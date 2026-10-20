package vampy_go

import (
	"context"
	"os"
	"testing"
	"uuid"

	"github.com/qvineox/vampy-go/vampy"
	"github.com/stretchr/testify/require"
)

func TestClient_GetRepositoryDashboard(t *testing.T) {
	testRepositoryUUID := os.Getenv("VAMPY_TEST_REPOSITORY_UUID")
	testRepositoryBranchUUID := os.Getenv("VAMPY_TEST_REPOSITORY_BRANCH_UUID")

	t.Run("get repository dashboard", func(t *testing.T) {
		if testRepositoryUUID == "" {
			t.Skip("missing VAMPY_TEST_REPOSITORY_UUID environment variable, skipping test...")
		}

		opts := GetRepositoryDashboardOptions{
			RepositoryID: uuid.MustParse(testRepositoryUUID),
		}

		dashboard, err := testClient.GetRepositoryDashboard(context.Background(), opts)
		require.NoError(t, err)
		require.NotNil(t, dashboard)
	})

	t.Run("get non-existent repository dashboard", func(t *testing.T) {
		opts := GetRepositoryDashboardOptions{
			RepositoryID: uuid.NewV7(),
		}

		dashboard, err := testClient.GetRepositoryDashboard(context.Background(), opts)
		require.Error(t, err)
		require.NotNil(t, dashboard)
		require.Empty(t, dashboard.Cards)
	})

	t.Run("get repository dashboard with branch", func(t *testing.T) {
		if testRepositoryUUID == "" || testRepositoryBranchUUID == "" {
			t.Skip("missing VAMPY_TEST_REPOSITORY_UUID or VAMPY_TEST_REPOSITORY_BRANCH_UUID environment variable, skipping test...")
		}

		opts := GetRepositoryDashboardOptions{
			RepositoryID:       uuid.MustParse(testRepositoryUUID),
			RepositoryBranchID: new(uuid.MustParse(testRepositoryBranchUUID)),
		}

		dashboard, err := testClient.GetRepositoryDashboard(context.Background(), opts)
		require.NoError(t, err)
		require.NotNil(t, dashboard)
	})

	t.Run("get repository dashboard with all branches", func(t *testing.T) {
		if testRepositoryUUID == "" {
			t.Skip("missing VAMPY_TEST_REPOSITORY_UUID environment variable, skipping test...")
		}

		opts := GetRepositoryDashboardOptions{
			RepositoryID: uuid.MustParse(testRepositoryUUID),
			AllBranches:  new(true),
		}

		dashboard, err := testClient.GetRepositoryDashboard(context.Background(), opts)
		require.NoError(t, err)
		require.NotNil(t, dashboard)

		_, _, _, _, _, totaOpenIssues := dashboard.GetOpenedIssues()
		require.NotZero(t, totaOpenIssues)
	})

	t.Run("get repository sast dashboard with all branches", func(t *testing.T) {
		if testRepositoryUUID == "" {
			t.Skip("missing VAMPY_TEST_REPOSITORY_UUID environment variable, skipping test...")
		}

		opts := GetRepositoryDashboardOptions{
			RepositoryID: uuid.MustParse(testRepositoryUUID),
			AllBranches:  new(true),
			ScannerType:  new(vampy.ScannerTypeSAST),
		}

		dashboard, err := testClient.GetRepositoryDashboard(context.Background(), opts)
		require.NoError(t, err)
		require.NotNil(t, dashboard)
	})

	t.Run("get repository dast dashboard with all branches", func(t *testing.T) {
		if testRepositoryUUID == "" {
			t.Skip("missing VAMPY_TEST_REPOSITORY_UUID environment variable, skipping test...")
		}

		opts := GetRepositoryDashboardOptions{
			RepositoryID: uuid.MustParse(testRepositoryUUID),
			AllBranches:  new(true),
			ScannerType:  new(vampy.ScannerTypeDAST),
		}

		dashboard, err := testClient.GetRepositoryDashboard(context.Background(), opts)
		require.NoError(t, err)
		require.NotNil(t, dashboard)
	})
}
