# Kube Descheduler Operator

Run the descheduler in your OpenShift cluster to move pods based on specific strategies.

## Deploy the operator

1. build and push the image to a registry (e.g. https://quay.io):
   ```sh
   $ podman build -t quay.io/<username>/ose-cluster-kube-descheduler-operator-bundle:latest -f Dockerfile .
   $ podman push quay.io/<username>/ose-cluster-kube-descheduler-operator-bundle:latest
   ```

1. build and push image index for operator-registry (pull and build https://github.com/operator-framework/operator-registry/ to get the `opm` binary)
   ```sh
   $ ./bin/linux-amd64-opm index add --bundles quay.io/<username>/ose-cluster-kube-descheduler-operator-bundle:latest --tag quay.io/<username>/ose-cluster-kube-descheduler-operator-bundle-index:1.0.0
   $ podman push quay.io/<username>/ose-cluster-kube-descheduler-operator-bundle-index:1.0.0
   ```

   Don't forget to increase the number of open files, .e.g. `ulimit -n 100000` in case the current limit is insufficient.

1. create and apply catalogsource manifest:
   ```yaml
   apiVersion: operators.coreos.com/v1alpha1
   kind: CatalogSource
   metadata:
     name: cluster-kube-descheduler-operator
     namespace: openshift-marketplace
   spec:
     sourceType: grpc
     image: quay.io/<username>/ose-cluster-kube-descheduler-operator-bundle-index:1.0.0
   ```

1. create `cluster-kube-descheduler-operator` namespace:
   ```
   $ oc create ns cluster-kube-descheduler-operator
   ```

1. open the console Operators -> OperatorHub, search for `descheduler operator` and install the operator

## Descheduler strategies

The Descheduler operator attempts to simplify the descheduler strategy names from their [upstream names](https://github.com/kubernetes-sigs/descheduler/#policy-and-strategies). Thus when set on the operator, these strategy names map to:

| Operator param | Descheduler strategy |
| ---- | ---- |
| `duplicates` | `RemoveDuplicates` |
| `interpodantiaffinity` | `RemovePodsViolatingInterPodAntiAffinity` |
| `lownodeutilization` | `LowNodeUtilization` |
| `nodeaffinity` | `RemovePodsViolatingNodeAffinity` |
| `nodetaints` | `RemovePodsViolatingNodeTaints` |

## Sample CR

A sample CR definition looks like below (the operator expects `config` CR under `openshift-kube-descheduler-operator` namespace):

```yaml
apiVersion: operator.openshift.io/v1beta1
kind: KubeDescheduler
metadata:
  name: config
  namespace: openshift-kube-descheduler-operator
spec:
  deschedulingIntervalSeconds: 1800
  strategies:
    - name: "lownodeutilization"
      params:
       - name: "cputhreshold"
         value: "10"
       - name: "memorythreshold"
         value: "20"
       - name: "podsthreshold"
         value: "30"
       - name: "memorytargetthreshold"
         value: "40"
       - name: "cputargetthreshold"
         value: "50"
       - name: "podstargetthreshold"
         value: "60"
       - name: "nodes"
         value: "3"
    - name: "duplicates"
```
The valid list of strategies are "lownodeutilization", "duplicates", "interpodantiaffinity", "nodeaffinity", and "nodetaints". Out of the above only lownodeutilization has parameters like cputhreshold, memorythreshold etc. Using the above strategies defined in CR we create a configmap in openshift-descheduler-operator namespace. As of now, adding new strategies could be done through code. DeschedulingIntervalSeconds field contains the number of seconds between a descheduler run (0 in this field will only run the descheduler once and exit). Nodes field indicate on how many nodes the lownodeutilization strategy should run.

## How does the descheduler operator work?

Descheduler operator at a high level is responsible for watching the above CR
- Create a configmap that could be used by descheduler.
- Run descheduler as a deployment mounting the configmap as a policy file in the pod.

The configmap created from above sample CR definition looks like this:

```yaml
apiVersion: "kubedeschedulers.operator.openshift.io/v1beta1"
kind: "DeschedulerPolicy"
strategies:
  "RemoveDuplicates":
     enabled: true
  "LowNodeUtilization":
     enabled: true
     params:
       nodeResourceUtilizationThresholds:
         thresholds:
           "cpu" : 10
           "memory": 20
           "pods": 30
         targetThresholds:
           "cpu" : 40
           "memory": 50
           "pods": 60
         numberOfNodes: 3
```

The above configmap would be mounted as a volume in descheduler pod created. Whenever we change strategies, parameters or schedule in the CR, the descheduler operator is responsible for identifying those changes and regenerating the configmap. For more information on how descheduler works, please visit [descheduler](https://docs.openshift.com/container-platform/3.11/admin_guide/scheduling/descheduler.html)


## Parameters
The Descheduler operator exposes the following parameters in its CRD:

* `deschedulingIntervalSeconds` - this sets the number of seconds between descheduler runs
* `image` - specifies the Descheduler container image to deploy
* `flags` - this allows additional descheduler flags to be set, and they will be appended to the descheduler pod. Therefore, they must be in the same format as would be passed to the descheduler binary (eg, `"--dry-run"`)
