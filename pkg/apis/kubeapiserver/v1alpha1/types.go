package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"

	operatorsv1alpha1api "github.com/openshift/api/operator/v1alpha1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubeAPISConfig provides information to configure kube-apiserver
type KubeAPIServerConfig struct {
	metav1.TypeMeta `json:",inline"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubeAPISOperatorConfig provides information to configure an operator to manage kube-apiserver.
type KubeAPIServerOperatorConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	Spec   KubeAPIServerOperatorConfigSpec   `json:"spec"`
	Status KubeAPIServerOperatorConfigStatus `json:"status"`
}

type KubeAPIServerOperatorConfigSpec struct {
	operatorsv1alpha1api.OperatorSpec `json:",inline"`

	// forceRedeploymentReason can be used to force the redeployment of the kube-apiserver by providing a unique string.
	// This provides a mechanism to kick a previously failed deployment and provide a reason why you think it will work
	// this time instead of failing again on the same config.
	ForceRedeploymentReason string `json:"forceRedeploymentReason"`

	// userConfig holds a sparse config that the user wants for this component.  It only needs to be the overrides from the defaults
	// it will end up overlaying in the following order:
	// 1. hardcoded default
	// 2. this config
	UserConfig runtime.RawExtension `json:"userConfig"`

	// observedConfig holds a sparse config that controller has observed from the cluster state.  It exists in spec because
	// it causes action for the operator
	ObservedConfig runtime.RawExtension `json:"observedConfig"`
}

type KubeAPIServerOperatorConfigStatus struct {
	operatorsv1alpha1api.StaticPodOperatorStatus `json:",inline"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubeAPISOperatorConfigList is a collection of items
type KubeAPIServerOperatorConfigList struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items contains the items
	Items []KubeAPIServerOperatorConfig `json:"items"`
}
