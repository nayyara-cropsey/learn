terraform {
  source = "../modules/sensor"
}

inputs = {
  ssh_connection = {
    host = "3.237.66.158"
    port = 22
  }

  ssh_credentials = {
    username    = "ubuntu"
    private_key = "${file("~/.ssh/nayyara-sensors.pem")}"
  }

  api_key      = "ME36XYic3QsNcCjpB1nS6mlxP89NFSdD6RvC1tF9BeMmwYiHiL7LngVOSFwU6haB"
  workspace_id = "75a76a71-5cc1-492c-a8b7-b546bd4959ae"
  persona_id   = "c0c51efb-0ec7-44c6-ab0a-1a97f44fc5cf"
}
