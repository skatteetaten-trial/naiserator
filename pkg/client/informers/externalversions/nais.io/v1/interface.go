// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	internalinterfaces "github.com/nais/naiserator/pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// AzureAdApplications returns a AzureAdApplicationInformer.
	AzureAdApplications() AzureAdApplicationInformer
	// Jwkers returns a JwkerInformer.
	Jwkers() JwkerInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// AzureAdApplications returns a AzureAdApplicationInformer.
func (v *version) AzureAdApplications() AzureAdApplicationInformer {
	return &azureAdApplicationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Jwkers returns a JwkerInformer.
func (v *version) Jwkers() JwkerInformer {
	return &jwkerInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
