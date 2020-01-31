package cloud

import (
	"sort"
	"strings"
)

const (
	GKE        = "gke"
	OKE        = "oke"
	EKS        = "eks"
	AKS        = "aks"
	AWS        = "aws"
	PKS        = "pks"
	IKS        = "iks"
	MINIKUBE   = "minikube"
	MINISHIFT  = "minishift"
	KUBERNETES = "kubernetes"
	OPENSHIFT  = "openshift"
	ICP        = "icp"
	JX_INFRA   = "jx-infra"
	ALIBABA    = "alibaba"
)

// KubernetesProviders list of all available Kubernetes providers
var KubernetesProviders = []string{ALIBABA, AKS, AWS, EKS, GKE, OKE, KUBERNETES, ICP, IKS, OPENSHIFT, MINISHIFT, MINIKUBE, JX_INFRA, PKS}

// KubernetesProviderOptions returns all the Kubernetes providers as a string
func KubernetesProviderOptions() string {
	values := []string{}
	values = append(values, KubernetesProviders...)
	sort.Strings(values)
	return strings.Join(values, ", ")
}
