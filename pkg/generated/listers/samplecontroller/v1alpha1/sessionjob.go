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

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "k8s.io/k8s-sym-client-controller/pkg/apis/samplecontroller/v1alpha1"
)

// SessionJobLister helps list SessionJobs.
type SessionJobLister interface {
	// List lists all SessionJobs in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.SessionJob, err error)
	// SessionJobs returns an object that can list and get SessionJobs.
	SessionJobs(namespace string) SessionJobNamespaceLister
	SessionJobListerExpansion
}

// sessionJobLister implements the SessionJobLister interface.
type sessionJobLister struct {
	indexer cache.Indexer
}

// NewSessionJobLister returns a new SessionJobLister.
func NewSessionJobLister(indexer cache.Indexer) SessionJobLister {
	return &sessionJobLister{indexer: indexer}
}

// List lists all SessionJobs in the indexer.
func (s *sessionJobLister) List(selector labels.Selector) (ret []*v1alpha1.SessionJob, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SessionJob))
	})
	return ret, err
}

// SessionJobs returns an object that can list and get SessionJobs.
func (s *sessionJobLister) SessionJobs(namespace string) SessionJobNamespaceLister {
	return sessionJobNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SessionJobNamespaceLister helps list and get SessionJobs.
type SessionJobNamespaceLister interface {
	// List lists all SessionJobs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.SessionJob, err error)
	// Get retrieves the SessionJob from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.SessionJob, error)
	SessionJobNamespaceListerExpansion
}

// sessionJobNamespaceLister implements the SessionJobNamespaceLister
// interface.
type sessionJobNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SessionJobs in the indexer for a given namespace.
func (s sessionJobNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SessionJob, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SessionJob))
	})
	return ret, err
}

// Get retrieves the SessionJob from the indexer for a given namespace and name.
func (s sessionJobNamespaceLister) Get(name string) (*v1alpha1.SessionJob, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("sessionjob"), name)
	}
	return obj.(*v1alpha1.SessionJob), nil
}
