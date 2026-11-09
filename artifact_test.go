package vampy_go

import (
	"context"
	"os"
	"testing"
	"uuid"

	"github.com/qvineox/vampy-go/vampy"
	"github.com/stretchr/testify/require"
)

func TestClient_GetArtifact(t *testing.T) {
	testArtifactUUID := os.Getenv("VAMPY_TEST_ARTIFACT_UUID")
	require.NotEmpty(t, testArtifactUUID)

	t.Run("get all artifact versions", func(t *testing.T) {
		ctx := context.Background()

		versions, err := testClient.ArtifactsVersionsGet(ctx, ArtifactsVersionsGetOptions{
			ArtifactID: uuid.MustParse(testArtifactUUID),
			Order:      new("-created"),
			Limit:      1000,
		})

		require.NoError(t, err)
		require.NotNil(t, versions)
		require.NotEmpty(t, versions.Items)
	})
}
func TestClient_DeleteArtifact(t *testing.T) {
	testArtifactUUID := os.Getenv("VAMPY_TEST_ARTIFACT_UUID")

	t.Run("delete artifact by uuid", func(t *testing.T) {
		ctx := context.Background()

		err := testClient.ArtifactsDeleteBatch(ctx, ArtifactsDeleteBatchOptions{
			Body: vampy.ArtifactDelete{
				Artifacts: []uuid.UUID{
					uuid.MustParse(testArtifactUUID),
				},
			},
		})

		require.NoError(t, err)

		//testClient.ArtifactVersionsDeleteBatch()
	})
}
