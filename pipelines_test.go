package vampy_go

import (
	"context"
	"os"
	"testing"
	"uuid"

	"github.com/qvineox/vampy-go/vampy"
	"github.com/stretchr/testify/require"
)

func TestClient_GetPipelines(t *testing.T) {
	testRepositoryUUID := os.Getenv("VAMPY_TEST_REPOSITORY_UUID")

	t.Run("get scan pipelines", func(t *testing.T) {
		if testRepositoryUUID == "" {
			t.Skip("missing VAMPY_TEST_REPOSITORY_UUID environment variable, skipping test...")
		}

		ctx := context.Background()

		pipelines, err := testClient.ScanPipelinesGet(ctx, ScanPipelinesGetOptions{
			RepositoryID: new(uuid.MustParse(testRepositoryUUID)),
			RelationType: []vampy.ScanPipelineRelationType{
				vampy.ScanPipelineRelationTypeREPOSITORY,
			},
			Order: new("-created"),
			Limit: 10,
		})

		require.NotNil(t, pipelines)
		require.Nil(t, err)

		require.NotEmpty(t, pipelines.Items)
	})
}
