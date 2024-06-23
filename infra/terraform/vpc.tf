resource "google_compute_network" "main" {
  name                    = "main"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "us_west1" {
  name          = "us-west-1"
  region        = "us-west1"
  network       = google_compute_network.main.name
  ip_cidr_range = local.vpc_subnetwork_ip_cidr_range
}
