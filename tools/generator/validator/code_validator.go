package validator

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/mantisec/terraform-provider-umbrella/tools/generator/config"
)

// CodeValidator validates generated code for quality and compliance
type CodeValidator struct {
	config     *config.AdvancedConfig
	fileSet    *token.FileSet
	violations []Violation
}

// Violation represents a code quality violation
type Violation struct {
	Type       string `json:"type"`
	Severity   string `json:"severity"`
	Message    string `json:"message"`
	File       string `json:"file"`
	Line       int    `json:"line"`
	Column     int    `json:"column"`
	Rule       string `json:"rule"`
	Suggestion string `json:"suggestion,omitempty"`
}

// NewCodeValidator creates a new code validator
func NewCodeValidator(config *config.AdvancedConfig) *CodeValidator {
	return &CodeValidator{
		config:     config,
		fileSet:    token.NewFileSet(),
		violations: make([]Violation, 0),
	}
}

// ValidateDirectory validates all Go files in a directory
func (v *CodeValidator) ValidateDirectory(dir string) ([]Violation, error) {
	v.violations = make([]Violation, 0)

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip non-Go files
		if !strings.HasSuffix(path, ".go") {
			return nil
		}

		// Skip test files for now (they have different rules)
		if strings.HasSuffix(path, "_test.go") {
			return nil
		}

		if err := v.validateFile(path); err != nil {
			return fmt.Errorf("failed to validate file %s: %w", path, err)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return v.violations, nil
}

// ValidateFile validates a single Go file
func (v *CodeValidator) ValidateFile(filePath string) ([]Violation, error) {
	v.violations = make([]Violation, 0)

	if err := v.validateFile(filePath); err != nil {
		return nil, err
	}

	return v.violations, nil
}

// validateFile performs validation on a single file
func (v *CodeValidator) validateFile(filePath string) error {
	// Read file content
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	// Parse the Go file
	file, err := parser.ParseFile(v.fileSet, filePath, content, parser.ParseComments)
	if err != nil {
		v.addViolation(Violation{
			Type:     "syntax_error",
			Severity: "error",
			Message:  fmt.Sprintf("Failed to parse Go file: %v", err),
			File:     filePath,
			Rule:     "syntax",
		})
		return nil // Continue with other validations
	}

	// Perform various validations
	v.validateSyntaxAndStructure(file, filePath)
	v.validateCodeQuality(file, filePath)
	v.validateSecurityRules(string(content), filePath)
	v.validateTerraformCompliance(file, filePath)
	v.validateNamingConventions(file, filePath)
	v.validateDocumentation(file, filePath)

	return nil
}

// validateSyntaxAndStructure validates basic syntax and structure
func (v *CodeValidator) validateSyntaxAndStructure(file *ast.File, filePath string) {
	ast.Inspect(file, func(n ast.Node) bool {
		switch node := n.(type) {
		case *ast.FuncDecl:
			v.validateFunction(node, filePath)
		case *ast.GenDecl:
			v.validateDeclaration(node, filePath)
		case *ast.IfStmt:
			v.validateIfStatement(node, filePath)
		}
		return true
	})
}

// validateFunction validates function-specific rules
func (v *CodeValidator) validateFunction(fn *ast.FuncDecl, filePath string) {
	if fn.Body == nil {
		return
	}

	// Check function length
	for _, rule := range v.config.QualityAssurance.CodeQuality.CustomRules {
		if rule.Name == "function_length" && rule.MaxLines > 0 {
			startPos := v.fileSet.Position(fn.Pos())
			endPos := v.fileSet.Position(fn.End())
			lineCount := endPos.Line - startPos.Line + 1

			if lineCount > rule.MaxLines {
				v.addViolation(Violation{
					Type:       "code_quality",
					Severity:   "warning",
					Message:    rule.Message,
					File:       filePath,
					Line:       startPos.Line,
					Column:     startPos.Column,
					Rule:       rule.Name,
					Suggestion: fmt.Sprintf("Consider breaking this %d-line function into smaller functions", lineCount),
				})
			}
		}
	}

	// Check cyclomatic complexity
	complexity := v.calculateCyclomaticComplexity(fn)
	for _, rule := range v.config.QualityAssurance.CodeQuality.CustomRules {
		if rule.Name == "cyclomatic_complexity" && rule.MaxComplexity > 0 {
			if complexity > rule.MaxComplexity {
				pos := v.fileSet.Position(fn.Pos())
				v.addViolation(Violation{
					Type:       "code_quality",
					Severity:   "warning",
					Message:    rule.Message,
					File:       filePath,
					Line:       pos.Line,
					Column:     pos.Column,
					Rule:       rule.Name,
					Suggestion: fmt.Sprintf("Function has complexity %d, consider refactoring", complexity),
				})
			}
		}
	}

	// Validate function naming
	if fn.Name != nil && fn.Name.IsExported() {
		if !v.isValidExportedFunctionName(fn.Name.Name) {
			pos := v.fileSet.Position(fn.Pos())
			v.addViolation(Violation{
				Type:       "naming",
				Severity:   "warning",
				Message:    "Exported function name should follow Go naming conventions",
				File:       filePath,
				Line:       pos.Line,
				Column:     pos.Column,
				Rule:       "function_naming",
				Suggestion: "Use PascalCase for exported functions",
			})
		}
	}
}

// validateDeclaration validates type and variable declarations
func (v *CodeValidator) validateDeclaration(decl *ast.GenDecl, filePath string) {
	for _, spec := range decl.Specs {
		switch s := spec.(type) {
		case *ast.TypeSpec:
			v.validateTypeDeclaration(s, filePath)
		case *ast.ValueSpec:
			v.validateValueDeclaration(s, filePath)
		}
	}
}

// validateTypeDeclaration validates type declarations
func (v *CodeValidator) validateTypeDeclaration(spec *ast.TypeSpec, filePath string) {
	if spec.Name.IsExported() && !v.isValidExportedTypeName(spec.Name.Name) {
		pos := v.fileSet.Position(spec.Pos())
		v.addViolation(Violation{
			Type:       "naming",
			Severity:   "warning",
			Message:    "Exported type name should follow Go naming conventions",
			File:       filePath,
			Line:       pos.Line,
			Column:     pos.Column,
			Rule:       "type_naming",
			Suggestion: "Use PascalCase for exported types",
		})
	}
}

// validateValueDeclaration validates variable and constant declarations
func (v *CodeValidator) validateValueDeclaration(spec *ast.ValueSpec, filePath string) {
	for _, name := range spec.Names {
		if name.IsExported() && !v.isValidExportedVariableName(name.Name) {
			pos := v.fileSet.Position(name.Pos())
			v.addViolation(Violation{
				Type:       "naming",
				Severity:   "warning",
				Message:    "Exported variable name should follow Go naming conventions",
				File:       filePath,
				Line:       pos.Line,
				Column:     pos.Column,
				Rule:       "variable_naming",
				Suggestion: "Use PascalCase for exported variables",
			})
		}
	}
}

// validateIfStatement validates if statements for potential issues
func (v *CodeValidator) validateIfStatement(stmt *ast.IfStmt, filePath string) {
	// Check for empty if blocks
	if stmt.Body != nil && len(stmt.Body.List) == 0 {
		pos := v.fileSet.Position(stmt.Pos())
		v.addViolation(Violation{
			Type:       "code_quality",
			Severity:   "warning",
			Message:    "Empty if block detected",
			File:       filePath,
			Line:       pos.Line,
			Column:     pos.Column,
			Rule:       "empty_blocks",
			Suggestion: "Remove empty if block or add implementation",
		})
	}
}

// validateCodeQuality performs general code quality checks
func (v *CodeValidator) validateCodeQuality(file *ast.File, filePath string) {
	// Check for TODO comments
	for _, commentGroup := range file.Comments {
		for _, comment := range commentGroup.List {
			if strings.Contains(strings.ToUpper(comment.Text), "TODO") {
				pos := v.fileSet.Position(comment.Pos())
				v.addViolation(Violation{
					Type:       "code_quality",
					Severity:   "info",
					Message:    "TODO comment found",
					File:       filePath,
					Line:       pos.Line,
					Column:     pos.Column,
					Rule:       "todo_comments",
					Suggestion: "Consider creating an issue or implementing the TODO",
				})
			}
		}
	}

	// Check for magic numbers
	ast.Inspect(file, func(n ast.Node) bool {
		if lit, ok := n.(*ast.BasicLit); ok && lit.Kind == token.INT {
			if v.isMagicNumber(lit.Value) {
				pos := v.fileSet.Position(lit.Pos())
				v.addViolation(Violation{
					Type:       "code_quality",
					Severity:   "warning",
					Message:    "Magic number detected",
					File:       filePath,
					Line:       pos.Line,
					Column:     pos.Column,
					Rule:       "magic_numbers",
					Suggestion: "Consider using a named constant",
				})
			}
		}
		return true
	})
}

// validateSecurityRules validates security-related rules
func (v *CodeValidator) validateSecurityRules(content, filePath string) {
	for _, rule := range v.config.Validation.SecurityRules {
		pattern, err := regexp.Compile(rule.Pattern)
		if err != nil {
			v.addViolation(Violation{
				Type:     "validator_error",
				Severity: "error",
				Message:  fmt.Sprintf("Invalid regex pattern in security rule %s: %v", rule.Name, err),
				File:     filePath,
				Rule:     rule.Name,
			})
			continue
		}

		matches := pattern.FindAllStringIndex(content, -1)
		for _, match := range matches {
			line := strings.Count(content[:match[0]], "\n") + 1
			v.addViolation(Violation{
				Type:     "security",
				Severity: rule.Severity,
				Message:  rule.Message,
				File:     filePath,
				Line:     line,
				Rule:     rule.Name,
			})
		}
	}
}

// validateTerraformCompliance validates Terraform-specific compliance
func (v *CodeValidator) validateTerraformCompliance(file *ast.File, filePath string) {
	// Check for proper Terraform provider patterns
	ast.Inspect(file, func(n ast.Node) bool {
		switch node := n.(type) {
		case *ast.FuncDecl:
			v.validateTerraformFunction(node, filePath)
		case *ast.StructType:
			v.validateTerraformStruct(node, filePath)
		}
		return true
	})
}

// validateTerraformFunction validates Terraform-specific function patterns
func (v *CodeValidator) validateTerraformFunction(fn *ast.FuncDecl, filePath string) {
	if fn.Name == nil {
		return
	}

	funcName := fn.Name.Name

	// Check for proper CRUD function naming
	if strings.HasSuffix(funcName, "Create") || strings.HasSuffix(funcName, "Read") ||
		strings.HasSuffix(funcName, "Update") || strings.HasSuffix(funcName, "Delete") {

		// Validate function signature for CRUD operations
		if fn.Type.Params == nil || len(fn.Type.Params.List) < 2 {
			pos := v.fileSet.Position(fn.Pos())
			v.addViolation(Violation{
				Type:       "terraform_compliance",
				Severity:   "error",
				Message:    "Terraform CRUD functions must have context and request parameters",
				File:       filePath,
				Line:       pos.Line,
				Column:     pos.Column,
				Rule:       "terraform_crud_signature",
				Suggestion: "Add context.Context and appropriate request/response parameters",
			})
		}
	}
}

// validateTerraformStruct validates Terraform-specific struct patterns
func (v *CodeValidator) validateTerraformStruct(structType *ast.StructType, filePath string) {
	// Check for proper field tags in Terraform structs
	for _, field := range structType.Fields.List {
		if field.Tag != nil {
			tagValue := field.Tag.Value
			if strings.Contains(tagValue, "tfsdk:") {
				// Validate tfsdk tag format
				if !v.isValidTFSDKTag(tagValue) {
					pos := v.fileSet.Position(field.Pos())
					v.addViolation(Violation{
						Type:       "terraform_compliance",
						Severity:   "warning",
						Message:    "Invalid tfsdk tag format",
						File:       filePath,
						Line:       pos.Line,
						Column:     pos.Column,
						Rule:       "tfsdk_tag_format",
						Suggestion: "Use proper tfsdk tag format: `tfsdk:\"field_name\"`",
					})
				}
			}
		}
	}
}

// validateNamingConventions validates Go naming conventions
func (v *CodeValidator) validateNamingConventions(file *ast.File, filePath string) {
	// This is handled in individual validation functions
	// Additional naming convention checks can be added here
}

// validateDocumentation validates documentation requirements
func (v *CodeValidator) validateDocumentation(file *ast.File, filePath string) {
	ast.Inspect(file, func(n ast.Node) bool {
		switch node := n.(type) {
		case *ast.FuncDecl:
			if node.Name != nil && node.Name.IsExported() {
				if !v.hasDocumentation(node, file) {
					pos := v.fileSet.Position(node.Pos())
					v.addViolation(Violation{
						Type:       "documentation",
						Severity:   "warning",
						Message:    "Exported function missing documentation",
						File:       filePath,
						Line:       pos.Line,
						Column:     pos.Column,
						Rule:       "exported_function_docs",
						Suggestion: fmt.Sprintf("Add documentation comment starting with '%s'", node.Name.Name),
					})
				}
			}
		case *ast.GenDecl:
			for _, spec := range node.Specs {
				if typeSpec, ok := spec.(*ast.TypeSpec); ok && typeSpec.Name.IsExported() {
					if !v.hasDocumentation(node, file) {
						pos := v.fileSet.Position(typeSpec.Pos())
						v.addViolation(Violation{
							Type:       "documentation",
							Severity:   "warning",
							Message:    "Exported type missing documentation",
							File:       filePath,
							Line:       pos.Line,
							Column:     pos.Column,
							Rule:       "exported_type_docs",
							Suggestion: fmt.Sprintf("Add documentation comment starting with '%s'", typeSpec.Name.Name),
						})
					}
				}
			}
		}
		return true
	})
}

// Helper functions

// addViolation adds a violation to the list
func (v *CodeValidator) addViolation(violation Violation) {
	v.violations = append(v.violations, violation)
}

// calculateCyclomaticComplexity calculates the cyclomatic complexity of a function
func (v *CodeValidator) calculateCyclomaticComplexity(fn *ast.FuncDecl) int {
	complexity := 1 // Base complexity

	ast.Inspect(fn, func(n ast.Node) bool {
		switch n.(type) {
		case *ast.IfStmt, *ast.ForStmt, *ast.RangeStmt, *ast.SwitchStmt, *ast.TypeSwitchStmt:
			complexity++
		case *ast.CaseClause:
			complexity++
		}
		return true
	})

	return complexity
}

// isValidExportedFunctionName checks if an exported function name follows Go conventions
func (v *CodeValidator) isValidExportedFunctionName(name string) bool {
	if len(name) == 0 {
		return false
	}

	// First character should be uppercase
	if name[0] < 'A' || name[0] > 'Z' {
		return false
	}

	// Check for common abbreviations that should be uppercase
	abbreviations := []string{"ID", "URL", "HTTP", "JSON", "XML", "API", "UUID"}
	for _, abbr := range abbreviations {
		if strings.Contains(name, strings.ToLower(abbr)) && !strings.Contains(name, abbr) {
			return false
		}
	}

	return true
}

// isValidExportedTypeName checks if an exported type name follows Go conventions
func (v *CodeValidator) isValidExportedTypeName(name string) bool {
	return v.isValidExportedFunctionName(name)
}

// isValidExportedVariableName checks if an exported variable name follows Go conventions
func (v *CodeValidator) isValidExportedVariableName(name string) bool {
	return v.isValidExportedFunctionName(name)
}

// isMagicNumber checks if a number literal is a magic number
func (v *CodeValidator) isMagicNumber(value string) bool {
	// Common non-magic numbers
	nonMagicNumbers := []string{"0", "1", "2", "10", "100", "1000"}
	for _, nonMagic := range nonMagicNumbers {
		if value == nonMagic {
			return false
		}
	}

	// Numbers with more than 2 digits are likely magic numbers
	return len(value) > 2
}

// isValidTFSDKTag checks if a tfsdk tag is properly formatted
func (v *CodeValidator) isValidTFSDKTag(tag string) bool {
	// Basic validation for tfsdk tag format
	pattern := regexp.MustCompile(`tfsdk:"[a-z_][a-z0-9_]*"`)
	return pattern.MatchString(tag)
}

// hasDocumentation checks if a declaration has documentation
func (v *CodeValidator) hasDocumentation(node ast.Node, file *ast.File) bool {
	nodePos := v.fileSet.Position(node.Pos())

	for _, commentGroup := range file.Comments {
		commentPos := v.fileSet.Position(commentGroup.End())

		// Check if comment is immediately before the node
		if commentPos.Line == nodePos.Line-1 {
			return true
		}
	}

	return false
}

// GetViolationSummary returns a summary of violations by type and severity
func (v *CodeValidator) GetViolationSummary(violations []Violation) map[string]map[string]int {
	summary := make(map[string]map[string]int)

	for _, violation := range violations {
		if summary[violation.Type] == nil {
			summary[violation.Type] = make(map[string]int)
		}
		summary[violation.Type][violation.Severity]++
	}

	return summary
}
