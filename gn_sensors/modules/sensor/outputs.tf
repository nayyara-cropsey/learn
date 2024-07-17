output "sensor_id" {
  description = "Sensor ID information"
  value       = "Please see /opt/greynoise/sensor.id on the server"
}

output "ssh_port" {
  description = "Custom SSH port assigned to server"
  value       = random_integer.custom_ssh_port.result
}
