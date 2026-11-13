package vampy_go

import (
	"bytes"
	"context"
	"encoding/json"
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

	const (
		delay time.Duration = 5 * time.Second
	)

	t.Run("async export repository issues", func(t *testing.T) {
		if testRepositoryUUID == "" || testRepositoryBranchUUID == "" {
			t.Skip("missing VAMPY_TEST_REPOSITORY_UUID or VAMPY_TEST_REPOSITORY_BRANCH_UUID environment variable, skipping test...")
		}

		ctx := context.Background()

		report, err := testClient.ScanIssuesAsyncExport(ctx, ScanIssuesAsyncExportOptions{
			Formatter:       new(vampy.AsyncExportFormatterJSON),
			IncludeComments: new(false),
			Order:           new("parser,severity"),
			TemplateSlug:    new(vampy.TemplateSlugDETAILED),
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
			retries int = 5
			state   vampy.Report
		)

		for retries > 0 {
			time.Sleep(delay)
			state, err = testClient.ReportGet(ctx, ReportGetOptions{
				ReportID: report.ID,
			})

			require.NoError(t, err)
			require.NotNil(t, state)

			status := vampy.ReportStatus(state.Status.(string))
			if status == vampy.ReportStatusPROGRESS || status == vampy.ReportStatusPENDING {
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

	testScanPipelineUUID := os.Getenv("VAMPY_TEST_SCAN_PIPELINE_UUID")

	t.Run("async export scan pipeline issues", func(t *testing.T) {
		if testScanPipelineUUID == "" {
			t.Skip("missing VAMPY_TEST_SCAN_PIPELINE_UUID environment variable, skipping test...")
		}

		ctx := context.Background()

		report, err := testClient.ScanIssuesAsyncExport(ctx, ScanIssuesAsyncExportOptions{
			Formatter:       new(vampy.AsyncExportFormatterJSON),
			IncludeComments: new(true),
			Order:           new("parser,severity"),
			TemplateSlug:    new(vampy.TemplateSlugDETAILED),
			Body: vampy.ScanIssueFilter{
				Repositories:   []uuid.UUID{uuid.MustParse(testRepositoryUUID)},
				ScanPipelineID: new(uuid.MustParse(testScanPipelineUUID)),
				HasSolution:    false,
				OldIssues:      new(false),
				FindingState:   vampy.FindingStateDETECTED,
			},
		})

		require.NotNil(t, report)
		require.Nil(t, err)

		var (
			retries int = 5
			state   vampy.Report
		)

		for retries > 0 {
			time.Sleep(delay)
			state, err = testClient.ReportGet(ctx, ReportGetOptions{
				ReportID: report.ID,
			})

			require.NoError(t, err)
			require.NotNil(t, state)

			status := vampy.ReportStatus(state.Status.(string))
			if status == vampy.ReportStatusPROGRESS || status == vampy.ReportStatusPENDING {
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
		require.NotEmpty(t, download)

		var detectedIssues []vampy.ScanIssueExport
		err = json.NewDecoder(bytes.NewReader(download)).Decode(&detectedIssues)
		require.NoError(t, err)

		require.Len(t, detectedIssues, 15)
	})
}
