package performance

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/mantisec/terraform-provider-umbrella/tools/generator/config"
	"github.com/mantisec/terraform-provider-umbrella/tools/generator/parser"
)

// ParallelProcessor handles parallel processing of API specifications
type ParallelProcessor struct {
	config     *config.AdvancedConfig
	maxWorkers int
	batchSize  int
	enabled    bool
}

// ProcessingJob represents a single processing job
type ProcessingJob struct {
	ID       string
	SpecFile string
	Spec     *parser.APISpec
	Priority int
}

// ProcessingResult represents the result of a processing job
type ProcessingResult struct {
	Job       ProcessingJob
	Success   bool
	Error     error
	Duration  time.Duration
	Generated []GeneratedFile
}

// GeneratedFile represents a generated file
type GeneratedFile struct {
	Path    string
	Content string
	Type    string // "resource", "data_source", "client_method", etc.
}

// ProcessingStats represents processing statistics
type ProcessingStats struct {
	TotalJobs     int           `json:"total_jobs"`
	CompletedJobs int           `json:"completed_jobs"`
	FailedJobs    int           `json:"failed_jobs"`
	TotalDuration time.Duration `json:"total_duration"`
	AvgDuration   time.Duration `json:"avg_duration"`
	Throughput    float64       `json:"throughput"` // jobs per second
}

// NewParallelProcessor creates a new parallel processor
func NewParallelProcessor(config *config.AdvancedConfig) *ParallelProcessor {
	maxWorkers := config.Performance.ParallelProcessing.MaxWorkers
	if maxWorkers <= 0 {
		maxWorkers = runtime.NumCPU()
	}

	batchSize := config.Performance.ParallelProcessing.BatchSize
	if batchSize <= 0 {
		batchSize = 10
	}

	return &ParallelProcessor{
		config:     config,
		maxWorkers: maxWorkers,
		batchSize:  batchSize,
		enabled:    config.Performance.ParallelProcessing.Enabled,
	}
}

// ProcessSpecs processes multiple API specifications in parallel
func (p *ParallelProcessor) ProcessSpecs(ctx context.Context, jobs []ProcessingJob, processor JobProcessor) ([]ProcessingResult, ProcessingStats, error) {
	if !p.enabled || len(jobs) == 0 {
		// Process sequentially if parallel processing is disabled
		return p.processSequentially(ctx, jobs, processor)
	}

	startTime := time.Now()
	results := make([]ProcessingResult, len(jobs))
	stats := ProcessingStats{
		TotalJobs: len(jobs),
	}

	// Create job and result channels
	jobChan := make(chan ProcessingJob, len(jobs))
	resultChan := make(chan ProcessingResult, len(jobs))

	// Start workers
	var wg sync.WaitGroup
	for i := 0; i < p.maxWorkers; i++ {
		wg.Add(1)
		go p.worker(ctx, &wg, jobChan, resultChan, processor)
	}

	// Send jobs to workers
	go func() {
		defer close(jobChan)
		for _, job := range jobs {
			select {
			case jobChan <- job:
			case <-ctx.Done():
				return
			}
		}
	}()

	// Collect results
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Process results
	resultIndex := 0
	for result := range resultChan {
		results[resultIndex] = result
		resultIndex++

		stats.CompletedJobs++
		if !result.Success {
			stats.FailedJobs++
		}
		stats.TotalDuration += result.Duration
	}

	// Calculate statistics
	stats.TotalDuration = time.Since(startTime)
	if stats.CompletedJobs > 0 {
		stats.AvgDuration = stats.TotalDuration / time.Duration(stats.CompletedJobs)
		stats.Throughput = float64(stats.CompletedJobs) / stats.TotalDuration.Seconds()
	}

	return results, stats, nil
}

// ProcessBatches processes jobs in batches for better memory management
func (p *ParallelProcessor) ProcessBatches(ctx context.Context, jobs []ProcessingJob, processor JobProcessor) ([]ProcessingResult, ProcessingStats, error) {
	if !p.enabled {
		return p.processSequentially(ctx, jobs, processor)
	}

	var allResults []ProcessingResult
	var combinedStats ProcessingStats
	combinedStats.TotalJobs = len(jobs)

	// Process jobs in batches
	for i := 0; i < len(jobs); i += p.batchSize {
		end := i + p.batchSize
		if end > len(jobs) {
			end = len(jobs)
		}

		batch := jobs[i:end]
		results, stats, err := p.ProcessSpecs(ctx, batch, processor)
		if err != nil {
			return nil, combinedStats, err
		}

		allResults = append(allResults, results...)
		combinedStats.CompletedJobs += stats.CompletedJobs
		combinedStats.FailedJobs += stats.FailedJobs
		combinedStats.TotalDuration += stats.TotalDuration
	}

	// Calculate final statistics
	if combinedStats.CompletedJobs > 0 {
		combinedStats.AvgDuration = combinedStats.TotalDuration / time.Duration(combinedStats.CompletedJobs)
		combinedStats.Throughput = float64(combinedStats.CompletedJobs) / combinedStats.TotalDuration.Seconds()
	}

	return allResults, combinedStats, nil
}

// worker processes jobs from the job channel
func (p *ParallelProcessor) worker(ctx context.Context, wg *sync.WaitGroup, jobChan <-chan ProcessingJob, resultChan chan<- ProcessingResult, processor JobProcessor) {
	defer wg.Done()

	for {
		select {
		case job, ok := <-jobChan:
			if !ok {
				return
			}

			startTime := time.Now()
			generated, err := processor.ProcessJob(ctx, job)
			duration := time.Since(startTime)

			result := ProcessingResult{
				Job:       job,
				Success:   err == nil,
				Error:     err,
				Duration:  duration,
				Generated: generated,
			}

			select {
			case resultChan <- result:
			case <-ctx.Done():
				return
			}

		case <-ctx.Done():
			return
		}
	}
}

// processSequentially processes jobs one by one
func (p *ParallelProcessor) processSequentially(ctx context.Context, jobs []ProcessingJob, processor JobProcessor) ([]ProcessingResult, ProcessingStats, error) {
	startTime := time.Now()
	results := make([]ProcessingResult, len(jobs))
	stats := ProcessingStats{
		TotalJobs: len(jobs),
	}

	for i, job := range jobs {
		select {
		case <-ctx.Done():
			return results[:i], stats, ctx.Err()
		default:
		}

		jobStartTime := time.Now()
		generated, err := processor.ProcessJob(ctx, job)
		duration := time.Since(jobStartTime)

		results[i] = ProcessingResult{
			Job:       job,
			Success:   err == nil,
			Error:     err,
			Duration:  duration,
			Generated: generated,
		}

		stats.CompletedJobs++
		if err != nil {
			stats.FailedJobs++
		}
	}

	// Calculate statistics
	stats.TotalDuration = time.Since(startTime)
	if stats.CompletedJobs > 0 {
		stats.AvgDuration = stats.TotalDuration / time.Duration(stats.CompletedJobs)
		stats.Throughput = float64(stats.CompletedJobs) / stats.TotalDuration.Seconds()
	}

	return results, stats, nil
}

// JobProcessor interface for processing individual jobs
type JobProcessor interface {
	ProcessJob(ctx context.Context, job ProcessingJob) ([]GeneratedFile, error)
}

// DefaultJobProcessor implements JobProcessor for standard code generation
type DefaultJobProcessor struct {
	generator Generator
	cache     interface{} // Will be properly typed when cache.go is available
}

// Generator interface for code generation
type Generator interface {
	GenerateFromSpec(spec *parser.APISpec, outputDir string) error
}

// NewDefaultJobProcessor creates a new default job processor
func NewDefaultJobProcessor(generator Generator, cache interface{}) *DefaultJobProcessor {
	return &DefaultJobProcessor{
		generator: generator,
		cache:     cache,
	}
}

// ProcessJob processes a single job
func (p *DefaultJobProcessor) ProcessJob(ctx context.Context, job ProcessingJob) ([]GeneratedFile, error) {
	// Generate code
	tempDir := fmt.Sprintf("/tmp/generator_%s_%d", job.ID, time.Now().UnixNano())
	if err := p.generator.GenerateFromSpec(job.Spec, tempDir); err != nil {
		return nil, fmt.Errorf("failed to generate code for %s: %w", job.SpecFile, err)
	}

	// Collect generated files (simplified - in real implementation, scan the temp directory)
	generated := []GeneratedFile{
		{
			Path:    fmt.Sprintf("%s/resource_%s.go", tempDir, job.ID),
			Content: "// Generated resource code",
			Type:    "resource",
		},
		{
			Path:    fmt.Sprintf("%s/data_source_%s.go", tempDir, job.ID),
			Content: "// Generated data source code",
			Type:    "data_source",
		},
	}

	return generated, nil
}

// WorkerPool manages a pool of workers for processing
type WorkerPool struct {
	workers     int
	jobQueue    chan ProcessingJob
	resultQueue chan ProcessingResult
	quit        chan bool
	wg          sync.WaitGroup
}

// NewWorkerPool creates a new worker pool
func NewWorkerPool(workers int) *WorkerPool {
	return &WorkerPool{
		workers:     workers,
		jobQueue:    make(chan ProcessingJob, workers*2),
		resultQueue: make(chan ProcessingResult, workers*2),
		quit:        make(chan bool),
	}
}

// Start starts the worker pool
func (wp *WorkerPool) Start(processor JobProcessor) {
	for i := 0; i < wp.workers; i++ {
		wp.wg.Add(1)
		go wp.worker(processor)
	}
}

// Stop stops the worker pool
func (wp *WorkerPool) Stop() {
	close(wp.quit)
	wp.wg.Wait()
	close(wp.jobQueue)
	close(wp.resultQueue)
}

// SubmitJob submits a job to the worker pool
func (wp *WorkerPool) SubmitJob(job ProcessingJob) {
	wp.jobQueue <- job
}

// GetResult gets a result from the worker pool
func (wp *WorkerPool) GetResult() ProcessingResult {
	return <-wp.resultQueue
}

// worker processes jobs in the worker pool
func (wp *WorkerPool) worker(processor JobProcessor) {
	defer wp.wg.Done()

	for {
		select {
		case job := <-wp.jobQueue:
			startTime := time.Now()
			generated, err := processor.ProcessJob(context.Background(), job)
			duration := time.Since(startTime)

			result := ProcessingResult{
				Job:       job,
				Success:   err == nil,
				Error:     err,
				Duration:  duration,
				Generated: generated,
			}

			wp.resultQueue <- result

		case <-wp.quit:
			return
		}
	}
}

// ProgressTracker tracks processing progress
type ProgressTracker struct {
	total     int
	completed int
	failed    int
	mutex     sync.RWMutex
	startTime time.Time
}

// NewProgressTracker creates a new progress tracker
func NewProgressTracker(total int) *ProgressTracker {
	return &ProgressTracker{
		total:     total,
		startTime: time.Now(),
	}
}

// Update updates the progress
func (pt *ProgressTracker) Update(success bool) {
	pt.mutex.Lock()
	defer pt.mutex.Unlock()

	pt.completed++
	if !success {
		pt.failed++
	}
}

// GetProgress returns current progress
func (pt *ProgressTracker) GetProgress() (int, int, int, float64) {
	pt.mutex.RLock()
	defer pt.mutex.RUnlock()

	percentage := float64(pt.completed) / float64(pt.total) * 100
	return pt.completed, pt.failed, pt.total, percentage
}

// GetETA returns estimated time of arrival
func (pt *ProgressTracker) GetETA() time.Duration {
	pt.mutex.RLock()
	defer pt.mutex.RUnlock()

	if pt.completed == 0 {
		return 0
	}

	elapsed := time.Since(pt.startTime)
	avgTimePerJob := elapsed / time.Duration(pt.completed)
	remaining := pt.total - pt.completed

	return avgTimePerJob * time.Duration(remaining)
}

// String returns a string representation of progress
func (pt *ProgressTracker) String() string {
	completed, failed, total, percentage := pt.GetProgress()
	eta := pt.GetETA()

	return fmt.Sprintf("Progress: %d/%d (%.1f%%) - Failed: %d - ETA: %v",
		completed, total, percentage, failed, eta.Round(time.Second))
}
