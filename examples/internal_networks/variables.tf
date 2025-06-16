variable "umbrella_api_key" {
  description = "Cisco Umbrella API Key"
  type        = string
  sensitive   = true
}

variable "umbrella_api_secret" {
  description = "Cisco Umbrella API Secret"
  type        = string
  sensitive   = true
}

variable "umbrella_org_id" {
  description = "Cisco Umbrella Organization ID"
  type        = string
}

variable "tunnel_id" {
  description = "ID of an existing tunnel for tunnel-based internal network"
  type        = number
  default     = null
}