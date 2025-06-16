# GPG Setup Guide for Terraform Provider Releases

This guide walks you through setting up GPG signing for secure Terraform provider releases.

## Overview

GPG (GNU Privacy Guard) signing is required for publishing providers to the Terraform Registry. It ensures the integrity and authenticity of your releases.

## Prerequisites

- Git installed and configured
- Access to your GitHub repository settings
- Command line access (PowerShell on Windows, Terminal on macOS/Linux)

## Step 1: Install GPG

### Windows
```powershell
# Using Chocolatey
choco install gnupg

# Using Scoop
scoop install gpg

# Or download from: https://gnupg.org/download/
```

### macOS
```bash
# Using Homebrew
brew install gnupg

# Using MacPorts
sudo port install gnupg2
```

### Linux
```bash
# Ubuntu/Debian
sudo apt-get install gnupg

# CentOS/RHEL/Fedora
sudo yum install gnupg2
# or
sudo dnf install gnupg2
```

## Step 2: Generate GPG Key

### Generate Key
```bash
gpg --full-generate-key
```

### Configuration Options
When prompted, choose:
1. **Key type**: `(1) RSA and RSA (default)`
2. **Key size**: `4096` (recommended for security)
3. **Expiration**: `0` (key does not expire) or set appropriate expiration
4. **Real name**: Your full name
5. **Email**: Your GitHub email address (important!)
6. **Comment**: Optional (e.g., "Terraform Provider Signing Key")
7. **Passphrase**: Choose a strong passphrase (you'll need this later)

### Example Output
```
gpg: key 1234567890ABCDEF marked as ultimately trusted
gpg: revocation certificate stored as '/home/user/.gnupg/openpgp-revocs.d/1234567890ABCDEF.rev'
public and secret key created and signed.

pub   rsa4096 2025-06-16 [SC]
      1234567890ABCDEF1234567890ABCDEF12345678
uid                      Your Name <your.email@example.com>
sub   rsa4096 2025-06-16 [E]
```

**Important**: Save the key ID (`1234567890ABCDEF12345678`) - you'll need it!

## Step 3: Export Keys

### Export Private Key
```bash
# Replace YOUR_KEY_ID with your actual key ID
gpg --armor --export-secret-keys YOUR_KEY_ID > private-key.asc
```

### Export Public Key
```bash
gpg --armor --export YOUR_KEY_ID > public-key.asc
```

### Get Key Fingerprint
```bash
gpg --list-secret-keys --keyid-format LONG
```

Example output:
```
sec   rsa4096/1234567890ABCDEF 2025-06-16 [SC]
      1234567890ABCDEF1234567890ABCDEF12345678
uid                 [ultimate] Your Name <your.email@example.com>
ssb   rsa4096/FEDCBA0987654321 2025-06-16 [E]
```

The fingerprint is: `1234567890ABCDEF1234567890ABCDEF12345678`

## Step 4: Configure GitHub Secrets

### Add Secrets to Repository
1. Go to your GitHub repository
2. Navigate to `Settings` > `Secrets and variables` > `Actions`
3. Click `New repository secret`

### Required Secrets

#### GPG_PRIVATE_KEY
- **Name**: `GPG_PRIVATE_KEY`
- **Value**: Contents of `private-key.asc` file
- **Note**: Include the entire content including `-----BEGIN PGP PRIVATE KEY BLOCK-----` and `-----END PGP PRIVATE KEY BLOCK-----`

#### PASSPHRASE
- **Name**: `PASSPHRASE`
- **Value**: The passphrase you set when creating the GPG key
- **Note**: Leave empty if you didn't set a passphrase (not recommended)

### Verify Secrets
After adding secrets, you should see:
- ✅ `GPG_PRIVATE_KEY`
- ✅ `PASSPHRASE`

## Step 5: Test GPG Configuration

### Test Signing Locally
```bash
# Create a test file
echo "test" > test.txt

# Sign the file
gpg --detach-sign --armor test.txt

# Verify the signature
gpg --verify test.txt.asc test.txt

# Clean up
rm test.txt test.txt.asc
```

### Test with GoReleaser
```bash
# Test release build (doesn't publish)
goreleaser release --snapshot --clean
```

## Step 6: Verify GitHub Actions Integration

### Check Workflow File
Ensure your `.github/workflows/release.yml` includes:

```yaml
- name: Import GPG key
  uses: crazy-max/ghaction-import-gpg@v6
  id: import_gpg
  with:
    gpg_private_key: ${{ secrets.GPG_PRIVATE_KEY }}
    passphrase: ${{ secrets.PASSPHRASE }}

- name: Run GoReleaser
  uses: goreleaser/goreleaser-action@v6
  with:
    distribution: goreleaser
    version: '~> v2'
    args: release --clean
  env:
    GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
    GPG_FINGERPRINT: ${{ steps.import_gpg.outputs.fingerprint }}
```

## Step 7: Upload Public Key to Terraform Registry

### When Publishing to Registry
1. Go to [registry.terraform.io](https://registry.terraform.io)
2. Sign in with GitHub
3. Navigate to your provider
4. Go to `Settings` > `Signing Keys`
5. Upload your public key (`public-key.asc`)

## Security Best Practices

### Key Management
- ✅ **Store private keys securely** - Never commit to version control
- ✅ **Use strong passphrases** - Protect your private key
- ✅ **Regular key rotation** - Consider expiring keys periodically
- ✅ **Backup keys safely** - Store in secure, encrypted location
- ✅ **Revoke compromised keys** - Have a revocation plan

### GitHub Secrets
- ✅ **Limit access** - Only necessary collaborators
- ✅ **Regular audit** - Review who has access
- ✅ **Environment-specific** - Use environment secrets for production
- ✅ **Monitor usage** - Check Actions logs for secret usage

### Release Process
- ✅ **Verify signatures** - Always check release signatures
- ✅ **Automated signing** - Use CI/CD for consistent signing
- ✅ **Audit trail** - Maintain logs of all releases
- ✅ **Rollback plan** - Be prepared to revoke bad releases

## Troubleshooting

### Common Issues

#### "gpg: signing failed: No such file or directory"
**Solution**: Ensure GPG is installed and in PATH
```bash
gpg --version
```

#### "gpg: signing failed: Inappropriate ioctl for device"
**Solution**: Set GPG_TTY environment variable
```bash
export GPG_TTY=$(tty)
```

#### "Error: failed to sign: exit status 2"
**Solution**: Check passphrase and key availability
```bash
gpg --list-secret-keys
```

#### GitHub Actions: "gpg: signing failed: No secret key"
**Solution**: Verify GitHub secrets are correctly configured
- Check `GPG_PRIVATE_KEY` contains full private key
- Verify `PASSPHRASE` matches your key passphrase

#### "gpg: can't connect to the agent: No such file or directory"
**Solution**: Start GPG agent
```bash
gpg-agent --daemon
```

### Debug Commands

```bash
# List all keys
gpg --list-keys
gpg --list-secret-keys

# Check key details
gpg --list-keys --keyid-format LONG

# Test key functionality
gpg --sign --armor --detach-sign --local-user YOUR_KEY_ID test.txt

# Check GPG agent status
gpg-connect-agent 'keyinfo --list' /bye
```

## Verification

### After Setup
1. ✅ GPG key generated with 4096-bit RSA
2. ✅ Private key exported and added to GitHub secrets
3. ✅ Public key exported for Terraform Registry
4. ✅ GitHub Actions workflow configured
5. ✅ Test signing works locally
6. ✅ GoReleaser test passes

### Before Each Release
1. ✅ GPG key is still valid (not expired)
2. ✅ GitHub secrets are accessible
3. ✅ Test release build succeeds
4. ✅ Signature verification works

## Additional Resources

- [GPG Documentation](https://gnupg.org/documentation/)
- [GitHub GPG Documentation](https://docs.github.com/en/authentication/managing-commit-signature-verification)
- [Terraform Registry Publishing](https://www.terraform.io/docs/registry/providers/publishing.html)
- [GoReleaser Documentation](https://goreleaser.com/customization/sign/)

## Support

If you encounter issues:
1. Check the troubleshooting section above
2. Review GitHub Actions logs for detailed error messages
3. Verify all prerequisites are met
4. Test GPG functionality locally before using in CI/CD

---

**Security Note**: Keep your private keys secure and never share them. If a key is compromised, revoke it immediately and generate a new one.