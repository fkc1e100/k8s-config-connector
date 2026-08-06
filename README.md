# S3NS Sovereign Kubernetes Config Connector (`kcc-s3ns`)

[![s3ns-tpc-ci](https://github.com/fkc1e100/kcc-s3ns/actions/workflows/s3ns-ci.yaml/badge.svg)](https://github.com/fkc1e100/kcc-s3ns/actions/workflows/s3ns-ci.yaml)

This repository ([`fkc1e100/kcc-s3ns`](https://github.com/fkc1e100/kcc-s3ns)) is a slimmed, sovereign-compliant fork of Google Cloud's **Kubernetes Config Connector (KCC)**, tailored specifically for **S3NS Trusted Private Cloud (TPC)** datacenters in France and applicable to international TPC environments (Germany, Canada, Japan).

---

## 🌟 Key Sovereign Capabilities

1. **Dynamic TPC Universe Domain Resolution (`s3nsapis.fr`)**:
   - Automatically routes REST and gRPC API calls to sovereign endpoints (`storage.s3nsapis.fr`, `compute.s3nsapis.fr`, `sqladmin.s3nsapis.fr`).
   - Driven by the generic `pkg/universe` module and configured via `GOOGLE_CLOUD_UNIVERSE_DOMAIN`.

2. **Workload Identity Domain Suffix (`s3ns.svc.id.goog`)**:
   - Integrates federated identity authentication using `WORKLOAD_IDENTITY_DOMAIN`.

3. **Slimmed 40 CRD Service Catalog**:
   - Pruned down from 464 upstream CRDs to **40 allowed S3NS resource kinds** (GKE Autopilot, Compute Engine, Storage, Cloud SQL, BigQuery, IAM, Secret Manager, Vertex AI).
   - Enforces strict runtime allowlisting in `pkg/controller/registration/`.

---

## 🚀 Quickstart Deployment

To deploy the slimmed S3NS KCC bundle on a GKE cluster:

```bash
# Clone the repository
git clone https://github.com/fkc1e100/kcc-s3ns.git
cd kcc-s3ns

# Deploy 40 S3NS CRDs & cnrm-controller-manager with S3NS universe domain
./scripts/deploy-s3ns-kcc.sh
```

---

## 📚 S3NS Documentation Suite

Explore detailed guides in the [`docs/s3ns/`](docs/s3ns/README.md) directory:

- 📖 **[S3NS Documentation Overview](docs/s3ns/README.md)**
- 📋 **[Supported Service Catalog & 40 CRD Allowlist](docs/s3ns/supported-services.md)**
- 🏗️ **[Technical Architecture & Endpoint Resolution](docs/s3ns/architecture.md)**
- 🔄 **[Upstream Sync & Maintenance Guide](docs/s3ns/upstream-sync-guide.md)**
- 🛠️ **[Direct Controller Developer Guide](docs/s3ns/dev-guide.md)**

---

## 🔄 Upstream Relationship

This repository is forked from [`GoogleCloudPlatform/k8s-config-connector`](https://github.com/GoogleCloudPlatform/k8s-config-connector) (baseline release `v1.154.1`). Upstream updates are regularly synchronized via [`scripts/sync-upstream.sh`](scripts/sync-upstream.sh).
