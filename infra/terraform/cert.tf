# Create a DNS authorization
resource "google_certificate_manager_dns_authorization" "main" {
  name   = "root"
  domain = local.domain
}

resource "google_certificate_manager_dns_authorization" "api" {
  name   = "api"
  domain = "api.${local.domain}"
}

# Add CNAME record to DNS configuration
resource "google_dns_record_set" "root_cert" {
  name         = google_certificate_manager_dns_authorization.main.dns_resource_record[0].name
  managed_zone = google_dns_managed_zone.main.name
  type         = google_certificate_manager_dns_authorization.main.dns_resource_record[0].type
  ttl          = 300
  rrdatas      = [google_certificate_manager_dns_authorization.main.dns_resource_record[0].data]
}

resource "google_dns_record_set" "api_cert" {
  name         = google_certificate_manager_dns_authorization.api.dns_resource_record[0].name
  managed_zone = google_dns_managed_zone.main.name
  type         = google_certificate_manager_dns_authorization.api.dns_resource_record[0].type
  ttl          = 300
  rrdatas      = [google_certificate_manager_dns_authorization.api.dns_resource_record[0].data]
}

# Create a Google-managed certificate referencing the DNS authorization
resource "google_certificate_manager_certificate" "main" {
  name = "app"
  managed {
    domains = [local.domain, "api.${local.domain}"]
    dns_authorizations = [
      google_certificate_manager_dns_authorization.main.id,
      google_certificate_manager_dns_authorization.api.id
    ]
  }
}

resource "google_certificate_manager_certificate_map" "main" {
  name = "app-certmap"
}

resource "google_certificate_manager_certificate_map_entry" "root" {
  name         = "app-certmap-entry-root"
  map          = google_certificate_manager_certificate_map.main.name
  hostname     = local.domain
  certificates = [google_certificate_manager_certificate.main.id]
}

resource "google_certificate_manager_certificate_map_entry" "api" {
  name         = "app-certmap-entry-api"
  map          = google_certificate_manager_certificate_map.main.name
  hostname     = "api.${local.domain}"
  certificates = [google_certificate_manager_certificate.main.id]
}
