# Persist state data to keep track of the resources terraform manages.
resource "google_storage_bucket" "tf_backend" {
  name          = var.tf_backend_bucket_name
  location      = "US"
  force_destroy = true

  versioning {
    enabled = true
  }
}
