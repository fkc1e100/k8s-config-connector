<p align="center">
  <img src="docs/s3ns/img/s3ns_logo.png" alt="S3NS Sovereign Cloud Logo" width="360" />
</p>

# S3NS Sovereign Kubernetes Config Connector (`kcc-s3ns`)

[![s3ns-tpc-ci](https://github.com/fkc1e100/kcc-s3ns/actions/workflows/s3ns-ci.yaml/badge.svg)](https://github.com/fkc1e100/kcc-s3ns/actions/workflows/s3ns-ci.yaml)
[![Baseline](https://img.shields.io/badge/KCC_Baseline-v1.154.1-blue.svg)](https://github.com/GoogleCloudPlatform/k8s-config-connector/releases/tag/v1.154.1)
[![Compliance](https://img.shields.io/badge/Compliance-SecNumCloud_3.2-green.svg)](https://www.s3ns.io)

This repository ([`fkc1e100/kcc-s3ns`](https://github.com/fkc1e100/kcc-s3ns)) is a slimmed, sovereign-compliant fork of Google Cloud's **Kubernetes Config Connector (KCC)**. It is engineered specifically for **S3NS Trusted Private Cloud (TPC)** datacenters in France and is architected to support international TPC environments (e.g. Germany `tpc.de`, Canada `tpc.ca`, Japan `tpc.jp`).

---

## 🏛️ Upstream Foundation & Sovereign Fork Rationale

### Why This Fork Exists
Upstream Google Cloud KCC is designed for global GCP (`googleapis.com`) and bundles **464 Custom Resource Definitions (CRDs)** covering the full range of GCP services. 

In sovereign TPC environments such as S3NS:
1. **Endpoint Isolation**: API calls must route to isolated sovereign domains (`*.s3nsapis.fr`) rather than `googleapis.com`.
2. **Workload Identity Domain**: Identity federations use custom pool domain suffixes (`s3ns.svc.id.goog`).
3. **Pruned Service Catalog**: Only ANSSI SecNumCloud approved services are available in TPC datacenters. Running unneeded upstream reconcilers introduces security risk and control plane memory overhead.

### Upstream Alignment
This fork is built directly on top of Google Cloud's official release **[`v1.154.1`](https://github.com/GoogleCloudPlatform/k8s-config-connector/releases/tag/v1.154.1)**. Upstream updates, security patches, and new Direct Controllers are continuously synchronized via [`scripts/sync-upstream.sh`](scripts/sync-upstream.sh).

---

## 🌟 Sovereign Architectural Highlights

1. **Dynamic TPC Universe Domain (`s3nsapis.fr`)**:
   - Automatically routes REST and gRPC API calls to sovereign endpoints (`storage.s3nsapis.fr`, `compute.s3nsapis.fr`, `sqladmin.s3nsapis.fr`).
   - Driven by the generic `pkg/universe` module and configured via `GOOGLE_CLOUD_UNIVERSE_DOMAIN`.

2. **Workload Identity Domain Suffix (`s3ns.svc.id.goog`)**:
   - Configures federated service account authentication using `WORKLOAD_IDENTITY_DOMAIN`.

3. **Slimmed 40-CRD Service Catalog**:
   - Pruned down from 464 upstream CRDs to **40 allowed S3NS resource kinds** (GKE Autopilot, Compute Engine, Cloud Storage, Cloud SQL, BigQuery, IAM, Secret Manager, Vertex AI).
   - Enforces strict runtime allowlisting in `pkg/controller/registration/`.

4. **80%+ Repository Slimming**:
   - Purged 225 unsupported sample directories, 85 mock GCP packages, and obsolete experimental code for maximum maintainability.

---

## 🚀 Quickstart Deployment

Deploy the slimmed S3NS KCC bundle onto a GKE Autopilot cluster:

```bash
# Clone the repository
git clone https://github.com/fkc1e100/kcc-s3ns.git
cd kcc-s3ns

# Deploy 40 S3NS CRDs & cnrm-controller-manager configured for s3nsapis.fr
./scripts/deploy-s3ns-kcc.sh
```

---

## 📚 S3NS Documentation Suite

Explore detailed guides in the [`docs/s3ns/`](docs/s3ns/README.md) directory:

| Guide | Description | Target Audience |
| :--- | :--- | :--- |
| 📖 **[S3NS Overview & Index](docs/s3ns/README.md)** | Main entry point for S3NS documentation | All Engineers |
| 📋 **[Supported Service Catalog](docs/s3ns/supported-services.md)** | 40-CRD allowlist and GCP service exclusion matrix | Architects & DevOps |
| 🏗️ **[Technical Architecture](docs/s3ns/architecture.md)** | Endpoint resolution, universe domain logic & Workload Identity | Core Developers |
| 🚀 **[Autopilot Deployment Guide](docs/s3ns/deployment.md)** | Step-by-step GKE Autopilot deployment instructions | S3NS Customers |
| 🔄 **[Upstream Sync Guide](docs/s3ns/upstream-sync-guide.md)** | Re-syncing upstream Google releases (`v1.155.0+`) | Maintenance Team |

---

## 🤝 Handover & S3NS Tracking Issues

7 tracking issues are logged on GitHub for the S3NS engineering team to execute upon handover:

* **[Issue #10](https://github.com/fkc1e100/kcc-s3ns/issues/10)**: `refactor`: Migrate GKE `ContainerCluster` & `ContainerNodePool` to Direct Controller.
* **[Issue #11](https://github.com/fkc1e100/kcc-s3ns/issues/11)**: `refactor`: Migrate `Project` & `Folder` to Direct Controller.
* **[Issue #12](https://github.com/fkc1e100/kcc-s3ns/issues/12)**: `refactor`: Migrate Private Service Connect to Direct Controller.
* **[Issue #13](https://github.com/fkc1e100/kcc-s3ns/issues/13)**: `ci`: Configure Nightly E2E Integration Test Pipeline on S3NS Autopilot.
* **[Issue #14](https://github.com/fkc1e100/kcc-s3ns/issues/14)**: `docs`: Establish S3NS Developer Workstation Onboarding & Local Emulation Guide.
* **[Issue #15](https://github.com/fkc1e100/kcc-s3ns/issues/15)**: `ops`: Automate Container Image Builds & Sync to `registry.s3ns.fr`.
* **[Issue #16](https://github.com/fkc1e100/kcc-s3ns/issues/16)**: `security`: Deploy Validating Admission Webhook for Autopilot & Machine Constraints.
