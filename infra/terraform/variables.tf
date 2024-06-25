variable "gcp_common" {
  type = object({
    org_id          = string
    billing_account = string
    folder_id       = string
  })
}

variable "region" {
  type    = string
  default = "us-west1"
}

variable "tf_backend_bucket_name" {
  type    = string
}

variable "gke" {
  type = object({
    cluster_name       = string
    machine_type       = string
    node_locations     = list(string)
    min_node_count     = number
    max_node_count     = number
    initial_node_count = number
  })
  default = {
    cluster_name       = "app"
    machine_type       = "t2d-standard-2"
    node_locations     = ["us-west1-a"]
    min_node_count     = 1
    max_node_count     = 4
    initial_node_count = 1
  }
}
