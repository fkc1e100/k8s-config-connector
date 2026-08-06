# S3NS KCC — End-User Customer Deployment Guide (GKE Autopilot)

This guide provides step-by-step instructions for S3NS customer administrators deploying **S3NS Sovereign KCC** onto a **GKE Autopilot cluster** in S3NS TPC (Trusted Private Cloud).

---

## 1. Prerequisites

Before starting deployment, ensure you have:
1. An active **GKE Autopilot cluster** running in your S3NS TPC environment (`europe-west9` / S3NS region).
2. `kubectl` authenticated with `cluster-admin` privileges on the target Autopilot cluster.
3. An IAM Service Account (`sa-kcc-controller@YOUR_PROJECT_ID.iam.gserviceaccount.com`) with S3NS IAM roles (e.g. `roles/owner` or specific service admin roles).

---

## 2. Workload Identity Pairing for S3NS

In S3NS TPC, Workload Identity uses the **`s3ns.svc.id.goog`** federated pool suffix.

Execute the following commands to bind the Kubernetes Service Account (`cnrm-controller-manager` in `kcc-system`) to your GCP IAM Service Account:

```bash
# Set your environment variables
export PROJECT_ID="your-s3ns-project-id"
export GSA_NAME="sa-kcc-controller"
export KSA_NAME="cnrm-controller-manager"
export NAMESPACE="kcc-system"

# Bind the Kubernetes SA to the IAM SA in S3NS Workload Identity Pool
gcloud iam service-accounts add-iam-policy-binding \
  "${GSA_NAME}@${PROJECT_ID}.iam.gserviceaccount.com" \
  --role="roles/iam.workloadIdentityUser" \
  --member="principal://iam.googleapis.com/projects/${PROJECT_NUMBER}/locations/global/workloadIdentityPools/${PROJECT_ID}.s3ns.svc.id.goog/subject/ns/${NAMESPACE}/sa/${KSA_NAME}"
```

---

## 3. Installation Options

### Option A: Installing via Official Release Bundle (`s3ns-kcc-release-bundle-*.tar.gz`)

For production environments, download the pre-packaged release bundle from [GitHub Releases](https://github.com/fkc1e100/kcc-s3ns/releases):

```bash
# Download and extract official S3NS release bundle
curl -L -O https://github.com/fkc1e100/kcc-s3ns/releases/download/v1.154.1-s3ns/s3ns-kcc-release-bundle-v1.154.1.tar.gz
mkdir -p s3ns-kcc && tar -xzvf s3ns-kcc-release-bundle-v1.154.1.tar.gz -C s3ns-kcc
cd s3ns-kcc

# 1. Apply the 40 slimmed S3NS CRD manifests
kubectl apply -f crds.yaml

# 2. Deploy cnrm-controller-manager with S3NS Autopilot Workload Identity
kubectl apply -f install-bundle-autopilot-workload-identity/0-cnrm-system.yaml
```

---

### Option B: Installing from Repository Checkout

If deploying directly from the source repository:

```bash
# Clone the repository
git clone https://github.com/fkc1e100/kcc-s3ns.git
cd kcc-s3ns

# Run the S3NS Autopilot deployment script
./scripts/deploy-s3ns-kcc.sh
```

---

## 4. Configuring Namespaced Mode (Multi-Tenant Isolation)

For multi-tenant S3NS customer deployments, configure **Namespaced Mode** so each tenant team reconciles resources exclusively in their authorized GCP project:

```yaml
apiVersion: core.cnrm.cloud.google.com/v1beta1
kind: ConfigConnectorContext
metadata:
  name: configconnectorcontext.core.cnrm.cloud.google.com
  namespace: team-a
spec:
  googleServiceAccount: "sa-team-a@your-s3ns-project-id.iam.gserviceaccount.com"
```

---

## 5. Deployment Verification

Verify that the controller manager pod is running and communicating with `s3nsapis.fr`:

```bash
# 1. Check pod status
kubectl get pods -n kcc-system

# 2. Inspect logs for s3nsapis.fr endpoint initialization
kubectl logs -n kcc-system deployment/cnrm-controller-manager -c manager | grep "s3nsapis.fr"
```
