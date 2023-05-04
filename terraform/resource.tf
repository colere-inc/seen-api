# https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/cloud_run_v2_service
resource "google_cloud_run_v2_service" "seen_api_service" {
  name     = "seen-api"
  location = var.region
  ingress  = "INGRESS_TRAFFIC_ALL"

  template {
    scaling {
      min_instance_count = 1
      max_instance_count = 8
    }

    containers {
      image = var.image
      env {
        name = "FREEE_COMPANY_ID"
        value = var.freee_company_id
      }
    }
  }
}
