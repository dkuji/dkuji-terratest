provider "google" {
  credentials = "${file("./credentials/dkuji-cloud-run-55573a6f90d4.json")}"
  project     = "dkuji-cloud-run"
  region      = "us-central1"
  version     = "v3.1.0"
}
