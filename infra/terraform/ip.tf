resource "google_compute_global_address" "frontend" {
  name = "frontend"
}

resource "google_compute_global_address" "api" {
  name = "api"
}
