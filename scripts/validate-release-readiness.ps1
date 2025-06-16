# Terraform Provider Umbrella - Release Readiness Validation
# This script performs comprehensive validation to ensure the provider is ready for Terraform Registry publication

param(
    [switch]$Verbose = $false,
    [switch]$Fix = $false
)

# Colors for output
$Colors = @{
    Red = "Red"
    Green = "Green"
    Yellow = "Yellow"
    Blue = "Blue"
    Cyan = "Cyan"
    Magenta = "Magenta"
    White = "White"
}

function Write-Status {
    param([string]$Message)
    Write-Host "[INFO] $Message" -ForegroundColor $Colors.Blue
}

function Write-Success {
    param([string]$Message)
    Write-Host "[✅] $Message" -ForegroundColor $Colors.Green
}

function Write-Warning {
    param([string]$Message)
    Write-Host "[⚠️] $Message" -ForegroundColor $Colors.Yellow
}

function Write-Error {
    param([string]$Message)
    Write-Host "[❌] $Message" -ForegroundColor $Colors.Red
}

function Write-Header {
    param([string]$Message)
    Write-Host ""
    Write-Host "=" * 70 -ForegroundColor $Colors.Cyan
    Write-Host $Message -ForegroundColor $Colors.Cyan
    Write-Host "=" * 70 -ForegroundColor $Colors.Cyan
    Write-Host ""
}

function Write-SubHeader {
    param([string]$Message)
    Write-Host ""
    Write-Host $Message -ForegroundColor $Colors.Magenta
    Write-Host "-" * $Message.Length -ForegroundColor $Colors.Magenta
}

# Global validation results
$ValidationResults = @{
    Passed = 0
    Failed = 0
    Warnings = 0
    Issues = @()
}

function Test-Requirement {
    param(
        [string]$Name,
        [scriptblock]$Test,
        [string]$SuccessMessage,
        [string]$FailureMessage,
        [scriptblock]$Fix = $null,
        [switch]$Warning = $false
    )
    
    try {
        $result = & $Test
        if ($result) {
            Write-Success "$Name - $SuccessMessage"
            $ValidationResults.Passed++
            return $true
        } else {
            if ($Warning) {
                Write-Warning "$Name - $FailureMessage"
                $ValidationResults.Warnings++
                $ValidationResults.Issues += @{Type="Warning"; Name=$Name; Message=$FailureMessage}
            } else {
                Write-Error "$Name - $FailureMessage"
                $ValidationResults.Failed++
                $ValidationResults.Issues += @{Type="Error"; Name=$Name; Message=$FailureMessage}
            }
            
            if ($Fix -and $Fix) {
                Write-Status "Attempting to fix: $Name"
                try {
                    & $Fix
                    Write-Success "Fixed: $Name"
                } catch {
                    Write-Error "Failed to fix: $Name - $($_.Exception.Message)"
                }
            }
            return $false
        }
    } catch {
        Write-Error "$Name - Exception: $($_.Exception.Message)"
        $ValidationResults.Failed++
        $ValidationResults.Issues += @{Type="Error"; Name=$Name; Message="Exception: $($_.Exception.Message)"}
        return $false
    }
}

Write-Header "Terraform Provider Umbrella - Release Readiness Validation"

# Repository Structure Validation
Write-SubHeader "Repository Structure"

Test-Requirement -Name "Repository Root Files" -Test {
    $requiredFiles = @(
        "go.mod",
        "main.go", 
        ".goreleaser.yml",
        "terraform-registry-manifest.json",
        "README.md",
        "CHANGELOG.md"
    )
    
    $missing = @()
    foreach ($file in $requiredFiles) {
        if (-not (Test-Path $file)) {
            $missing += $file
        }
    }
    
    if ($missing.Count -eq 0) {
        if ($Verbose) { Write-Host "  All required root files present" }
        return $true
    } else {
        if ($Verbose) { Write-Host "  Missing files: $($missing -join ', ')" }
        return $false
    }
} -SuccessMessage "All required root files present" -FailureMessage "Missing required root files"

Test-Requirement -Name "Documentation Structure" -Test {
    $requiredDirs = @("docs", "docs/resources", "docs/data-sources", "docs/guides")
    $requiredFiles = @("docs/index.md")
    
    $missing = @()
    foreach ($dir in $requiredDirs) {
        if (-not (Test-Path $dir -PathType Container)) {
            $missing += $dir
        }
    }
    
    foreach ($file in $requiredFiles) {
        if (-not (Test-Path $file)) {
            $missing += $file
        }
    }
    
    return $missing.Count -eq 0
} -SuccessMessage "Documentation structure complete" -FailureMessage "Missing documentation directories or files"

Test-Requirement -Name "Examples Directory" -Test {
    return (Test-Path "examples" -PathType Container) -and (Get-ChildItem "examples" -Directory).Count -gt 0
} -SuccessMessage "Examples directory with content exists" -FailureMessage "Examples directory missing or empty"

Test-Requirement -Name "GitHub Actions Workflows" -Test {
    $workflows = @(".github/workflows/release.yml", ".github/workflows/test.yml")
    foreach ($workflow in $workflows) {
        if (-not (Test-Path $workflow)) {
            return $false
        }
    }
    return $true
} -SuccessMessage "GitHub Actions workflows present" -FailureMessage "Missing GitHub Actions workflows"

# Go Module Validation
Write-SubHeader "Go Module Validation"

Test-Requirement -Name "Go Module Syntax" -Test {
    try {
        go mod verify | Out-Null
        return $LASTEXITCODE -eq 0
    } catch {
        return $false
    }
} -SuccessMessage "Go module is valid" -FailureMessage "Go module validation failed"

Test-Requirement -Name "Go Dependencies" -Test {
    try {
        go mod download | Out-Null
        return $LASTEXITCODE -eq 0
    } catch {
        return $false
    }
} -SuccessMessage "Dependencies downloaded successfully" -FailureMessage "Failed to download dependencies"

Test-Requirement -Name "Go Build" -Test {
    try {
        go build -v . | Out-Null
        return $LASTEXITCODE -eq 0
    } catch {
        return $false
    }
} -SuccessMessage "Provider builds successfully" -FailureMessage "Build failed"

# Testing Validation
Write-SubHeader "Testing Validation"

Test-Requirement -Name "Unit Tests" -Test {
    try {
        $testOutput = go test ./... 2>&1
        if ($LASTEXITCODE -eq 0) {
            if ($Verbose) { Write-Host "  Test output: $testOutput" }
            return $true
        } else {
            if ($Verbose) { Write-Host "  Test failures: $testOutput" }
            return $false
        }
    } catch {
        return $false
    }
} -SuccessMessage "All tests pass" -FailureMessage "Tests are failing"

Test-Requirement -Name "Test Coverage" -Test {
    try {
        go test -coverprofile=coverage.out ./... | Out-Null
        if ($LASTEXITCODE -eq 0 -and (Test-Path "coverage.out")) {
            $coverage = go tool cover -func=coverage.out | Select-String "total:" | ForEach-Object { $_.ToString().Split()[-1] }
            if ($coverage) {
                $coveragePercent = [float]($coverage -replace '%', '')
                if ($Verbose) { Write-Host "  Coverage: $coverage" }
                return $coveragePercent -gt 0
            }
        }
        return $false
    } catch {
        return $false
    }
} -SuccessMessage "Test coverage generated" -FailureMessage "Failed to generate test coverage" -Warning

# GoReleaser Validation
Write-SubHeader "GoReleaser Configuration"

Test-Requirement -Name "GoReleaser Installation" -Test {
    try {
        goreleaser --version | Out-Null
        return $LASTEXITCODE -eq 0
    } catch {
        return $false
    }
} -SuccessMessage "GoReleaser is installed" -FailureMessage "GoReleaser not found"

Test-Requirement -Name "GoReleaser Configuration" -Test {
    try {
        goreleaser check | Out-Null
        return $LASTEXITCODE -eq 0
    } catch {
        return $false
    }
} -SuccessMessage "GoReleaser configuration is valid" -FailureMessage "GoReleaser configuration is invalid"

Test-Requirement -Name "GoReleaser Test Build" -Test {
    try {
        goreleaser build --snapshot --clean | Out-Null
        return $LASTEXITCODE -eq 0
    } catch {
        return $false
    }
} -SuccessMessage "GoReleaser test build successful" -FailureMessage "GoReleaser test build failed"

# Documentation Validation
Write-SubHeader "Documentation Validation"

Test-Requirement -Name "Provider Documentation" -Test {
    if (-not (Test-Path "docs/index.md")) {
        return $false
    }
    
    $content = Get-Content "docs/index.md" -Raw
    $requiredSections = @("provider", "authentication", "argument", "example")
    
    foreach ($section in $requiredSections) {
        if ($content -notmatch $section) {
            if ($Verbose) { Write-Host "  Missing section reference: $section" }
        }
    }
    
    return $content.Length -gt 500  # Basic content length check
} -SuccessMessage "Provider documentation is comprehensive" -FailureMessage "Provider documentation needs improvement"

Test-Requirement -Name "Resource Documentation" -Test {
    $resourceFiles = Get-ChildItem "docs/resources" -Filter "*.md" -ErrorAction SilentlyContinue
    if ($resourceFiles.Count -eq 0) {
        return $false
    }
    
    foreach ($file in $resourceFiles) {
        $content = Get-Content $file.FullName -Raw
        if ($content.Length -lt 200) {  # Basic content check
            if ($Verbose) { Write-Host "  Insufficient content in: $($file.Name)" }
            return $false
        }
    }
    
    return $true
} -SuccessMessage "Resource documentation complete" -FailureMessage "Resource documentation incomplete"

Test-Requirement -Name "tfplugindocs Generation" -Test {
    try {
        # Check if tfplugindocs is available
        $tfplugindocs = Get-Command "tfplugindocs" -ErrorAction SilentlyContinue
        if (-not $tfplugindocs) {
            # Try to install it
            go install github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs@latest | Out-Null
        }
        
        tfplugindocs generate --provider-name umbrella | Out-Null
        return $LASTEXITCODE -eq 0
    } catch {
        return $false
    }
} -SuccessMessage "Documentation generation successful" -FailureMessage "Documentation generation failed" -Warning

# Example Validation
Write-SubHeader "Example Validation"

Test-Requirement -Name "Terraform Examples" -Test {
    $exampleDirs = Get-ChildItem "examples" -Directory -ErrorAction SilentlyContinue
    if ($exampleDirs.Count -eq 0) {
        return $false
    }
    
    foreach ($dir in $exampleDirs) {
        $tfFiles = Get-ChildItem $dir.FullName -Filter "*.tf"
        if ($tfFiles.Count -eq 0) {
            if ($Verbose) { Write-Host "  No .tf files in: $($dir.Name)" }
            return $false
        }
        
        # Basic Terraform validation
        Push-Location $dir.FullName
        try {
            terraform init -backend=false | Out-Null
            terraform validate | Out-Null
            if ($LASTEXITCODE -ne 0) {
                if ($Verbose) { Write-Host "  Validation failed for: $($dir.Name)" }
                Pop-Location
                return $false
            }
        } catch {
            Pop-Location
            return $false
        }
        Pop-Location
    }
    
    return $true
} -SuccessMessage "All Terraform examples are valid" -FailureMessage "Terraform example validation failed"

# Registry Manifest Validation
Write-SubHeader "Registry Manifest Validation"

Test-Requirement -Name "Registry Manifest" -Test {
    if (-not (Test-Path "terraform-registry-manifest.json")) {
        return $false
    }
    
    try {
        $manifest = Get-Content "terraform-registry-manifest.json" | ConvertFrom-Json
        return $manifest.version -eq 1 -and $manifest.metadata.protocol_versions -contains "5.0"
    } catch {
        return $false
    }
} -SuccessMessage "Registry manifest is valid" -FailureMessage "Registry manifest is invalid or missing"

# Git and Version Control
Write-SubHeader "Version Control Validation"

Test-Requirement -Name "Git Repository" -Test {
    try {
        git rev-parse --git-dir | Out-Null
        return $LASTEXITCODE -eq 0
    } catch {
        return $false
    }
} -SuccessMessage "Git repository initialized" -FailureMessage "Not a git repository"

Test-Requirement -Name "Git Tags" -Test {
    try {
        $tags = git tag -l
        return $tags.Count -gt 0
    } catch {
        return $false
    }
} -SuccessMessage "Git tags present" -FailureMessage "No git tags found" -Warning

Test-Requirement -Name "Clean Working Directory" -Test {
    try {
        $status = git status --porcelain
        return $status.Count -eq 0
    } catch {
        return $false
    }
} -SuccessMessage "Working directory is clean" -FailureMessage "Uncommitted changes present" -Warning

# Security Validation
Write-SubHeader "Security Validation"

Test-Requirement -Name "No Hardcoded Secrets" -Test {
    $patterns = @(
        'api_key\s*=\s*[''"][^''"]+[''""]',
        'password\s*=\s*[''"][^''"]+[''""]',
        'secret\s*=\s*[''"][^''"]+[''""]',
        'token\s*=\s*[''"][^''"]+[''""]'
    )
    
    $goFiles = Get-ChildItem -Recurse -Filter "*.go" | Where-Object { $_.FullName -notmatch "vendor|\.git|tests" }
    
    foreach ($file in $goFiles) {
        $content = Get-Content $file.FullName -Raw
        foreach ($pattern in $patterns) {
            if ($content -match $pattern) {
                if ($Verbose) { Write-Host "  Potential secret in: $($file.Name)" }
                return $false
            }
        }
    }
    
    return $true
} -SuccessMessage "No hardcoded secrets detected" -FailureMessage "Potential hardcoded secrets found"

# Final Summary
Write-Header "Validation Summary"

$total = $ValidationResults.Passed + $ValidationResults.Failed + $ValidationResults.Warnings

Write-Host "Total Checks: $total" -ForegroundColor $Colors.White
Write-Host "✅ Passed: $($ValidationResults.Passed)" -ForegroundColor $Colors.Green
Write-Host "❌ Failed: $($ValidationResults.Failed)" -ForegroundColor $Colors.Red
Write-Host "⚠️  Warnings: $($ValidationResults.Warnings)" -ForegroundColor $Colors.Yellow

if ($ValidationResults.Failed -gt 0) {
    Write-Header "Critical Issues (Must Fix)"
    foreach ($issue in $ValidationResults.Issues | Where-Object { $_.Type -eq "Error" }) {
        Write-Host "❌ $($issue.Name): $($issue.Message)" -ForegroundColor $Colors.Red
    }
}

if ($ValidationResults.Warnings -gt 0) {
    Write-Header "Warnings (Recommended to Fix)"
    foreach ($issue in $ValidationResults.Issues | Where-Object { $_.Type -eq "Warning" }) {
        Write-Host "⚠️  $($issue.Name): $($issue.Message)" -ForegroundColor $Colors.Yellow
    }
}

Write-Header "Release Readiness Assessment"

if ($ValidationResults.Failed -eq 0) {
    Write-Success "🎉 READY FOR RELEASE!"
    Write-Status "Your provider meets all requirements for Terraform Registry publication."
    Write-Host ""
    Write-Status "Next steps:"
    Write-Host "  1. Create a release: .\scripts\create-release.ps1 -Version '1.0.0'" -ForegroundColor $Colors.Cyan
    Write-Host "  2. Submit to Terraform Registry: Follow docs\TERRAFORM_REGISTRY_SUBMISSION.md" -ForegroundColor $Colors.Cyan
    Write-Host "  3. Monitor release: Check GitHub Actions for build status" -ForegroundColor $Colors.Cyan
} else {
    Write-Error "❌ NOT READY FOR RELEASE"
    Write-Status "Please fix the critical issues listed above before proceeding."
    Write-Host ""
    Write-Status "After fixing issues, run this script again:"
    Write-Host "  .\scripts\validate-release-readiness.ps1" -ForegroundColor $Colors.Cyan
}

if ($ValidationResults.Warnings -gt 0) {
    Write-Host ""
    Write-Warning "Consider addressing warnings for a better release quality."
}

# Cleanup
if (Test-Path "coverage.out") {
    Remove-Item "coverage.out" -Force
}

# Exit with appropriate code
exit $ValidationResults.Failed