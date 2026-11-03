package vampy_go

import (
	"context"
	"fmt"
	"testing"
	"uuid"

	"github.com/qvineox/vampy-go/vampy"
	"github.com/stretchr/testify/require"
)

func TestClient_CreateRepository(t *testing.T) {
	var createdRepositoryUUID uuid.UUID

	t.Cleanup(func() {
		if createdRepositoryUUID == uuid.Nil() {
			t.Skip("repository has not been created")
		}

		err := testClient.RepositoryDelete(context.Background(), RepositoryDeleteOptions{
			RepositoryID: createdRepositoryUUID,
		})

		require.NoError(t, err)
	})

	t.Run("create new repository", func(t *testing.T) {
		name := fmt.Sprintf("test-repository-%s", uuid.New().String())
		opts := RepositoryCreateOptions{
			vampy.RepositoryCreate{
				Slug:        name,
				Name:        name,
				Type:        vampy.RepositoryTypeSOURCECODE,
				Criticality: vampy.CriticalityLOW,
				Description: "Auto-test created repository",
				URL:         "https://example.com",
			},
		}

		repository, err := testClient.RepositoryCreate(context.Background(), opts)
		require.NoError(t, err)
		require.NotNil(t, repository)
		require.NotNil(t, repository.ID)
		require.True(t, *repository.ID != uuid.Nil())

		createdRepositoryUUID = *repository.ID
	})
}
