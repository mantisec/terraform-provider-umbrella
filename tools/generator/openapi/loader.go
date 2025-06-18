package openapi

import (
	"fmt"
	"os"

	"github.com/pb33f/libopenapi"
	v3 "github.com/pb33f/libopenapi/datamodel/high/v3"
)

// Loader handles loading and parsing OpenAPI documents using libopenapi
type Loader struct {
	// Future: add caching, validation options, etc.
}

// NewLoader creates a new OpenAPI loader
func NewLoader() *Loader {
	return &Loader{}
}

// LoadFromFile loads an OpenAPI document from a file path
func (l *Loader) LoadFromFile(path string) (*v3.Document, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", path, err)
	}

	return l.LoadFromBytes(data)
}

// LoadFromBytes loads an OpenAPI document from byte data
func (l *Loader) LoadFromBytes(data []byte) (*v3.Document, error) {
	// Create document from bytes
	doc, err := libopenapi.NewDocument(data)
	if err != nil {
		return nil, fmt.Errorf("failed to create document: %w", err)
	}

	// Check document version
	info := doc.GetSpecInfo()
	if info.SpecType != "openapi" {
		return nil, fmt.Errorf("unsupported spec type: %s (expected openapi)", info.SpecType)
	}

	if info.Version != "3.0.0" && info.Version != "3.0.1" && info.Version != "3.0.2" && info.Version != "3.0.3" && info.Version != "3.1.0" {
		return nil, fmt.Errorf("unsupported OpenAPI version: %s", info.Version)
	}

	// Build the v3 model with full reference resolution
	model, errs := doc.BuildV3Model()
	if len(errs) > 0 {
		// Collect all errors
		var errMsgs []string
		for _, e := range errs {
			errMsgs = append(errMsgs, e.Error())
		}
		return nil, fmt.Errorf("failed to build v3 model: %v", errMsgs)
	}

	if model == nil {
		return nil, fmt.Errorf("failed to build v3 model: model is nil")
	}

	return &model.Model, nil
}

// ValidateDocument performs basic validation on the loaded document
func (l *Loader) ValidateDocument(doc *v3.Document) error {
	if doc == nil {
		return fmt.Errorf("document is nil")
	}

	if doc.Info == nil {
		return fmt.Errorf("document info is missing")
	}

	if doc.Info.Title == "" {
		return fmt.Errorf("document title is required")
	}

	if doc.Info.Version == "" {
		return fmt.Errorf("document version is required")
	}

	if doc.Paths == nil || doc.Paths.PathItems == nil || doc.Paths.PathItems.Len() == 0 {
		return fmt.Errorf("document has no paths defined")
	}

	return nil
}

// GetDocumentInfo extracts basic information from the document
func (l *Loader) GetDocumentInfo(doc *v3.Document) DocumentInfo {
	info := DocumentInfo{}

	if doc.Info != nil {
		info.Title = doc.Info.Title
		info.Version = doc.Info.Version
		if doc.Info.Description != "" {
			info.Description = doc.Info.Description
		}
	}

	if doc.Servers != nil {
		for _, server := range doc.Servers {
			if server.URL != "" {
				serverInfo := ServerInfo{
					URL: server.URL,
				}
				if server.Description != "" {
					serverInfo.Description = server.Description
				}
				info.Servers = append(info.Servers, serverInfo)
			}
		}
	}

	// Count paths and operations
	if doc.Paths != nil && doc.Paths.PathItems != nil {
		info.PathCount = doc.Paths.PathItems.Len()
		for pair := doc.Paths.PathItems.First(); pair != nil; pair = pair.Next() {
			pathItem := pair.Value()
			if pathItem.Get != nil {
				info.OperationCount++
			}
			if pathItem.Post != nil {
				info.OperationCount++
			}
			if pathItem.Put != nil {
				info.OperationCount++
			}
			if pathItem.Delete != nil {
				info.OperationCount++
			}
			if pathItem.Patch != nil {
				info.OperationCount++
			}
			if pathItem.Head != nil {
				info.OperationCount++
			}
			if pathItem.Options != nil {
				info.OperationCount++
			}
			if pathItem.Trace != nil {
				info.OperationCount++
			}
		}
	}

	// Count schemas
	if doc.Components != nil && doc.Components.Schemas != nil {
		info.SchemaCount = doc.Components.Schemas.Len()
	}

	return info
}

// DocumentInfo contains basic information about a loaded document
type DocumentInfo struct {
	Title          string
	Version        string
	Description    string
	Servers        []ServerInfo
	PathCount      int
	OperationCount int
	SchemaCount    int
}

// ServerInfo contains information about an API server
type ServerInfo struct {
	URL         string
	Description string
}

// String returns a string representation of the document info
func (di DocumentInfo) String() string {
	return fmt.Sprintf("OpenAPI %s v%s: %d paths, %d operations, %d schemas",
		di.Title, di.Version, di.PathCount, di.OperationCount, di.SchemaCount)
}
