#!/usr/bin/env pwsh
# PowerShell script to clean generated files from API specifications

param(
    [switch]$All,
    [switch]$Help
)

if ($Help) {
    Write-Host "Clean Generated Files Script"
    Write-Host ""
    Write-Host "Usage:"
    Write-Host "  .\scripts\clean-generated.ps1           - Clean only generated files"
    Write-Host "  .\scripts\clean-generated.ps1 -All      - Clean generated files + build artifacts"
    Write-Host "  .\scripts\clean-generated.ps1 -Help     - Show this help"
    Write-Host ""
    Write-Host "Generated files that will be cleaned:"
    Write-Host "  - internal/provider/generated_*.go"
    Write-Host "  - internal/provider/tests/*_test.go (except provider_test.go)"
    Write-Host "  - docs/resources/*.md"
    exit 0
}

Write-Host "Cleaning generated files from API specifications..." -ForegroundColor Yellow

# Clean generated Go files in internal/provider/
Write-Host "Removing generated Go files..." -ForegroundColor Cyan
$generatedFiles = Get-ChildItem -Path "internal/provider/generated_*.go" -ErrorAction SilentlyContinue
if ($generatedFiles) {
    $generatedFiles | Remove-Item -Force
    Write-Host "  Removed $($generatedFiles.Count) generated Go files" -ForegroundColor Green
} else {
    Write-Host "  No generated Go files found" -ForegroundColor Gray
}

# Also clean old naming convention files that have generation markers
Write-Host "Checking for old generated files with markers..." -ForegroundColor Cyan
$oldGeneratedFiles = @()
$resourceFiles = Get-ChildItem -Path "internal/provider/resource_*.go" -ErrorAction SilentlyContinue
$dataSourceFiles = Get-ChildItem -Path "internal/provider/data_source_*.go" -ErrorAction SilentlyContinue
foreach ($file in ($resourceFiles + $dataSourceFiles)) {
    $firstLine = Get-Content $file.FullName -First 1 -ErrorAction SilentlyContinue
    if ($firstLine -and $firstLine.Contains("Code generated")) {
        $oldGeneratedFiles += $file
    }
}
if ($oldGeneratedFiles) {
    $oldGeneratedFiles | Remove-Item -Force
    Write-Host "  Removed $($oldGeneratedFiles.Count) old generated files" -ForegroundColor Green
} else {
    Write-Host "  No old generated files found" -ForegroundColor Gray
}

# Clean generated test files (keeping manually created ones)
Write-Host "Removing generated test files..." -ForegroundColor Cyan
$testFiles = Get-ChildItem -Path "internal/provider/tests/*_test.go" -ErrorAction SilentlyContinue | Where-Object { $_.Name -ne "provider_test.go" }
if ($testFiles) {
    $testFiles | Remove-Item -Force
    Write-Host "  Removed $($testFiles.Count) generated test files" -ForegroundColor Green
    Write-Host "  Kept provider_test.go (manually maintained)" -ForegroundColor Gray
} else {
    Write-Host "  No generated test files found" -ForegroundColor Gray
}

# Clean generated documentation files
Write-Host "Removing generated documentation files..." -ForegroundColor Cyan
$docFiles = Get-ChildItem -Path "docs/resources/*.md" -ErrorAction SilentlyContinue
if ($docFiles) {
    $docFiles | Remove-Item -Force
    Write-Host "  Removed $($docFiles.Count) generated documentation files" -ForegroundColor Green
    Write-Host "  Kept docs/index.md (manually maintained)" -ForegroundColor Gray
} else {
    Write-Host "  No generated documentation files found" -ForegroundColor Gray
}

if ($All) {
    Write-Host "Cleaning build artifacts..." -ForegroundColor Yellow

    # Clean binary files
    $binaryFiles = Get-ChildItem -Path "terraform-provider-umbrella*" -ErrorAction SilentlyContinue
    if ($binaryFiles) {
        $binaryFiles | Remove-Item -Force
        Write-Host "  Removed $($binaryFiles.Count) binary files" -ForegroundColor Green
    }

    # Clean coverage files
    $coverageFiles = Get-ChildItem -Path "coverage.*" -ErrorAction SilentlyContinue
    if ($coverageFiles) {
        $coverageFiles | Remove-Item -Force
        Write-Host "  Removed $($coverageFiles.Count) coverage files" -ForegroundColor Green
    }
}

Write-Host "Generated files cleaned successfully!" -ForegroundColor Green

# Show what files are still present
Write-Host ""
Write-Host "Remaining files:" -ForegroundColor Yellow
Write-Host "Manual Go files:" -ForegroundColor Cyan
Get-ChildItem -Path "internal/provider/*.go" -ErrorAction SilentlyContinue | Where-Object { $_.Name -notlike "generated_*" } | ForEach-Object { Write-Host "  $($_.Name)" -ForegroundColor Gray }

Write-Host "Manual test files:" -ForegroundColor Cyan
Get-ChildItem -Path "internal/provider/tests/*.go" -ErrorAction SilentlyContinue | ForEach-Object { Write-Host "  $($_.Name)" -ForegroundColor Gray }

Write-Host "Manual documentation:" -ForegroundColor Cyan
Get-ChildItem -Path "docs/*.md" -ErrorAction SilentlyContinue | ForEach-Object { Write-Host "  $($_.Name)" -ForegroundColor Gray }
