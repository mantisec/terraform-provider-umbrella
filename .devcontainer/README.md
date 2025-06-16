# Dev Container Setup for Terraform Provider Umbrella

This directory contains the configuration for a VS Code Dev Container that provides a complete, consistent development environment for the Terraform Provider Umbrella project.

## 🚀 Quick Start

1. **Prerequisites:**
   - VS Code with the "Dev Containers" extension installed
   - Docker Desktop running on your machine

2. **Open in Dev Container:**
   - Open this project in VS Code
   - When prompted, click "Reopen in Container"
   - Or use Command Palette: `Dev Containers: Reopen in Container`

3. **Wait for Setup:**
   - The container will build and install all dependencies automatically
   - This may take 5-10 minutes on first run

4. **Start Developing:**
   - All tools are pre-installed and configured
   - Use the integrated terminal (it will be Linux/bash even on Windows)
   - Run `make help` to see available commands

## 📦 What's Included

### **Base Environment**
- **OS:** Debian Bullseye (Linux)
- **Go:** Latest stable version (1.21+)
- **Shell:** Zsh with Oh My Zsh
- **User:** `vscode` with sudo access

### **Development Tools**
- **Go Tools:**
  - `golangci-lint` - Comprehensive Go linter
  - `goimports` - Import formatter
  - `godoc` - Documentation generator
  - `dlv` - Delve debugger
  - `staticcheck` - Static analysis
  - `gomodifytags` - Struct tag modifier
  - `impl` - Interface implementation generator
  - `gotests` - Test generator

- **Terraform Tools:**
  - `terraform` - Latest stable version
  - `tflint` - Terraform linter
  - `terragrunt` - Terraform wrapper
  - `tfplugindocs` - Plugin documentation generator
  - `terraform-docs` - Documentation generator
  - `tfsec` - Security scanner
  - `checkov` - Additional security scanner

- **Build & Release Tools:**
  - `make` - Build automation
  - `goreleaser` - Release automation
  - `docker` - Container support
  - `git` - Version control
  - `github-cli` - GitHub integration

- **Development Utilities:**
  - `pre-commit` - Git hooks
  - `jq` - JSON processor
  - `tree` - Directory visualization
  - `curl`, `wget` - HTTP clients

### **VS Code Extensions**
- **Go Development:**
  - Go extension with full language support
  - Debugging and testing support
  
- **Terraform Development:**
  - HashiCorp Terraform extension
  - Syntax highlighting and validation
  
- **General Development:**
  - Makefile tools
  - YAML support
  - JSON support
  - Git integration (GitLens)
  - Spell checker
  - Test explorer

## 🔧 Configuration

### **Environment Variables**
```bash
GO111MODULE=on
GOPROXY=https://proxy.golang.org,direct
GOSUMDB=sum.golang.org
CGO_ENABLED=0
TF_LOG=INFO
TF_CLI_CONFIG_FILE=/workspaces/terraform-provider-umbrella/.terraformrc
```

### **Shell Aliases**
The container includes helpful aliases:

**Terraform:**
```bash
tf='terraform'
tfi='terraform init'
tfp='terraform plan'
tfa='terraform apply'
tfd='terraform destroy'
tfv='terraform validate'
tff='terraform fmt'
```

**Go Development:**
```bash
gob='go build'
gor='go run'
got='go test'
gom='go mod'
gof='go fmt'
gov='go vet'
gol='golangci-lint run'
```

**Project Specific:**
```bash
gen='go run tools/generator/cmd/generate/main.go'
gen-full='make generate-full'
clean-gen='make clean-generated'
build-provider='go build -o terraform-provider-umbrella'
test-provider='go test ./...'
```

### **Utility Functions**
```bash
tf-init-local()        # Initialize Terraform with local provider
provider-test()        # Run provider acceptance tests
provider-build()       # Build the provider binary
provider-install-local() # Install provider locally for testing
```

## 📁 File Structure

```
.devcontainer/
├── devcontainer.json     # Main dev container configuration
├── setup.sh             # Post-create setup script
├── pre-build.sh          # Pre-build script (runs on host)
├── on-create.sh          # On-create script (runs once)
├── update-content.sh     # Update script (runs on updates)
└── README.md            # This file
```

## 🛠️ Development Workflow

### **1. Code Generation**
```bash
# Clean and regenerate all code
make clean-generated
make generate-full

# Or use aliases
clean-gen
gen-full
```

### **2. Building**
```bash
# Build the provider
make build

# Or use alias
build-provider
```

### **3. Testing**
```bash
# Run unit tests
make test

# Run acceptance tests (requires API credentials)
TF_ACC=1 go test ./internal/provider/tests/ -v

# Or use function
provider-test
```

### **4. Local Development**
```bash
# Install provider locally for testing
make install-local

# Or use function
provider-install-local
```

### **5. Code Quality**
```bash
# Format code
make fmt

# Run linter
make lint

# Run all quality checks
make fmt && make lint && make test
```

## 🔍 Troubleshooting

### **Container Won't Start**
- Ensure Docker Desktop is running
- Check that you have enough disk space (container needs ~2GB)
- Try rebuilding: Command Palette → "Dev Containers: Rebuild Container"

### **Go Tools Not Working**
- Restart VS Code
- Check that GOPATH is set: `echo $GOPATH`
- Reinstall tools: `go install -a github.com/golangci/golangci-lint/cmd/golangci-lint@latest`

### **Terraform Commands Fail**
- Check Terraform version: `terraform version`
- Verify CLI config: `cat ~/.terraformrc`
- Ensure provider is built: `make build`

### **Permission Issues**
- All scripts should be executable
- If needed, run: `chmod +x .devcontainer/*.sh`

## 🎯 Benefits

### **For Windows Users**
- **No WSL Required:** Full Linux environment in container
- **Native Tools:** All Unix tools work natively
- **Consistent Environment:** Same setup across all team members
- **No Path Issues:** No Windows/Unix path conflicts

### **For All Users**
- **Isolated Environment:** No conflicts with host system
- **Reproducible:** Identical setup for all developers
- **Pre-configured:** All tools installed and configured
- **Version Locked:** Consistent tool versions across team

## 🔄 Updates

### **Updating the Container**
1. Pull latest changes from Git
2. Command Palette → "Dev Containers: Rebuild Container"
3. Wait for rebuild to complete

### **Adding New Tools**
1. Edit `.devcontainer/setup.sh`
2. Add installation commands
3. Rebuild container to test
4. Commit changes

## 📚 Additional Resources

- [VS Code Dev Containers Documentation](https://code.visualstudio.com/docs/devcontainers/containers)
- [Dev Container Features](https://containers.dev/features)
- [Terraform Provider Development](https://developer.hashicorp.com/terraform/plugin)
- [Go Development in VS Code](https://code.visualstudio.com/docs/languages/go)

## 🆘 Support

If you encounter issues with the dev container setup:

1. Check this README for troubleshooting steps
2. Try rebuilding the container
3. Check the container logs in VS Code
4. Ask for help in the team chat with error details

The dev container provides a complete, isolated development environment that eliminates platform-specific issues and ensures consistent development experience across the team.