package vampy_go

import (
	"context"
	"testing"

	"github.com/qvineox/vampy-go/vampy"
)

func TestRepositoryListGet(t *testing.T) {
	c := testClient(t)
	ctx := testCtx(t)
	got, err := c.RepositoryListGet(ctx, RepositoryListGetOptions{Paging: vampy.Paging{Limit: 1}})
	if err != nil {
		t.Fatalf("RepositoryListGet: %v", err)
	}
	requireSlice(t, got.Items)
}

func TestRepositoryCreate(t *testing.T) {
	c := testClient(t)
	ctx := testCtx(t)
	slug := uniqueName(t, "test-slug")
	got, err := c.RepositoryCreate(ctx, RepositoryCreateOptions{Body: vampy.RepositoryCreate{Repository: vampy.Repository{Name: uniqueName(t, "Test Repo"), Slug: slug, URL: "https://example.com/" + slug}}})
	if err != nil {
		t.Fatalf("RepositoryCreate: %v", err)
	}
	requireSchema(t, got)
	t.Cleanup(func() {
		_ = c.RepositoryDelete(context.Background(), RepositoryDeleteOptions{RepositoryID: uuidFromString(t, string(got.ID))})
	})
}

func TestRepositoryImportFromVCS(t *testing.T) {
	// requires an existing resource or external service
	t.Skip("requires an existing resource or external service")
}

func TestRepositoryDetailsGet(t *testing.T) {
	c := testClient(t)
	ctx := testCtx(t)
	slug := uniqueName(t, "test-slug")
	created, err := c.RepositoryCreate(ctx, RepositoryCreateOptions{Body: vampy.RepositoryCreate{Repository: vampy.Repository{Name: uniqueName(t, "Test Repo"), Slug: slug, URL: "https://example.com/" + slug}}})
	if err != nil {
		t.Fatalf("RepositoryCreate: %v", err)
	}
	t.Cleanup(func() {
		_ = c.RepositoryDelete(context.Background(), RepositoryDeleteOptions{RepositoryID: uuidFromString(t, string(created.ID))})
	})
	got, err := c.RepositoryDetailsGet(ctx, RepositoryDetailsGetOptions{RepositoryID: uuidFromString(t, string(created.ID))})
	if err != nil {
		t.Fatalf("RepositoryDetailsGet: %v", err)
	}
	requireSchema(t, got)
}

func TestRepositoryDelete(t *testing.T) {
	c := testClient(t)
	ctx := testCtx(t)
	slug := uniqueName(t, "test-slug")
	created, err := c.RepositoryCreate(ctx, RepositoryCreateOptions{Body: vampy.RepositoryCreate{Repository: vampy.Repository{Name: uniqueName(t, "Test Repo"), Slug: slug, URL: "https://example.com/" + slug}}})
	if err != nil {
		t.Fatalf("RepositoryCreate: %v", err)
	}
	if err := c.RepositoryDelete(ctx, RepositoryDeleteOptions{RepositoryID: uuidFromString(t, string(created.ID))}); err != nil {
		t.Fatalf("RepositoryDelete: %v", err)
	}
}

func TestRepositoryPatch(t *testing.T) {
	c := testClient(t)
	ctx := testCtx(t)
	slug := uniqueName(t, "test-slug")
	name2 := uniqueName(t, "test-res-renamed")
	created, err := c.RepositoryCreate(ctx, RepositoryCreateOptions{Body: vampy.RepositoryCreate{Repository: vampy.Repository{Name: uniqueName(t, "Test Repo"), Slug: slug, URL: "https://example.com/" + slug}}})
	if err != nil {
		t.Fatalf("RepositoryCreate: %v", err)
	}
	t.Cleanup(func() {
		_ = c.RepositoryDelete(context.Background(), RepositoryDeleteOptions{RepositoryID: uuidFromString(t, string(created.ID))})
	})
	got, err := c.RepositoryPatch(ctx, RepositoryPatchOptions{RepositoryID: uuidFromString(t, string(created.ID)), Body: vampy.RepositoryPatch{Name: name2}})
	if err != nil {
		t.Fatalf("RepositoryPatch: %v", err)
	}
	requireSchema(t, got)
}

func TestRepositoryUserGroupsGet(t *testing.T) {
	// requires an existing resource
	t.Skip("requires an existing resource")
}

func TestRepositoryUsersSet(t *testing.T) {
	// requires an existing resource or external service
	t.Skip("requires an existing resource or external service")
}

func TestRepositoryUsersUnset(t *testing.T) {
	// requires an existing resource or external service
	t.Skip("requires an existing resource or external service")
}

func TestRepositoryUserGroupsSet(t *testing.T) {
	// requires an existing resource or external service
	t.Skip("requires an existing resource or external service")
}

func TestRepositoryUserGroupsUnset(t *testing.T) {
	// requires an existing resource or external service
	t.Skip("requires an existing resource or external service")
}

func TestRepositoryVCSLink(t *testing.T) {
	// requires an existing resource or external service
	t.Skip("requires an existing resource or external service")
}

func TestRepositoryVCSUnlink(t *testing.T) {
	// requires an existing resource or external service
	t.Skip("requires an existing resource or external service")
}

func TestRepositoryBranchListGet(t *testing.T) {
	c := testSpaceClient(t)
	ctx := testCtx(t)
	got, err := c.RepositoryBranchListGet(ctx, RepositoryBranchListGetOptions{Paging: vampy.Paging{Limit: 1}})
	if err != nil {
		t.Fatalf("RepositoryBranchListGet: %v", err)
	}
	requireSlice(t, got.Items)
}

func TestRepositoryBranchDelete(t *testing.T) {
	// requires an existing resource or external service
	t.Skip("requires an existing resource or external service")
}
