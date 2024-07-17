# --- required variables ---
variable "api_key" {
  description = "GreyNoise API key needed to deploy sensor in a given workspace"
  type        = string
  # not marked as sensitive to avoid suppressing output in remote-exec provisioner
}

variable "workspace_id" {
  description = "Workspace ID for the given workspace"
  type        = string
}

variable "persona_id" {
  description = "The ID of the persona to setup the sensor with"
  type        = string
}

variable "ssh_connection" {
  description = "SSH connection parameters"
  type        = object({
    host = string
    port = number
  })
}

variable "ssh_credentials" {
  description = "SSH credentials for the sensor server"
  type        = object({
    username    = string
    password    = optional(string)
    private_key = optional(string)
  })
  sensitive = true
}

# --- optional variables ---
variable "public_ip" {
  description = "Public IP for configuring sensor (if different from host)"
  type        = optional(string)
  default     = null
}
