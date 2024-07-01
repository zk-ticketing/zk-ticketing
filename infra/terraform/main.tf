terraform {
  backend "gcs" {
    bucket = "terraform-proof-pass"
    prefix = "terraform/"
  }
}

locals {
  project_name = "Proof Pass"
  project_id   = "zk-ticketing"

  enable_services = [
    "storage-api.googleapis.com",
    "compute.googleapis.com",
    "container.googleapis.com",
    "dns.googleapis.com",
    "run.googleapis.com",
    "secretmanager.googleapis.com",
    "certificatemanager.googleapis.com"
  ]

  # VPC
  vpc_subnetwork_ip_cidr_range = "10.0.1.0/24"

  domain = "proofpass.io"
}

provider "google" {
  project = local.project_id
  region  = var.region
}

data "google_client_config" "current" {}

resource "google_project" "main" {
  name            = local.project_name
  project_id      = local.project_id
  folder_id       = var.gcp_common.folder_id
  billing_account = var.gcp_common.billing_account
  labels = {
    "terraform" : "true"
  }
}

resource "google_project_service" "main" {
  project = google_project.main.project_id

  for_each = toset(local.enable_services)
  service  = each.value
}

