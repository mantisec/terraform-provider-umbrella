# Cisco Umbrella API Analysis Report

## Executive Summary

This report provides a comprehensive analysis of Cisco's Umbrella API documentation, including available API specifications, endpoint inventory, and recommendations for Terraform provider enhancement.

## Current Implementation Status

### ✅ **Currently Implemented APIs**

1. **Authentication API** (`/auth/v2`)
   - OAuth2 client credentials flow
   - Automatic token refresh
   - **Status**: ✅ Complete

2. **Policies API** (`/policies/v2`)
   - Destination Lists management
   - Rulesets management  
   - Rules management
   - **Status**: ✅ Partially implemented

3. **Secure Internet Gateway API** (`/v2`)
   - IPSec tunnels management
   - **Status**: ✅ Complete

4. **SAML Configuration API** (`/v2`)
   - SSO integration
   - **Status**: ✅ Complete

## 📋 **Complete API Inventory**

### 1. **Authentication API** ✅ IMPLEMENTED
- **Base URL**: `https://api.umbrella.com/auth/v2`
- **Swagger Spec**: ✅ Downloaded (`api-specs/auth_api.yaml`)
- **Endpoints**:
  - `POST /token` - Create authorization token

### 2. **Investigate API** 🆕 NEW OPPORTUNITY
- **Base URL**: `https://api.umbrella.com/investigate/v2`
- **Swagger Spec**: ✅ Downloaded (`api-specs/investigate_api.yaml`)
- **Size**: 137KB+ comprehensive specification
- **Categories**:

#### **Domain Intelligence**
- `GET /domains/categorization/{domain}` - Domain status and categorization
- `POST /domains/categorization` - Bulk domain categorization
- `GET /domains/volume/{domain}` - Domain query volume
- `GET /domains/risk-score/{domain}` - Domain risk scoring

#### **Threat Intelligence**
- `GET /security/name/{domain}` - Security scores and features
- `GET /timeline/{name}` - Tagging timeline for domains/IPs/URLs
- `GET /recommendations/name/{domain}` - Co-occurring domains
- `GET /links/name/{domain}` - Related domains

#### **DNS Intelligence**
- `GET /pdns/name/{domain}` - Passive DNS for domains
- `GET /pdns/domain/{domain}` - Resource records for domains
- `GET /pdns/ip/{ip}` - Resource records for IPs
- `GET /pdns/raw/{anystring}` - Raw DNS data

#### **Network Intelligence**
- `GET /bgp_routes/ip/{ip}/as_for_ip.json` - BGP route information
- `GET /bgp_routes/asn/{asn}/prefixes_for_asn.json` - ASN prefix information

#### **WHOIS Intelligence**
- `GET /whois/{domain}` - WHOIS information
- `GET /whois/{domain}/history` - Historical WHOIS data
- `GET /whois/nameservers/{nameserver}` - Nameserver WHOIS
- `GET /whois/nameservers` - Multiple nameservers WHOIS
- `GET /whois/emails/{email}` - Email-based WHOIS search
- `GET /whois/search/{searchField}/{regexExpression}` - RegEx WHOIS search

#### **Malware Analysis Integration**
- `GET /samples/{destination}` - Malware samples for domains/IPs/URLs
- `GET /sample/{hash}` - Sample information by hash
- `GET /sample/{hash}/artifacts` - Sample artifacts
- `GET /sample/{hash}/connections` - Sample network connections
- `GET /sample/{hash}/behaviors` - Sample behaviors

#### **Discovery & Search**
- `GET /search/{expression}` - Domain search by regex
- `GET /topmillion` - Top million domains list
- `GET /subdomains/{domain}` - Subdomain enumeration

### 3. **Policies API** ⚠️ PARTIALLY IMPLEMENTED
- **Base URL**: `https://api.umbrella.com/policies/v2`
- **Current Implementation**: Destination Lists only
- **Missing Endpoints**: Need to identify via documentation crawling

### 4. **Deployments API** 🔍 NEEDS INVESTIGATION
- **Status**: Mentioned in documentation but specification not found
- **Potential endpoints**: Network deployment management

### 5. **Admin API** 🔍 NEEDS INVESTIGATION  
- **Status**: Mentioned in documentation but specification not found
- **Potential endpoints**: Organization and user management

### 6. **Reports API** 🔍 NEEDS INVESTIGATION
- **Status**: Mentioned in documentation but specification not found
- **Potential endpoints**: Analytics and reporting data

## 🎯 **Terraform Resource Mapping Analysis**

### **High Priority - Investigate API Data Sources**
These endpoints are perfect for Terraform data sources (read-only):

#### **Domain Intelligence Data Sources**
```hcl
data "umbrella_domain_categorization" "example" {
  domain = "example.com"
}

data "umbrella_domain_risk_score" "example" {
  domain = "example.com"
}

data "umbrella_domain_security_info" "example" {
  domain = "example.com"
}
```

#### **Threat Intelligence Data Sources**
```hcl
data "umbrella_domain_timeline" "example" {
  domain = "example.com"
}

data "umbrella_related_domains" "example" {
  domain = "example.com"
}

data "umbrella_passive_dns" "example" {
  domain = "example.com"
}
```

#### **Network Intelligence Data Sources**
```hcl
data "umbrella_bgp_info" "example" {
  ip = "1.2.3.4"
}

data "umbrella_whois" "example" {
  domain = "example.com"
}
```

### **Medium Priority - Management Resources**
Need to find Admin/Deployments API specs for:

#### **Organization Management**
```hcl
resource "umbrella_organization" "example" {
  name = "My Organization"
  # Additional configuration
}

resource "umbrella_user" "example" {
  email = "user@example.com"
  role  = "admin"
}
```

#### **Network Deployments**
```hcl
resource "umbrella_network_device" "example" {
  name = "Branch Office Router"
  ip   = "203.0.113.1"
}

resource "umbrella_site" "example" {
  name = "Branch Office"
  # Network configuration
}
```

## 🔧 **Technical Implementation Recommendations**

### **Phase 1: Investigate API Data Sources**
1. **Implement read-only data sources** for threat intelligence
2. **Focus on most valuable endpoints**:
   - Domain categorization and risk scoring
   - Passive DNS lookups
   - WHOIS information
   - Security timeline data

### **Phase 2: Complete Policies API**
1. **Find missing Policies API specification**
2. **Implement missing policy management resources**
3. **Add comprehensive policy rule management**

### **Phase 3: Admin & Deployments APIs**
1. **Locate Admin API specification**
2. **Locate Deployments API specification**  
3. **Implement organization and user management**
4. **Implement network deployment resources**

### **Phase 4: Reports API**
1. **Locate Reports API specification**
2. **Implement reporting data sources**
3. **Add analytics and metrics data sources**

## 📊 **API Specifications Found**

| API | Specification | Size | Status |
|-----|---------------|------|--------|
| Authentication | `api-specs/auth_api.yaml` | 2.9KB | ✅ Downloaded |
| Investigate | `api-specs/investigate_api.yaml` | 137KB | ✅ Downloaded |
| Policies | Not found | - | 🔍 Need to locate |
| Deployments | Not found | - | 🔍 Need to locate |
| Admin | Not found | - | 🔍 Need to locate |
| Reports | Not found | - | 🔍 Need to locate |

## 🚀 **Next Steps**

### **Immediate Actions**
1. **Analyze Investigate API** for data source implementation
2. **Find remaining API specifications** through:
   - Direct URL pattern testing
   - Documentation deep-dive
   - Cisco developer portal exploration

### **Development Priorities**
1. **High Value Data Sources**: Domain intelligence, threat data
2. **Complete Policy Management**: Missing policy endpoints
3. **Organization Management**: Admin API implementation
4. **Network Management**: Deployments API implementation

### **Rate Limiting Considerations**
- **Investigate API**: Requires `investigate.investigate:read` and `investigate.bulk:read` scopes
- **Bulk operations**: Limited to 1000 domains per request, 100KB payload limit
- **Passive DNS**: Maximum 10,000 records per request
- **WHOIS**: Limited to 500 results for email/nameserver searches

## 🔐 **Authentication & Scopes**

### **Required OAuth2 Scopes**
- `investigate.investigate:read` - Standard investigate operations
- `investigate.bulk:read` - Bulk investigate operations
- Additional scopes needed for other APIs (TBD)

### **Security Considerations**
- All APIs use OAuth2 client credentials flow
- Tokens expire and require refresh
- Rate limiting applies to all endpoints
- Some features require additional licenses (e.g., Cisco Secure Malware Analytics)

## 📈 **Business Value Assessment**

### **High Value Additions**
1. **Threat Intelligence Data Sources** - Enable security automation
2. **Domain Risk Assessment** - Automated security scoring
3. **Passive DNS Lookups** - Infrastructure intelligence
4. **WHOIS Data** - Domain ownership tracking

### **Medium Value Additions**
1. **Organization Management** - Complete administrative control
2. **User Management** - Access control automation
3. **Network Deployment** - Infrastructure as code
4. **Reporting Data** - Analytics and metrics

This comprehensive analysis provides a roadmap for expanding the Terraform provider to cover the full Cisco Umbrella API ecosystem, with clear priorities and implementation guidance.