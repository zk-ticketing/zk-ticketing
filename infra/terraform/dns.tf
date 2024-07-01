resource "google_dns_managed_zone" "main" {
  name     = "proofpass-io"
  dns_name = "proofpass.io."

  visibility = "public"
}

resource "google_dns_record_set" "root" {
  managed_zone = google_dns_managed_zone.main.name
  name         = "proofpass.io."
  type         = "A"
  ttl          = 300
  rrdatas      = [google_compute_global_address.frontend.address]
}

resource "google_dns_record_set" "api" {
  managed_zone = google_dns_managed_zone.main.name
  name         = "api.proofpass.io."
  type         = "A"
  ttl          = 300
  rrdatas      = [google_compute_global_address.api.address]
}
