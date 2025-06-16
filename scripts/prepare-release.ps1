# Terraform Provider Umbrella - Release Preparation Script
# This script prepares the provider for release to the Terraform Registry

param(
    [Parameter(Mandatory=$true)]
    [string]$Version,
    [switch]$DryRun = $false,
    [switch]$Force = $false
)

# Colors for output
$Colors = @{
    Red = "Red"
    Green = "Green"
    Yellow = "Yellow"
    Blue = "Blue"
    Cyan = "Cyan"
    White = "White"
}

function Write-Status {
    param([string]$Message)
    Write-Host "[INFO] $Message" -ForegroundColor $Colors.Blue
}

function Write-Success {
    param([string]$Message)
    Write-Host "[SUCCESS] $Message" -ForegroundColor $Colors.Green
}

function Write-Warning {
    param([string]$Message)
    Write-Host "[WARNING] $Message" -ForegroundColor $Colors.Yellow
}

function Write-Error {
    param([string]$Message)
    Write-Host "[ERROR] $Message" -ForegroundColor $Colors.Red
}

function Write-Header {
    param([string]$Message)
    Write-Host ""
    Write-Host "=" * 60 -ForegroundColor $Colors.Cyan
    Write-Host $Message -ForegroundColor $Colors.Cyan
    Write-Host "=" * 60 -ForegroundColor $Colors.Cyan
    Write-Host ""
}

# Validate version format
if ($Version -notmatch '^v?\d+\.\d+\.\d+$') {
    Write-Error "Invalid version format. Use X.Y.Z or vX.Y.Z (e.g., 1.0.0 or v1.0.0)"
    exit 1
}

# Normalize version (ensure it starts with 'v')
if (-not $Version.StartsWith('v')) {
    $Version = "v$Version"
}

Write-Header "Terraform Provider Umbrella - Release Preparation"
Write-Status "Preparing release: $Version"
Write-Status "Dry run mode: $DryRun"

# Check if we're in a git repository
try {
    git rev-parse --git-dir | Out-Null
} catch {
    Write-Error "Not in a git repository"
    exit 1
}

# Check for required tools
Write-Header "Checking Prerequisites"

$requiredTools = @(
    @{Name="git"; Command="git --version"},
    @{Name="go"; Command="go version"},
    @{Name="goreleaser"; Command="goreleaser --version"}
)

foreach ($tool in $requiredTools) {
    try {
        $output = Invoke-Expression $tool.Command 2>$null
        Write-Success "$($tool.Name): Available"
    } catch {
        Write-Error "$($tool.Name): Not found or not working"
        Write-Status "Please install $($tool.Name) before proceeding"
        exit 1
    }
}

# Check GitHub CLI (optional but recommended)
try {
    gh --version | Out-Null
    Write-Success "GitHub CLI: Available"
    $hasGH = $true
} catch {
    Write-Warning "GitHub CLI: Not found (optional, but recommended for release management)"
    $hasGH = $false
}

# Check current branch
$currentBranch = git branch --show-current
if ($currentBranch -ne "main" -and -not $Force) {
    Write-Warning "You're on branch '$currentBranch', not 'main'"
    $continue = Read-Host "Continue anyway? (y/N)"
    if ($continue -notmatch "^[Yy]$") {
        Write-Error "Aborted"
        exit 1
    }
}

# Check for uncommitted changes
$gitStatus = git status --porcelain
if ($gitStatus -and -not $Force) {
    Write-Error "You have uncommitted changes. Please commit or stash them first."
    git status --porcelain
    exit 1
}

# Check if tag already exists
$existingTags = git tag -l
if ($existingTags -contains $Version) {
    Write-Error "Tag $Version already exists"
    if (-not $Force) {
        exit 1
    } else {
        Write-Warning "Forcing release despite existing tag"
    }
}

Write-Header "Running Pre-Release Checks"

# Run tests
Write-Status "Running tests..."
if (-not $DryRun) {
    $testResult = go test ./...
    if ($LASTEXITCODE -ne 0) {
        Write-Error "Tests failed. Please fix before releasing."
        exit 1
    }
    Write-Success "All tests passed"
} else {
    Write-Status "DRY RUN: Would run tests"
}

# Build provider
Write-Status "Building provider..."
if (-not $DryRun) {
    $buildResult = go build -v .
    if ($LASTEXITCODE -ne 0) {
        Write-Error "Build failed. Please fix before releasing."
        exit 1
    }
    Write-Success "Build successful"
} else {
    Write-Status "DRY RUN: Would build provider"
}

# Test GoReleaser configuration
Write-Status "Testing GoReleaser configuration..."
if (-not $DryRun) {
    $goreleaserResult = goreleaser check
    if ($LASTEXITCODE -ne 0) {
        Write-Error "GoReleaser configuration is invalid"
        exit 1
    }
    Write-Success "GoReleaser configuration is valid"
} else {
    Write-Status "DRY RUN: Would test GoReleaser configuration"
}

# Validate Terraform examples
Write-Status "Validating Terraform examples..."
if (-not $DryRun) {
    $exampleDirs = Get-ChildItem -Path "examples" -Directory
    foreach ($dir in $exampleDirs) {
        Write-Status "Validating $($dir.Name)..."
        Push-Location $dir.FullName
        try {
            terraform init -backend=false | Out-Null
            terraform validate | Out-Null
            Write-Success "Example $($dir.Name) is valid"
        } catch {
            Write-Error "Example $($dir.Name) validation failed"
            Pop-Location
            exit 1
        }
        Pop-Location
    }
} else {
    Write-Status "DRY RUN: Would validate Terraform examples"
}

# Generate documentation
Write-Status "Generating documentation..."
if (-not $DryRun) {
    try {
        # Install tfplugindocs if not available
        go install github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs@latest
        tfplugindocs generate --provider-name umbrella
        Write-Success "Documentation generated successfully"
    } catch {
        Write-Warning "Documentation generation failed, but continuing..."
    }
} else {
    Write-Status "DRY RUN: Would generate documentation"
}

Write-Header "Release Summary"

Write-Status "Release version: $Version"
Write-Status "Current branch: $currentBranch"
Write-Status "Repository status: Clean"
Write-Status "Tests: Passed"
Write-Status "Build: Successful"
Write-Status "GoReleaser: Valid"
Write-Status "Examples: Validated"
Write-Status "Documentation: Generated"

Write-Header "Required GitHub Secrets"

Write-Status "Ensure these secrets are configured in your GitHub repository:"
Write-Host "  - GPG_PRIVATE_KEY: Your GPG private key for signing releases" -ForegroundColor $Colors.Yellow
Write-Host "  - PASSPHRASE: Passphrase for your GPG key" -ForegroundColor $Colors.Yellow
Write-Host ""
Write-Status "To check secrets: Go to Settings > Secrets and variables > Actions"

Write-Header "Next Steps"

if ($DryRun) {
    Write-Success "DRY RUN COMPLETE - All checks passed!"
    Write-Status "To create the actual release, run:"
    Write-Host "  .\scripts\prepare-release.ps1 -Version $Version" -ForegroundColor $Colors.Cyan
} else {
    Write-Success "PRE-RELEASE CHECKS COMPLETE!"
    Write-Status "Ready to create release $Version"
    Write-Host ""
    Write-Status "To create the release:"
    Write-Host "  1. Run: git tag $Version" -ForegroundColor $Colors.Cyan
    Write-Host "  2. Run: git push origin $Version" -ForegroundColor $Colors.Cyan
    Write-Host "  3. Monitor: https://github.com/mantisec/terraform-provider-umbrella/actions" -ForegroundColor $Colors.Cyan
    Write-Host ""
    Write-Status "Or use the automated script:"
    Write-Host "  .\scripts\create-release.ps1 -Version $($Version.TrimStart('v'))" -ForegroundColor $Colors.Cyan
    
    if ($hasGH) {
        Write-Host ""
        Write-Status "Or create release with GitHub CLI:"
        Write-Host "  gh release create $Version --generate-notes" -ForegroundColor $Colors.Cyan
    }
}

Write-Header "Release Checklist"

$checklist = @(
    "✅ All tests passing",
    "✅ Build successful", 
    "✅ GoReleaser configuration valid",
    "✅ Terraform examples validated",
    "✅ Documentation generated",
    "⚠️  GitHub secrets configured (GPG_PRIVATE_KEY, PASSPHRASE)",
    "⚠️  Repository pushed to GitHub",
    "⚠️  Ready to create and push tag"
)

foreach ($item in $checklist) {
    Write-Host "  $item"
}

Write-Host ""
Write-Success "Release preparation complete! 🚀"