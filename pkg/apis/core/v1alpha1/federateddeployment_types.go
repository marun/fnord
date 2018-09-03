/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// FederatedDeploymentSpec defines the desired state of FederatedDeployment
type FederatedDeploymentSpec struct {
	Template appsv1.Deployment `json:"template,omitempty"`
}

// FederatedDeploymentStatus defines the observed state of FederatedDeployment
type FederatedDeploymentStatus struct {
	ClusterStatuses []FederatedDeploymentClusterStatus `json:"clusterStatuses,omitempty"`
}

// FederatedDeploymentClusterStatus is the observed status for a named cluster
type FederatedDeploymentClusterStatus struct {
	ClusterName string                  `json:"clusterName,omitempty"`
	Status      appsv1.DeploymentStatus `json:"status,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FederatedDeployment
// +k8s:openapi-gen=true
// +kubebuilder:resource:path=federateddeployments
type FederatedDeployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FederatedDeploymentSpec   `json:"spec,omitempty"`
	Status FederatedDeploymentStatus `json:"status,omitempty"`
}
