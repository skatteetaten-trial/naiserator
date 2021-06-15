package gcp

import (
	"fmt"
	nais_io_v1 "github.com/nais/liberator/pkg/apis/nais.io/v1"
	google_iam "github.com/nais/naiserator/pkg/resourcecreator/google/iam"
	google_sql "github.com/nais/naiserator/pkg/resourcecreator/google/sql"
	google_storagebucket "github.com/nais/naiserator/pkg/resourcecreator/google/storagebucket"
	"github.com/nais/naiserator/pkg/resourcecreator/resource"
)

func Create(source resource.Source, ast *resource.Ast, resourceOptions resource.Options, naisGCP *nais_io_v1.GCP) error {
	if len(resourceOptions.GoogleProjectId) <= 0 {
		return nil
	}

	googleServiceAccount := google_iam.CreateServiceAccount(source, resourceOptions.GoogleProjectId)
	googleServiceAccountBinding := google_iam.CreatePolicy(source, &googleServiceAccount, resourceOptions.GoogleProjectId)
	ast.AppendOperation(resource.OperationCreateOrUpdate, &googleServiceAccount)
	ast.AppendOperation(resource.OperationCreateOrUpdate, &googleServiceAccountBinding)

	if naisGCP != nil {
		if len(resourceOptions.GoogleTeamProjectId) == 0 {
			return fmt.Errorf("cannot create GCP resource(s) in non team namespace")
		}

		google_storagebucket.Create(source, ast, resourceOptions, googleServiceAccount, naisGCP.Buckets)
		err := google_sql.CreateInstance(source, ast, resourceOptions, &naisGCP.SqlInstances)
		if err != nil {
			return err
		}
		err = google_iam.CreatePolicyMember(source, ast, resourceOptions, naisGCP.Permissions)
		if err != nil {
			return err
		}
	}

	return nil
}
