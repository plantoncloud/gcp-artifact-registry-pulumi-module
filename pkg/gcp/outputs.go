package gcp

import (
	"context"

	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/iac/v1/stackjob/enums/stackjoboperationtype"

	"github.com/pkg/errors"
	"github.com/plantoncloud/artifact-store-pulumi-blueprint/pkg/gcp/repo/docker"
	"github.com/plantoncloud/artifact-store-pulumi-blueprint/pkg/gcp/repo/maven"
	"github.com/plantoncloud/artifact-store-pulumi-blueprint/pkg/gcp/repo/npm"
	"github.com/plantoncloud/artifact-store-pulumi-blueprint/pkg/gcp/repo/python"
	"github.com/plantoncloud/artifact-store-pulumi-blueprint/pkg/gcp/serviceaccount"
	code2cloudv1developafsmodel "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/artifactstore/model"
	code2cloudv1developafsstackgcpmodel "github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/artifactstore/stack/gcp/model"
	"github.com/plantoncloud/stack-job-runner-golang-sdk/pkg/stack/output/backend"
)

func Outputs(ctx context.Context, input *code2cloudv1developafsstackgcpmodel.ArtifactStoreGcpStackInput) (*code2cloudv1developafsstackgcpmodel.ArtifactStoreGcpStackOutputs, error) {
	stackOutput, err := backend.StackOutput(input.StackJob)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get stack output")
	}
	return OutputMapTransformer(stackOutput, input), nil
}

func OutputMapTransformer(stackOutput map[string]interface{}, input *code2cloudv1developafsstackgcpmodel.ArtifactStoreGcpStackInput) *code2cloudv1developafsstackgcpmodel.ArtifactStoreGcpStackOutputs {
	if input.StackJob.Spec.OperationType != stackjoboperationtype.StackJobOperationType_apply || stackOutput == nil {
		return &code2cloudv1developafsstackgcpmodel.ArtifactStoreGcpStackOutputs{}
	}
	artifactStoreId := input.ResourceInput.ArtifactStore.Metadata.Id
	dockerRepoName := docker.GetRepoName(artifactStoreId)
	mavenRepoName := maven.GetRepoName(artifactStoreId)
	npmRepoName := npm.GetRepoName(artifactStoreId)
	pythonRepoName := python.GetRepoName(artifactStoreId)
	return &code2cloudv1developafsstackgcpmodel.ArtifactStoreGcpStackOutputs{
		GcpArtifactRegistryStatus: &code2cloudv1developafsmodel.ArtifactStoreGcpArtifactRegistryStatus{
			ReaderServiceAccountEmail:     backend.GetVal(stackOutput, serviceaccount.GetReaderServiceAccountEmailOutputName(artifactStoreId)),
			ReaderServiceAccountKeyBase64: backend.GetVal(stackOutput, serviceaccount.GetReaderServiceAccountKeyOutputName(artifactStoreId)),
			WriterServiceAccountEmail:     backend.GetVal(stackOutput, serviceaccount.GetWriterServiceAccountEmailOutputName(artifactStoreId)),
			WriterServiceAccountKeyBase64: backend.GetVal(stackOutput, serviceaccount.GetWriterServiceAccountKeyOutputName(artifactStoreId)),
			DockerRepoName:                backend.GetVal(stackOutput, docker.GetDockerRepoNameOutputName(dockerRepoName)),
			DockerRepoHostname:            backend.GetVal(stackOutput, docker.GetDockerRepoHostnameOutputName(dockerRepoName)),
			DockerRepoUrl:                 backend.GetVal(stackOutput, docker.GetDockerRepoUrlOutputName(dockerRepoName)),
			MavenRepoName:                 backend.GetVal(stackOutput, maven.GetMavenRepoNameOutputName(mavenRepoName)),
			MavenRepoUrl:                  backend.GetVal(stackOutput, maven.GetMavenRepoUrlOutputName(mavenRepoName)),
			NpmRepoName:                   backend.GetVal(stackOutput, npm.GetNpmRepoNameOutputName(npmRepoName)),
			PythonRepoName:                backend.GetVal(stackOutput, python.GetPythonRepoNameOutputName(pythonRepoName)),
		},
	}
}
