// Code generated by informer-gen. DO NOT EDIT.

package externalversions

import (
	"fmt"

	v1beta1 "github.com/nais/naiserator/pkg/apis/iam.cnrm.cloud.google.com/v1beta1"
	v1 "github.com/nais/naiserator/pkg/apis/nais.io/v1"
	v1alpha1 "github.com/nais/naiserator/pkg/apis/nais.io/v1alpha1"
	v1alpha3 "github.com/nais/naiserator/pkg/apis/networking.istio.io/v1alpha3"
	sqlcnrmcloudgooglecomv1beta1 "github.com/nais/naiserator/pkg/apis/sql.cnrm.cloud.google.com/v1beta1"
	storagecnrmcloudgooglecomv1beta1 "github.com/nais/naiserator/pkg/apis/storage.cnrm.cloud.google.com/v1beta1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=iam.cnrm.cloud.google.com, Version=v1beta1
	case v1beta1.SchemeGroupVersion.WithResource("iampolicies"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Iam().V1beta1().IAMPolicies().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("iampolicymembers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Iam().V1beta1().IAMPolicyMembers().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("iamserviceaccounts"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Iam().V1beta1().IAMServiceAccounts().Informer()}, nil

		// Group=nais.io, Version=v1
	case v1.SchemeGroupVersion.WithResource("azureadapplications"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Nais().V1().AzureAdApplications().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("jwkers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Nais().V1().Jwkers().Informer()}, nil

		// Group=nais.io, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithResource("applications"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Nais().V1alpha1().Applications().Informer()}, nil

		// Group=networking.istio.io, Version=v1alpha3
	case v1alpha3.SchemeGroupVersion.WithResource("serviceentries"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Networking().V1alpha3().ServiceEntries().Informer()}, nil
	case v1alpha3.SchemeGroupVersion.WithResource("virtualservices"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Networking().V1alpha3().VirtualServices().Informer()}, nil

		// Group=sql.cnrm.cloud.google.com, Version=v1beta1
	case sqlcnrmcloudgooglecomv1beta1.SchemeGroupVersion.WithResource("sqldatabases"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Sql().V1beta1().SQLDatabases().Informer()}, nil
	case sqlcnrmcloudgooglecomv1beta1.SchemeGroupVersion.WithResource("sqlinstances"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Sql().V1beta1().SQLInstances().Informer()}, nil
	case sqlcnrmcloudgooglecomv1beta1.SchemeGroupVersion.WithResource("sqlusers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Sql().V1beta1().SQLUsers().Informer()}, nil

		// Group=storage.cnrm.cloud.google.com, Version=v1beta1
	case storagecnrmcloudgooglecomv1beta1.SchemeGroupVersion.WithResource("storagebuckets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Storage().V1beta1().StorageBuckets().Informer()}, nil
	case storagecnrmcloudgooglecomv1beta1.SchemeGroupVersion.WithResource("storagebucketaccesscontrols"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Storage().V1beta1().StorageBucketAccessControls().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
