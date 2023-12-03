# https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/cloud_run_v2_service
resource "google_cloud_run_v2_service" "default" {
  name     = "seen-api"
  location = var.region

  template {
    scaling {
      min_instance_count = 1
      max_instance_count = 8
    }
    containers {
      image = var.image
      env {
        name  = "GCP_PROJECT_ID"
        value = var.gcp_project_id
      }
    }
  }
}

resource "google_cloud_run_v2_service_iam_member" "member" {
  name       = google_cloud_run_v2_service.default.name
  location   = google_cloud_run_v2_service.default.location
  role       = "roles/run.invoker"
  member     = "allUsers"
  depends_on = [google_cloud_run_v2_service.default]
}
