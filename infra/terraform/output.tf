output "dns_name_servers" {
  value = google_dns_managed_zone.main.name_servers
}

output "frontend_ip" {
  value = google_compute_global_address.frontend.address
}

output "api_ip" {
  value = google_compute_global_address.api.address
}
