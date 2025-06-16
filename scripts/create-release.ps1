# Create Release Script for terraform-provider-umbrella (PowerShell version)
# This script helps create new releases by tagging and pushing to trigger the release workflow

param(
    [string]$Version = "",
    [switch]$Force = $false
)

# Colors for output
$Colors = @{
    Red = "Red"
    Green = "Green"
    Yellow = "Yellow"
    Blue = "Blue"
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

# Check if we're in a git repository
try {
    git rev-parse --git-dir | Out-Null
} catch {
    Write-Error "Not in a git repository"
    exit 1
}

# Check if we're on main branch
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

# Get current version
$currentVersion = git tag -l | Where-Object { $_ -match '^v\d+\.\d+\.\d+$' } | Sort-Object { [version]($_ -replace '^v', '') } | Select-Object -Last 1

if (-not $currentVersion) {
    $currentVersion = "v0.0.0"
    Write-Warning "No previous version found, starting from $currentVersion"
} else {
    Write-Status "Current version: $currentVersion"
}

# Parse version components
if ($currentVersion -match '^v(\d+)\.(\d+)\.(\d+)$') {
    $major = [int]$Matches[1]
    $minor = [int]$Matches[2]
    $patch = [int]$Matches[3]
} else {
    $major = 0
    $minor = 0
    $patch = 0
}

# Calculate next versions
$nextPatch = "v$major.$minor.$($patch + 1)"
$nextMinor = "v$major.$($minor + 1).0"
$nextMajor = "v$($major + 1).0.0"

if (-not $Version) {
    Write-Host ""
    Write-Status "Available version options:"
    Write-Host "  1) Patch release: $nextPatch (bug fixes)"
    Write-Host "  2) Minor release: $nextMinor (new features, backward compatible)"
    Write-Host "  3) Major release: $nextMajor (breaking changes)"
    Write-Host "  4) Custom version"
    Write-Host ""

    $versionChoice = Read-Host "Select version type (1-4)"

    switch ($versionChoice) {
        "1" { $Version = $nextPatch }
        "2" { $Version = $nextMinor }
        "3" { $Version = $nextMajor }
        "4" { 
            $Version = Read-Host "Enter custom version (e.g., v1.2.3)"
            if ($Version -notmatch '^v\d+\.\d+\.\d+$') {
                Write-Error "Invalid version format. Use vX.Y.Z (e.g., v1.2.3)"
                exit 1
            }
        }
        default {
            Write-Error "Invalid choice"
            exit 1
        }
    }
}

# Validate version format
if ($Version -notmatch '^v\d+\.\d+\.\d+$') {
    Write-Error "Invalid version format. Use vX.Y.Z (e.g., v1.2.3)"
    exit 1
}

# Check if tag already exists
$existingTags = git tag -l
if ($existingTags -contains $Version) {
    Write-Error "Tag $Version already exists"
    exit 1
}

Write-Status "Creating release $Version"

# Confirm before proceeding
if (-not $Force) {
    Write-Host ""
    Write-Warning "This will:"
    Write-Host "  - Create and push tag: $Version"
    Write-Host "  - Trigger the GitHub Actions release workflow"
    Write-Host "  - Create a GitHub release with binaries"
    Write-Host ""
    $continue = Read-Host "Continue? (y/N)"

    if ($continue -notmatch "^[Yy]$") {
        Write-Error "Aborted"
        exit 1
    }
}

try {
    # Pull latest changes
    Write-Status "Pulling latest changes..."
    git pull origin main

    # Create and push tag
    Write-Status "Creating tag $Version..."
    git tag -a $Version -m "Release $Version"

    Write-Status "Pushing tag to origin..."
    git push origin $Version

    Write-Success "Tag $Version created and pushed!"
    Write-Status "GitHub Actions will now build and create the release."
    
    # Get repository URL
    $repoUrl = git config --get remote.origin.url
    $repoPath = if ($repoUrl -match 'github\.com[:/]([^.]+)') { $Matches[1] } else { "your-repo" }
    
    Write-Status "Check the progress at: https://github.com/$repoPath/actions"
    Write-Host ""
    Write-Success "Release process initiated! 🎉"
    Write-Status "The release will be available at: https://github.com/$repoPath/releases"

} catch {
    Write-Error "Failed to create release: $($_.Exception.Message)"
    exit 1
}