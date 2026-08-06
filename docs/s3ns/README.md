<p align="center">
  <img src="img/s3ns_logo.png" alt="S3NS Sovereign Cloud Logo" width="300" />
</p>

# S3NS KCC — Sovereign Cloud Documentation Suite

Welcome to the documentation suite for **KCC in S3NS TPC (Trusted Private Cloud)**, hosted in [`fkc1e100/kcc-s3ns`](https://github.com/fkc1e100/kcc-s3ns).

This fork provides a slimmed, sovereign-compliant version of Google Cloud's Kubernetes Config Connector (KCC) tailored for S3NS TPC datacenters in France.

---

## 📚 Documentation Index

1. **[Architecture Specification](architecture.md)**
   - Technical breakdown of `s3nsapis.fr` universe domain routing.
   - Workload Identity domain integration (`s3ns.svc.id.goog`).
   - Validating Admission Webhook for GKE Autopilot & machine constraints.

2. **[Supported Services & Service Catalog](supported-services.md)**
   - Official 40 CRD allowlist (Compute, GKE Autopilot, Storage, SQL, BigQuery, IAM, Secret Manager).
   - Exclusion matrix for unsupported GCP services.

3. **[Upstream Sync & Maintenance Guide](upstream-sync-guide.md)**
   - Re-syncing upstream releases (`v1.155.0+`) via `scripts/sync-upstream.sh`.
   - Private container registry mirroring & offline deployment.

4. **[Developer & Testing Guide](dev-guide.md)**
   - Guide to developing Direct Controllers for S3NS.
   - Running `mockgcp` integration tests with URL normalization.
   - Using `scripts/prune-s3ns-crds.sh` to keep the bundle slim.

---

## 🛠️ Upstream Developer Resources

For general KCC direct controller development standards, refer to:
- [`docs/develop-resources/`](../develop-resources/README.md) — Direct Controller Development Guide
- [`docs/dev/`](../dev/) — Local Development & Testing Workflow
