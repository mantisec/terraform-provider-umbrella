# 🐳 Dev Container Setup Guide

This guide will help you set up the VS Code Dev Container for the Terraform Provider Umbrella project, which will solve all Windows compatibility issues and provide a consistent Linux development environment.

## 🚀 Quick Start

### Prerequisites

1. **Docker Desktop** - Download and install from [docker.com](https://www.docker.com/products/docker-desktop/)
2. **VS Code** - Download from [code.visualstudio.com](https://code.visualstudio.com/)
3. **Dev Containers Extension** - Install from VS Code marketplace

### Setup Steps

1. **Start Docker Desktop**

   - Ensure Docker Desktop is running
   - You should see the Docker icon in your system tray

2. **Open Project in VS Code**

   ```bash
   cd c:\Users\asus\Projects\Mantisec\terraform-provider-umbrella
   code .
   ```

3. **Open in Dev Container**

   - VS Code should automatically detect the dev container configuration
   - Click "Reopen in Container" when prompted
   - **OR** use Command Palette (`Ctrl+Shift+P`) → "Dev Containers: Reopen in Container"

4. **Wait for Setup**

   - First build takes 5-10 minutes
   - All dependencies are installed automatically
   - You'll see progress in the VS Code terminal

5. **Verify Setup**
   ```bash
   # These commands should work in the integrated terminal
   make help
   go version
   terraform version
   ```

## 🎯 What This Solves

### **Windows Compatibility Issues**

- ❌ **Before:** `make` command not found
- ✅ **After:** Full Linux environment with all Unix tools

- ❌ **Before:** Path separator issues (Windows `\` vs Unix `/`)
- ✅ **After:** Consistent Unix paths everywhere

- ❌ **Before:** PowerShell vs Bash script compatibility
- ✅ **After:** Native Bash/Zsh environment

- ❌ **Before:** Windows-specific Go build issues
- ✅ **After:** Linux Go environment with proper toolchain

### **Development Environment**

- ❌ **Before:** Manual tool installation and configuration
- ✅ **After:** Pre-configured environment with all tools

- ❌ **Before:** Version mismatches between team members
- ✅ **After:** Identical environment for everyone

- ❌ **Before:** Complex setup for new developers
- ✅ **After:** One-click setup with dev container

## 🛠️ Included Tools & Configuration

### **Development Tools**

- **Go 1.21+** with full toolchain
- **Terraform** latest stable
- **Make** for build automation
- **Git** with proper configuration
- **Docker** for containerized workflows

### **Go Development Tools**

- `golangci-lint` - Comprehensive linter
- `goimports` - Import formatter
- `dlv` - Delve debugger
- `staticcheck` - Static analysis
- `gotests` - Test generator

### **Terraform Tools**

- `tflint` - Terraform linter
- `terragrunt` - Terraform wrapper
- `tfplugindocs` - Plugin documentation
- `terraform-docs` - Documentation generator
- `tfsec` - Security scanner

### **VS Code Configuration**

- **Extensions:** All necessary extensions pre-installed
- **Settings:** Optimized for Go and Terraform development
- **Debugging:** Pre-configured launch configurations
- **Tasks:** Build, test, and generate tasks ready to use
- **Terminal:** Zsh with Oh My Zsh for better experience

### **Shell Aliases & Functions**

```bash
# Terraform shortcuts
tf='terraform'
tfi='terraform init'
tfp='terraform plan'
tfa='terraform apply'

# Go shortcuts
gob='go build'
gor='go run'
got='go test'
gol='golangci-lint run'

# Project shortcuts
gen-full='make generate-full'
clean-gen='make clean-generated'
build-provider='go build -o terraform-provider-umbrella'
test-provider='go test ./...'
```

## 🔧 Roo Code Extension Integration

The dev container is specifically configured to ensure the **Roo Code VS Code extension** uses the container's shell:

### **Terminal Configuration**

```json
{
  "terminal.integrated.defaultProfile.linux": "zsh",
  "terminal.integrated.automationProfile.linux": {
    "path": "/bin/zsh",
    "args": ["-l"]
  }
}
```

### **What This Means**

- ✅ All Roo Code commands execute in the Linux container
- ✅ `make` commands work perfectly
- ✅ Unix paths and tools available
- ✅ Consistent environment regardless of host OS

## 📋 Development Workflow

### **1. Code Generation**

```bash
# Clean and regenerate (now works on Windows!)
make clean-generated
make generate-full
```

### **2. Building**

```bash
# Build the provider
make build

# Install locally for testing
make install-local
```

### **3. Testing**

```bash
# Unit tests
make test

# Acceptance tests (requires API credentials)
TF_ACC=1 go test ./internal/provider/tests/ -v
```

### **4. Code Quality**

```bash
# Format and lint
make fmt
make lint

# All quality checks
make fmt && make lint && make test
```

## 🐛 Troubleshooting

### **Container Won't Start**

1. Ensure Docker Desktop is running
2. Check available disk space (needs ~2GB)
3. Try: Command Palette → "Dev Containers: Rebuild Container"

### **Extensions Not Working**

1. Wait for container to fully initialize
2. Reload VS Code window (`Ctrl+Shift+P` → "Developer: Reload Window")
3. Check that all recommended extensions are installed

### **Terminal Issues**

1. Open new terminal (` Ctrl+Shift+``  `)
2. Verify shell: `echo $SHELL` should show `/bin/zsh`
3. If needed: Command Palette → "Terminal: Select Default Profile"

### **Go Tools Not Found**

1. Check Go environment: `go env`
2. Verify GOPATH: `echo $GOPATH`
3. Reinstall tools: Run setup script again

### **Make Commands Fail**

1. Verify make is installed: `make --version`
2. Check current directory: `pwd`
3. Ensure you're in the project root

## 🎉 Benefits

### **For You (Windows User)**

- **No WSL Required:** Full Linux environment without WSL complexity
- **Native Unix Tools:** All commands work as documented
- **Consistent Experience:** Same as Linux/Mac developers
- **No Path Issues:** No more Windows/Unix path conflicts

### **For the Team**

- **Reproducible Environment:** Everyone has identical setup
- **Faster Onboarding:** New developers up and running in minutes
- **Consistent Builds:** No "works on my machine" issues
- **Easy Updates:** Update container config, rebuild for everyone

### **For the Project**

- **CI/CD Consistency:** Same environment as production builds
- **Documentation Accuracy:** All examples work for everyone
- **Reduced Support:** Fewer environment-specific issues
- **Better Collaboration:** Consistent tooling and workflows

## 🚀 Next Steps

1. **Open in Dev Container** - Follow the quick start above
2. **Verify Setup** - Run `make help` to see all available commands
3. **Start Developing** - All tools are ready to use!
4. **Generate Code** - Try `make generate-full` to test the generator
5. **Build Provider** - Run `make build` to create the binary

## 📚 Additional Resources

- [VS Code Dev Containers Documentation](https://code.visualstudio.com/docs/devcontainers/containers)
- [Docker Desktop Documentation](https://docs.docker.com/desktop/)
- [Terraform Provider Development Guide](https://developer.hashicorp.com/terraform/plugin)

---

**🎯 The dev container eliminates all Windows compatibility issues and provides a professional, consistent development environment for the Terraform Provider Umbrella project!**
