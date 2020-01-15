// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/nais/naiserator/pkg/apis/iam.cnrm.cloud.google.com/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// IAMPolicyMemberLister helps list IAMPolicyMembers.
type IAMPolicyMemberLister interface {
	// List lists all IAMPolicyMembers in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.IAMPolicyMember, err error)
	// IAMPolicyMembers returns an object that can list and get IAMPolicyMembers.
	IAMPolicyMembers(namespace string) IAMPolicyMemberNamespaceLister
	IAMPolicyMemberListerExpansion
}

// iAMPolicyMemberLister implements the IAMPolicyMemberLister interface.
type iAMPolicyMemberLister struct {
	indexer cache.Indexer
}

// NewIAMPolicyMemberLister returns a new IAMPolicyMemberLister.
func NewIAMPolicyMemberLister(indexer cache.Indexer) IAMPolicyMemberLister {
	return &iAMPolicyMemberLister{indexer: indexer}
}

// List lists all IAMPolicyMembers in the indexer.
func (s *iAMPolicyMemberLister) List(selector labels.Selector) (ret []*v1alpha1.IAMPolicyMember, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IAMPolicyMember))
	})
	return ret, err
}

// IAMPolicyMembers returns an object that can list and get IAMPolicyMembers.
func (s *iAMPolicyMemberLister) IAMPolicyMembers(namespace string) IAMPolicyMemberNamespaceLister {
	return iAMPolicyMemberNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// IAMPolicyMemberNamespaceLister helps list and get IAMPolicyMembers.
type IAMPolicyMemberNamespaceLister interface {
	// List lists all IAMPolicyMembers in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.IAMPolicyMember, err error)
	// Get retrieves the IAMPolicyMember from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.IAMPolicyMember, error)
	IAMPolicyMemberNamespaceListerExpansion
}

// iAMPolicyMemberNamespaceLister implements the IAMPolicyMemberNamespaceLister
// interface.
type iAMPolicyMemberNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all IAMPolicyMembers in the indexer for a given namespace.
func (s iAMPolicyMemberNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.IAMPolicyMember, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IAMPolicyMember))
	})
	return ret, err
}

// Get retrieves the IAMPolicyMember from the indexer for a given namespace and name.
func (s iAMPolicyMemberNamespaceLister) Get(name string) (*v1alpha1.IAMPolicyMember, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("iampolicymember"), name)
	}
	return obj.(*v1alpha1.IAMPolicyMember), nil
}
