# Creates a network layer
resource "google_compute_network" "neocargo-network" {
  name = "${var.platform-name}"
}

# Creates a firewall with some sane defaults, allowing ports 22, 80 and 443 to be open
# This is ssh, http and https.
resource "google_compute_firewall" "ssh" {
  name    = "${var.platform-name}-ssh"
  network = "${google_compute_network.neocargo-network.name}"

  allow {
    protocol = "icmp"
  }

  allow {
    protocol = "tcp"
    ports    = ["22", "80", "443"]
  }

   source_ranges = ["0.0.0.0/0"]
}

# Creates a new DNS zone
resource "google_dns_managed_zone" "neocargo" {
  name        = "neocargo-com"
  dns_name    = "neocargo.com."
  description = "neocargo.com DNS zone"
}

# Creates a new subnet for our platform within our selected region
resource "google_compute_subnetwork" "neocargo" {
  name          = "dev-${var.platform-name}-${var.gcloud-region}"
  ip_cidr_range = "10.1.2.0/24"
  network       = "${google_compute_network.neocargo-network.self_link}"
  region        = "${var.gcloud-region}"
}

# Creates a container cluster called 'neocargo-cluster'
# Attaches new cluster to our network and our subnet,
# Ensures at least one instance is running
resource "google_container_cluster" "neocargo-cluster" {
  name = "neocargo-cluster"
  network = "${google_compute_network.neocargo-network.name}"
  subnetwork = "${google_compute_subnetwork.neocargo.name}"
  zone = "${var.gcloud-zone}"

  initial_node_count = 1

  master_auth {
    username = <redacted>
    password = <redacted>
  }

  node_config {

    # Defines the type/size instance to use
    # Standard is a sensible starting point
    machine_type = "n1-standard-1"

    # Grants OAuth access to the following API's within the cluster
    oauth_scopes = [
      "https://www.googleapis.com/auth/projecthosting",
      "https://www.googleapis.com/auth/devstorage.full_control",
      "https://www.googleapis.com/auth/monitoring",
      "https://www.googleapis.com/auth/logging.write",
      "https://www.googleapis.com/auth/compute",
      "https://www.googleapis.com/auth/cloud-platform"
    ]
  }
}

# Creates a new DNS range for cluster
resource "google_dns_record_set" "dev-k8s-endpoint-neocargo" {
  name  = "k8s.dev.${google_dns_managed_zone.neocargo.dns_name}"
  type  = "A"
  ttl   = 300

  managed_zone = "${google_dns_managed_zone.neocargo.name}"

  rrdatas = ["${google_container_cluster.neocargo-cluster.endpoint}"]
}