terraform {
  backend "gcs" {
  }
  required_version = ">= 0.12.0"
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "5.7.0"
    }
  }
}

provider "google" {
  # project: GOOGLE_PROJECT を読み込んでいる
  # credentials: GOOGLE_CREDENTIALS を読み込んでいる
  region = "asia-northeast1"
}
