provider "google" {
  credentials = "${file("gcp-cred.json")}"
  project     = "${var.gcloud-project}"
  region      = "${var.gcloud-region}"
}