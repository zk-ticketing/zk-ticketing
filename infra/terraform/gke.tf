# Google recommends custom service accounts that have cloud-platform scope and
# permissions granted via IAM Roles.
resource "google_service_account" "gke_node" {
  account_id   = "gke-node-sa"
  display_name = "GKE Node Service Account"
}

resource "google_container_cluster" "main" {
  name     = var.gke.cluster_name
  location = var.region

  # We can't create a cluster with no node pool defined, but we want to only use
  # separately managed node pools. So we create the smallest possible default
  # node pool and immediately delete it.
  remove_default_node_pool = true
  initial_node_count       = 1

  # IMPORTANT: Creates a VPC native cluster instead of Route based cluster. Otherwise NEG won't work.
  networking_mode = "VPC_NATIVE"

  network    = google_compute_network.main.name
  subnetwork = google_compute_subnetwork.us_west1.name

  gateway_api_config {
    # Enable gateway API so that we can create LB with Gateway resource annotations.
    channel = "CHANNEL_STANDARD"
  }

  logging_config {
    # Supported values include: SYSTEM_COMPONENTS, APISERVER, CONTROLLER_MANAGER, SCHEDULER, and WORKLOADS.
    enable_components = []
  }

  monitoring_config {
    # Supported values include: SYSTEM_COMPONENTS, APISERVER, CONTROLLER_MANAGER, and SCHEDULER
    enable_components = []
    managed_prometheus {
      enabled = false
    }
  }

  #   # GSA requires workload identity. e.g. external DNS
  #   workload_identity_config {
  #     workload_pool = "${google_project.main.project_id}.svc.id.goog"
  #   }
}

resource "google_container_node_pool" "main" {
  name           = "main"
  location       = google_container_cluster.main.location
  node_locations = var.gke.node_locations
  cluster        = google_container_cluster.main.name

  autoscaling {
    min_node_count = var.gke.min_node_count
    max_node_count = var.gke.max_node_count
  }

  initial_node_count = var.gke.initial_node_count

  node_config {
    oauth_scopes = [
      #   "https://www.googleapis.com/auth/logging.write",
      #   "https://www.googleapis.com/auth/monitoring",
      "https://www.googleapis.com/auth/cloud-platform",
      "https://www.googleapis.com/auth/devstorage.read_only",
    ]

    service_account = google_service_account.gke_node.email

    machine_type = var.gke.machine_type
    tags         = ["gke-node"]
    metadata = {
      disable-legacy-endpoints = "true"
    }
  }
}

provider "kubernetes" {
  host                   = "https://${google_container_cluster.main.endpoint}"
  cluster_ca_certificate = base64decode(google_container_cluster.main.master_auth.0.cluster_ca_certificate)
  token                  = data.google_client_config.current.access_token
}

resource "kubernetes_namespace" "app" {
  provider = kubernetes
  metadata {
    name = "app"
  }
}
