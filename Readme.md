## Getting Started

To get started with the Kubebuilder-Operator, follow the steps below:

1. Clone the repository from GitHub:

   ```bash
   git clone https://github.com/YashPimple/Kubebuilder-Operator.git
   cd Kubebuilder-Operator

2. Generate the Kubernetes manifests:

   ```bash
   make manifests

3. Apply the Custom Resource Definition (CRD):

   ```bash
   kubectl apply -f config/crd/bases/demo.demo.kubebuilder.io_demovolumes.yaml

4. Install Instances of Custom Resources
   ```bash
   kubectl apply -f config/samples/demo_v1_demovolumes.yaml

5.  ```bash
    kubectl get crds

6. Run your operator locally:
   ```bash
   make run

7. Finally, check the created PVC
   ```bash
   kubectl get pvc -w
