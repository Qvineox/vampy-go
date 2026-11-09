package vampy_go

import (
	"context"
	"os"
	"testing"
	"time"
	"uuid"

	"github.com/qvineox/vampy-go/vampy"
	"github.com/stretchr/testify/require"
)

func TestClient_AsyncExportReport(t *testing.T) {
	testRepositoryUUID := os.Getenv("VAMPY_TEST_REPOSITORY_UUID")
	testRepositoryBranchUUID := os.Getenv("VAMPY_TEST_REPOSITORY_BRANCH_UUID")

	t.Run("async export repository issues", func(t *testing.T) {
		ctx := context.Background()

		report, err := testClient.ScanIssuesAsyncExport(ctx, ScanIssuesAsyncExportOptions{
			Formatter:       new(vampy.AsyncExportFormatterCSV),
			IncludeComments: new(false),
			Order:           new("parser,severity"),
			TemplateSlug:    new(vampy.TemplateSlugDETAILED),
			CSVDialect:      new(vampy.CSVExportDialectEXCEL),
			CSVDelimiter:    new(vampy.CSVExportDelimiterSEMICOLON),
			Body: vampy.ScanIssueFilter{
				Repositories:       []uuid.UUID{uuid.MustParse(testRepositoryUUID)},
				RepositoryBranches: []uuid.UUID{uuid.MustParse(testRepositoryBranchUUID)},
				HasSolution:        false,
				OldIssues:          new(false),
			},
		})

		require.NotNil(t, report)
		require.Nil(t, err)

		var (
			retries int           = 5
			delay   time.Duration = 2 * time.Second

			state vampy.Report
		)

		for retries > 0 {
			time.Sleep(delay)
			state, err = testClient.ReportGet(ctx, ReportGetOptions{
				ReportID: report.ID,
			})

			require.NoError(t, err)
			require.NotNil(t, state)

			if state.Status == vampy.ReportStatusPROGRESS {
				retries--
				continue
			} else {
				break
			}
		}

		require.EqualValues(t, vampy.ReportStatusSUCCESS, state.Status)

		download, err := testClient.ReportDownload(ctx, ReportDownloadOptions{
			ReportID: report.ID,
		})

		require.NoError(t, err)
		require.NotNil(t, download)
	})
}
