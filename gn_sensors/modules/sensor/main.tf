terraform {
  required_providers {
    null = {
      source  = "hashicorp/null"
      version = ">= 3.0.0"
    }

    random = {
      source  = "hashicorp/random"
      version = ">= 3.3.0"
    }
  }
}

locals {
  public_ip = var.public_ip == null ? var.ssh_connection.host : var.public_ip

  bootstrap_cmd = templatefile("${path.module}/bootstrap.tftpl", {
    api_key         = var.api_key
    workspace_id    = var.workspace_id
    public_ip       = local.public_ip
    custom_ssh_port = random_integer.custom_ssh_port.result
  })

  persona_deploy_cmd = templatefile("${path.module}/persona_deploy.tftpl", {
    api_key      = var.api_key
    workspace_id = var.workspace_id
    persona_id   = var.persona_id
  })
}

resource "random_integer" "custom_ssh_port" {
  min = 55000
  max = 65535

  keepers = {
    host = var.ssh_connection.host
  }
}

resource "null_resource" "bootstrap" {
  triggers = {
    workspace = var.workspace_id
  }

  connection {
    host        = var.ssh_connection.host
    port        = var.ssh_connection.port
    user        = var.ssh_credentials.username
    password    = var.ssh_credentials.password
    private_key = var.ssh_credentials.private_key
  }

  provisioner "remote-exec" {
    inline = [
      local.bootstrap_cmd,
    ]
    # failure is expected as SSH connection will be lost
    # once bootstrap completes and changes SSH port
    on_failure = continue
  }
}

resource "null_resource" "persona_deploy" {
  triggers = {
    workspace = var.workspace_id
    persona   = var.persona_id
  }

  connection {
    host        = var.ssh_connection.host
    port        = random_integer.custom_ssh_port.result
    user        = var.ssh_credentials.username
    password    = var.ssh_credentials.password
    private_key = var.ssh_credentials.private_key
  }

  provisioner "remote-exec" {
    inline = [
      local.persona_deploy_cmd,
    ]
  }

  depends_on = [
    null_resource.bootstrap,
  ]
}
