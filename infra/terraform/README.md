# Terraform

Terraform state data is stored in a remote gcs bucket. The bucket is versioned and encrypted. Bucket name is terraform-proof-pass.

## Components

### GKE

Main application is deployed on GKE.

## Setup

Create file terraform.tfvars to provide values for the variables in main.tf. The file should look like this:

```
gcp_common = {
  org_id          = "123"
  billing_account = "123-456-7890"
  folder_id       = "123"
}

tf_backend_bucket_name = "bucket"

<everything else you want to pass as optional variables>
```

To setup a brand new terraform project, follow these steps:

1. Comment out the terraform backend config block in the main.tf file.
2. `terraform init` to initialize the terraform project.
3. `terraform apply` to create the remote state bucket and all other resources.
4. Uncomment the terraform backend config block in the main.tf file.
5. `terraform init` to reinitialize the terraform project with the remote state backend. This will move the state file to the remote bucket, and all future terraform commands will use the remote state.

To join an existing terraform project that uses the remote state backend, follow these steps:

1. Ask the project owner to provide you with access to the remote state bucket.
2. `terraform init` to initialize the terraform project with the remote state backend.
