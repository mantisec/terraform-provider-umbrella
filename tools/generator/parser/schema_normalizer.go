package parser

import (
	"fmt"
	"strings"
)

// SchemaNormalizer handles normalization and reference resolution for OpenAPI schemas
type SchemaNormalizer struct {
	resolvedRefs map[string]*Schema
}

// NewSchemaNormalizer creates a new schema normalizer
func NewSchemaNormalizer() *SchemaNormalizer {
	return &SchemaNormalizer{
		resolvedRefs: make(map[string]*Schema),
	}
}

// NormalizeSpec normalizes the entire OpenAPI specification
func (n *SchemaNormalizer) NormalizeSpec(spec *APISpec) error {
	// First pass: collect all component schemas
	if spec.Components.Schemas != nil {
		for name, schema := range spec.Components.Schemas {
			n.resolvedRefs["#/components/schemas/"+name] = schema
		}
	}

	// Second pass: resolve references in paths
	for path, pathItem := range spec.Paths {
		if err := n.normalizePathItem(&pathItem); err != nil {
			return fmt.Errorf("failed to normalize path %s: %w", path, err)
		}
	}

	// Third pass: resolve references in component schemas
	if spec.Components.Schemas != nil {
		for _, schema := range spec.Components.Schemas {
			if err := n.resolveSchemaReferences(schema); err != nil {
				return fmt.Errorf("failed to resolve schema references: %w", err)
			}
		}
	}

	return nil
}

// normalizePathItem normalizes all operations in a path item
func (n *SchemaNormalizer) normalizePathItem(pathItem *PathItem) error {
	operations := []*Operation{pathItem.Get, pathItem.Post, pathItem.Put, pathItem.Delete, pathItem.Patch}

	for _, op := range operations {
		if op != nil {
			if err := n.normalizeOperation(op); err != nil {
				return err
			}
		}
	}

	return nil
}

// normalizeOperation normalizes an operation
func (n *SchemaNormalizer) normalizeOperation(op *Operation) error {
	// Normalize parameters
	for _, param := range op.Parameters {
		if param.Schema != nil {
			if err := n.resolveSchemaReferences(param.Schema); err != nil {
				return err
			}
		}
	}

	// Normalize request body
	if op.RequestBody != nil {
		for _, mediaType := range op.RequestBody.Content {
			if mediaType.Schema != nil {
				if err := n.resolveSchemaReferences(mediaType.Schema); err != nil {
					return err
				}
			}
		}
	}

	// Normalize responses
	for _, response := range op.Responses {
		for _, mediaType := range response.Content {
			if mediaType.Schema != nil {
				if err := n.resolveSchemaReferences(mediaType.Schema); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

// resolveSchemaReferences resolves $ref references in a schema
func (n *SchemaNormalizer) resolveSchemaReferences(schema *Schema) error {
	if schema == nil {
		return nil
	}

	// Handle $ref
	if schema.Ref != "" {
		resolved, exists := n.resolvedRefs[schema.Ref]
		if !exists {
			return fmt.Errorf("unresolved reference: %s", schema.Ref)
		}
		// Copy resolved schema properties to current schema
		*schema = *resolved
		schema.Ref = "" // Clear the reference
	}

	// Recursively resolve properties
	for _, prop := range schema.Properties {
		if err := n.resolveSchemaReferences(prop); err != nil {
			return err
		}
	}

	// Resolve items schema for arrays
	if schema.Items != nil {
		if err := n.resolveSchemaReferences(schema.Items); err != nil {
			return err
		}
	}

	// Resolve allOf, oneOf, anyOf
	for _, subSchema := range schema.AllOf {
		if err := n.resolveSchemaReferences(subSchema); err != nil {
			return err
		}
	}
	for _, subSchema := range schema.OneOf {
		if err := n.resolveSchemaReferences(subSchema); err != nil {
			return err
		}
	}
	for _, subSchema := range schema.AnyOf {
		if err := n.resolveSchemaReferences(subSchema); err != nil {
			return err
		}
	}

	return nil
}

// NormalizeResponsePatterns analyzes response patterns and normalizes them
func (n *SchemaNormalizer) NormalizeResponsePatterns(spec *APISpec) map[string]ResponsePattern {
	patterns := make(map[string]ResponsePattern)

	for path, pathItem := range spec.Paths {
		operations := map[string]*Operation{
			"GET":    pathItem.Get,
			"POST":   pathItem.Post,
			"PUT":    pathItem.Put,
			"DELETE": pathItem.Delete,
			"PATCH":  pathItem.Patch,
		}

		for method, op := range operations {
			if op == nil {
				continue
			}

			pattern := n.analyzeResponsePattern(op)
			patternKey := fmt.Sprintf("%s %s", method, path)
			patterns[patternKey] = pattern
		}
	}

	return patterns
}

// ResponsePattern represents a normalized response pattern
type ResponsePattern struct {
	SuccessSchema *Schema
	ErrorSchema   *Schema
	IsList        bool
	HasPagination bool
	StatusCodes   []string
}

// analyzeResponsePattern analyzes the response pattern of an operation
func (n *SchemaNormalizer) analyzeResponsePattern(op *Operation) ResponsePattern {
	pattern := ResponsePattern{
		StatusCodes: make([]string, 0),
	}

	for statusCode, response := range op.Responses {
		pattern.StatusCodes = append(pattern.StatusCodes, statusCode)

		// Check if this is a success response (2xx)
		if strings.HasPrefix(statusCode, "2") {
			for contentType, mediaType := range response.Content {
				if strings.Contains(contentType, "json") && mediaType.Schema != nil {
					pattern.SuccessSchema = mediaType.Schema

					// Check if it's a list response
					if mediaType.Schema.Type == "array" {
						pattern.IsList = true
					}

					// Check for pagination indicators
					if n.hasPaginationIndicators(mediaType.Schema) {
						pattern.HasPagination = true
					}
				}
			}
		} else {
			// Error response
			for contentType, mediaType := range response.Content {
				if strings.Contains(contentType, "json") && mediaType.Schema != nil {
					pattern.ErrorSchema = mediaType.Schema
				}
			}
		}
	}

	return pattern
}

// hasPaginationIndicators checks if a schema has pagination indicators
func (n *SchemaNormalizer) hasPaginationIndicators(schema *Schema) bool {
	if schema == nil || schema.Properties == nil {
		return false
	}

	paginationFields := []string{"page", "limit", "offset", "total", "count", "next", "previous", "has_more"}

	for field := range schema.Properties {
		fieldLower := strings.ToLower(field)
		for _, paginationField := range paginationFields {
			if strings.Contains(fieldLower, paginationField) {
				return true
			}
		}
	}

	return false
}
