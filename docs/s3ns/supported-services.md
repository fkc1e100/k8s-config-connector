# S3NS TPC Supported Services & GVK Catalog

This document defines the **official 40 CRD allowlist** supported by the `fkc1e100/kcc-s3ns` fork in S3NS Trusted Private Cloud (TPC).

---

## 1. Supported Services Table

| Category | GCP Service | KCC Resource Kind (`GVK`) | API Domain | Controller Engine |
| :--- | :--- | :--- | :--- | :--- |
| **Compute** | GKE Autopilot | `ContainerCluster`, `ContainerNodePool` | `container.s3nsapis.fr` | TF2CRD (Issue #10) |
| **Compute** | Compute Engine | `ComputeInstance`, `ComputeDisk` | `compute.s3nsapis.fr` | Direct |
| **Networking** | VPC & Subnets | `ComputeNetwork`, `ComputeSubnetwork` | `compute.s3nsapis.fr` | Direct |
| **Networking** | Firewall Rules | `ComputeFirewall` | `compute.s3nsapis.fr` | Direct |
| **Networking** | Cloud Router & NAT | `ComputeRouter`, `ComputeRouterNAT` | `compute.s3nsapis.fr` | Direct |
| **Networking** | Static IPs | `ComputeAddress` | `compute.s3nsapis.fr` | Direct |
| **Networking** | Regional LB | `ComputeForwardingRule`, `ComputeBackendService` | `compute.s3nsapis.fr` | Direct |
| **Storage** | Cloud Storage | `StorageBucket` | `storage.s3nsapis.fr` | Direct |
| **Database** | Cloud SQL | `SQLInstance`, `SQLDatabase`, `SQLUser` | `sqladmin.s3nsapis.fr` | Direct |
| **Analytics** | BigQuery | `BigQueryDataset`, `BigQueryTable` | `bigquery.s3nsapis.fr` | Direct |
| **Hierarchy** | Resource Manager | `Project`, `Folder` | `cloudresourcemanager.s3nsapis.fr` | TF2CRD (Issue #11) |
| **IAM** | Identity & Access | `IAMPolicy`, `IAMPartialPolicy`, `IAMPolicyMember`, `IAMServiceAccount` | `iam.s3nsapis.fr` | Direct / Native |
| **Security** | Secret Manager | `SecretManagerSecret`, `SecretManagerSecretVersion` | `secretmanager.s3nsapis.fr` | Direct |
| **AI / ML** | Vertex AI (A3 GPU) | `VertexAIFeaturestore`, `AIPlatformModel` | `aiplatform.s3nsapis.fr` | Direct |

---

## 2. Exclusion Matrix (Unsupported Resources)

The following resources are excluded from S3NS TPC:

- `BinaryAuthorizationPolicy` (Binary Authorization)
- `TPUNode` (Cloud TPU Acceleration)
- `ConnectGateway` (GKE Enterprise Connect Relay)
- `PolicyController` (GKE Enterprise Policy Controller)
- `ContainerCluster` with `spec.autopilot.enabled == false` (Standard GKE Mode)
- Global `ComputeBackendService` (Anycast Load Balancing)
