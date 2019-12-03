// generated by your friendly code generator. DO NOT EDIT.
// to refresh this file, run `go generate` in your shell.

package updater

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	autoscalingv1 "k8s.io/api/autoscaling/v1"
	clientV1Alpha1 "github.com/nais/naiserator/pkg/client/clientset/versioned"
	extensionsv1beta1 "k8s.io/api/extensions/v1beta1"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	typed_core_v1 "k8s.io/client-go/kubernetes/typed/core/v1"
	typed_apps_v1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	typed_autoscaling_v1 "k8s.io/client-go/kubernetes/typed/autoscaling/v1"
	typed_extensions_v1beta1 "k8s.io/client-go/kubernetes/typed/extensions/v1beta1"
	typed_networking_v1 "k8s.io/client-go/kubernetes/typed/networking/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	rbacv1 "k8s.io/api/rbac/v1"
	"github.com/nais/naiserator/pkg/apis/rbac.istio.io/v1alpha1"
	typed_rbac_v1 "k8s.io/client-go/kubernetes/typed/rbac/v1"
	istio_v1alpha1 "github.com/nais/naiserator/pkg/client/clientset/versioned/typed/rbac.istio.io/v1alpha1"
	typed_networking_istio_io_v1alpha3 "github.com/nais/naiserator/pkg/client/clientset/versioned/typed/networking.istio.io/v1alpha3"
	networking_istio_io_v1alpha3 "github.com/nais/naiserator/pkg/apis/networking.istio.io/v1alpha3"
	typed_iam_cnrm_cloud_google_com_v1alpha1 "github.com/nais/naiserator/pkg/client/clientset/versioned/typed/iam.cnrm.cloud.google.com/v1alpha1"
	iam_cnrm_cloud_google_com_v1alpha1 "github.com/nais/naiserator/pkg/apis/iam.cnrm.cloud.google.com/v1alpha1"

)

{{range .}}

func {{.Name}}(client {{.Interface}}, old, new {{.Type}}) func() error {
	log.Infof("creating or updating {{ .Type }} for %s", new.Name)
	if old == nil {
		return func() error {
			_, err := client.Create(new)
			return err
		}
	}

	CopyMeta(old, new)
	{{if .TransformFunc}}
		{{.TransformFunc}}(old, new)
	{{end}}

	return func() error {
		_, err := client.Update(new)
		return err
	}
}

{{end}}

func CreateOrUpdate(clientSet kubernetes.Interface, customClient clientV1Alpha1.Interface, resource runtime.Object) func() error {
	switch new := resource.(type) {
	{{range .}}
		case {{.Type}}:
		c := {{.Client}}(new.Namespace)
		old, err := c.Get(new.Name, metav1.GetOptions{})
		if err != nil {
			if !errors.IsNotFound(err) {
				return func() error { return err }
			}
			return {{.Name}}(c, nil, new)
		}
		return {{.Name}}(c, old, new)
	{{end}}
	default:
		panic(fmt.Errorf("BUG! You didn't specify a case for type '%T' in the file hack/generator/updater.go", new))
	}
}

func CreateOrRecreate(clientSet kubernetes.Interface, customClient clientV1Alpha1.Interface, resource runtime.Object) func() error {
	switch new := resource.(type) {
	{{range .}}
		case {{.Type}}:
		c := {{.Client}}(new.Namespace)
		return func() error {
            log.Infof("pre-deleting {{ .Type }} for %s", new.Name)
			err := c.Delete(new.Name, &metav1.DeleteOptions{})
			if err != nil && !errors.IsNotFound(err) {
				return err
			}
            log.Infof("creating new {{ .Type }} for %s", new.Name)
            _, err = c.Create(new)
            return err
		}
	{{end}}
	default:
		panic(fmt.Errorf("BUG! You didn't specify a case for type '%T' in the file hack/generator/updater.go", new))
	}
}

func DeleteIfExists(clientSet kubernetes.Interface, customClient clientV1Alpha1.Interface, resource runtime.Object) func() error {
	switch new := resource.(type) {
	{{range .}}
		case {{.Type}}:
		c := {{.Client}}(new.Namespace)
		return func() error {
            log.Infof("deleting {{ .Type }} for %s", new.Name)
			err := c.Delete(new.Name, &metav1.DeleteOptions{})
			if err != nil && errors.IsNotFound(err) {
				return nil
			}
			return err
		}
	{{end}}
	default:
		panic(fmt.Errorf("BUG! You didn't specify a case for type '%T' in the file hack/generator/updater.go", new))
	}
}
