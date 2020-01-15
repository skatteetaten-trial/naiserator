// Code generated by lister-gen. DO NOT EDIT.

package v1alpha3

import (
	v1alpha3 "github.com/nais/naiserator/pkg/apis/sql.cnrm.cloud.google.com/v1alpha3"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// SQLUserLister helps list SQLUsers.
type SQLUserLister interface {
	// List lists all SQLUsers in the indexer.
	List(selector labels.Selector) (ret []*v1alpha3.SQLUser, err error)
	// SQLUsers returns an object that can list and get SQLUsers.
	SQLUsers(namespace string) SQLUserNamespaceLister
	SQLUserListerExpansion
}

// sQLUserLister implements the SQLUserLister interface.
type sQLUserLister struct {
	indexer cache.Indexer
}

// NewSQLUserLister returns a new SQLUserLister.
func NewSQLUserLister(indexer cache.Indexer) SQLUserLister {
	return &sQLUserLister{indexer: indexer}
}

// List lists all SQLUsers in the indexer.
func (s *sQLUserLister) List(selector labels.Selector) (ret []*v1alpha3.SQLUser, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha3.SQLUser))
	})
	return ret, err
}

// SQLUsers returns an object that can list and get SQLUsers.
func (s *sQLUserLister) SQLUsers(namespace string) SQLUserNamespaceLister {
	return sQLUserNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SQLUserNamespaceLister helps list and get SQLUsers.
type SQLUserNamespaceLister interface {
	// List lists all SQLUsers in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha3.SQLUser, err error)
	// Get retrieves the SQLUser from the indexer for a given namespace and name.
	Get(name string) (*v1alpha3.SQLUser, error)
	SQLUserNamespaceListerExpansion
}

// sQLUserNamespaceLister implements the SQLUserNamespaceLister
// interface.
type sQLUserNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SQLUsers in the indexer for a given namespace.
func (s sQLUserNamespaceLister) List(selector labels.Selector) (ret []*v1alpha3.SQLUser, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha3.SQLUser))
	})
	return ret, err
}

// Get retrieves the SQLUser from the indexer for a given namespace and name.
func (s sQLUserNamespaceLister) Get(name string) (*v1alpha3.SQLUser, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha3.Resource("sqluser"), name)
	}
	return obj.(*v1alpha3.SQLUser), nil
}
