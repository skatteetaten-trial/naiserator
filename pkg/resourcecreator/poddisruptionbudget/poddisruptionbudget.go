package poddisruptionbudget

import (
	nais_io_v1 "github.com/nais/liberator/pkg/apis/nais.io/v1"
	policyv1beta1 "k8s.io/api/policy/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"

	"github.com/nais/naiserator/pkg/resourcecreator/resource"
)

type Source interface {
	resource.Source
	GetReplicas() *nais_io_v1.Replicas
}

func Create(source Source, ast *resource.Ast) {
	replicas := source.GetReplicas()

	if *replicas.Max == 1 || *replicas.Min == 1 {
		return
	}

	maxUnavailable := intstr.FromInt(1)

	podDisruptionBudget := &policyv1beta1.PodDisruptionBudget{
		TypeMeta: metav1.TypeMeta{
			Kind:       "PodDisruptionBudget",
			APIVersion: "policy/v1beta1",
		},
		ObjectMeta: resource.CreateObjectMeta(source),
		Spec: policyv1beta1.PodDisruptionBudgetSpec{
			MaxUnavailable: &maxUnavailable,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": source.GetName(),
				},
			},
		},
	}

	ast.AppendOperation(resource.OperationCreateOrUpdate, podDisruptionBudget)
}
