/*
Copyright The Kubernetes Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/kubescape/storage/pkg/apis/softwarecomposition/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// SBOMSPDXv2p3FilteredLister helps list SBOMSPDXv2p3Filtereds.
// All objects returned here must be treated as read-only.
type SBOMSPDXv2p3FilteredLister interface {
	// List lists all SBOMSPDXv2p3Filtereds in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.SBOMSPDXv2p3Filtered, err error)
	// SBOMSPDXv2p3Filtereds returns an object that can list and get SBOMSPDXv2p3Filtereds.
	SBOMSPDXv2p3Filtereds(namespace string) SBOMSPDXv2p3FilteredNamespaceLister
	SBOMSPDXv2p3FilteredListerExpansion
}

// sBOMSPDXv2p3FilteredLister implements the SBOMSPDXv2p3FilteredLister interface.
type sBOMSPDXv2p3FilteredLister struct {
	indexer cache.Indexer
}

// NewSBOMSPDXv2p3FilteredLister returns a new SBOMSPDXv2p3FilteredLister.
func NewSBOMSPDXv2p3FilteredLister(indexer cache.Indexer) SBOMSPDXv2p3FilteredLister {
	return &sBOMSPDXv2p3FilteredLister{indexer: indexer}
}

// List lists all SBOMSPDXv2p3Filtereds in the indexer.
func (s *sBOMSPDXv2p3FilteredLister) List(selector labels.Selector) (ret []*v1beta1.SBOMSPDXv2p3Filtered, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.SBOMSPDXv2p3Filtered))
	})
	return ret, err
}

// SBOMSPDXv2p3Filtereds returns an object that can list and get SBOMSPDXv2p3Filtereds.
func (s *sBOMSPDXv2p3FilteredLister) SBOMSPDXv2p3Filtereds(namespace string) SBOMSPDXv2p3FilteredNamespaceLister {
	return sBOMSPDXv2p3FilteredNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SBOMSPDXv2p3FilteredNamespaceLister helps list and get SBOMSPDXv2p3Filtereds.
// All objects returned here must be treated as read-only.
type SBOMSPDXv2p3FilteredNamespaceLister interface {
	// List lists all SBOMSPDXv2p3Filtereds in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.SBOMSPDXv2p3Filtered, err error)
	// Get retrieves the SBOMSPDXv2p3Filtered from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta1.SBOMSPDXv2p3Filtered, error)
	SBOMSPDXv2p3FilteredNamespaceListerExpansion
}

// sBOMSPDXv2p3FilteredNamespaceLister implements the SBOMSPDXv2p3FilteredNamespaceLister
// interface.
type sBOMSPDXv2p3FilteredNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SBOMSPDXv2p3Filtereds in the indexer for a given namespace.
func (s sBOMSPDXv2p3FilteredNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.SBOMSPDXv2p3Filtered, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.SBOMSPDXv2p3Filtered))
	})
	return ret, err
}

// Get retrieves the SBOMSPDXv2p3Filtered from the indexer for a given namespace and name.
func (s sBOMSPDXv2p3FilteredNamespaceLister) Get(name string) (*v1beta1.SBOMSPDXv2p3Filtered, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("sbomspdxv2p3filtered"), name)
	}
	return obj.(*v1beta1.SBOMSPDXv2p3Filtered), nil
}