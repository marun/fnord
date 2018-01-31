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

// This file was automatically generated by lister-gen

package v1alpha1

import (
	v1alpha1 "github.com/marun/fnord/pkg/apis/federation/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// FederatedSecretOverridesLister helps list FederatedSecretOverrideses.
type FederatedSecretOverridesLister interface {
	// List lists all FederatedSecretOverrideses in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.FederatedSecretOverrides, err error)
	// FederatedSecretOverrideses returns an object that can list and get FederatedSecretOverrideses.
	FederatedSecretOverrideses(namespace string) FederatedSecretOverridesNamespaceLister
	FederatedSecretOverridesListerExpansion
}

// federatedSecretOverridesLister implements the FederatedSecretOverridesLister interface.
type federatedSecretOverridesLister struct {
	indexer cache.Indexer
}

// NewFederatedSecretOverridesLister returns a new FederatedSecretOverridesLister.
func NewFederatedSecretOverridesLister(indexer cache.Indexer) FederatedSecretOverridesLister {
	return &federatedSecretOverridesLister{indexer: indexer}
}

// List lists all FederatedSecretOverrideses in the indexer.
func (s *federatedSecretOverridesLister) List(selector labels.Selector) (ret []*v1alpha1.FederatedSecretOverrides, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FederatedSecretOverrides))
	})
	return ret, err
}

// FederatedSecretOverrideses returns an object that can list and get FederatedSecretOverrideses.
func (s *federatedSecretOverridesLister) FederatedSecretOverrideses(namespace string) FederatedSecretOverridesNamespaceLister {
	return federatedSecretOverridesNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FederatedSecretOverridesNamespaceLister helps list and get FederatedSecretOverrideses.
type FederatedSecretOverridesNamespaceLister interface {
	// List lists all FederatedSecretOverrideses in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.FederatedSecretOverrides, err error)
	// Get retrieves the FederatedSecretOverrides from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.FederatedSecretOverrides, error)
	FederatedSecretOverridesNamespaceListerExpansion
}

// federatedSecretOverridesNamespaceLister implements the FederatedSecretOverridesNamespaceLister
// interface.
type federatedSecretOverridesNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FederatedSecretOverrideses in the indexer for a given namespace.
func (s federatedSecretOverridesNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.FederatedSecretOverrides, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FederatedSecretOverrides))
	})
	return ret, err
}

// Get retrieves the FederatedSecretOverrides from the indexer for a given namespace and name.
func (s federatedSecretOverridesNamespaceLister) Get(name string) (*v1alpha1.FederatedSecretOverrides, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("federatedsecretoverrides"), name)
	}
	return obj.(*v1alpha1.FederatedSecretOverrides), nil
}
