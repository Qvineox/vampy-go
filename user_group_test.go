package vampy_go

import (
	"context"
	"testing"

	"github.com/qvineox/vampy-go/vampy"
)

func TestGetUserGroups(t *testing.T) {
	c := testSpaceClient(t)
	ctx := testCtx(t)
	got, err := c.GetUserGroups(ctx, GetUserGroupsOptions{Paging: vampy.Paging{Limit: 1}})
	if err != nil {
		t.Fatalf("GetUserGroups: %v", err)
	}
	requireSlice(t, got.Items)
}

func TestPostUserGroup(t *testing.T) {
	c := testSpaceClient(t)
	ctx := testCtx(t)
	name := uniqueName(t, "test-res")
	got, err := c.PostUserGroup(ctx, PostUserGroupOptions{Body: vampy.UserGroupCreateEdit{Name: name}})
	if err != nil {
		t.Fatalf("PostUserGroup: %v", err)
	}
	requireSchema(t, got)
	t.Cleanup(func() {
		_ = c.DeleteUserGroup(context.Background(), DeleteUserGroupOptions{UserGroupID: uuidFromString(t, string(got.ID))})
	})
}

func TestGetUserGroup(t *testing.T) {
	c := testSpaceClient(t)
	ctx := testCtx(t)
	name := uniqueName(t, "test-res")
	created, err := c.PostUserGroup(ctx, PostUserGroupOptions{Body: vampy.UserGroupCreateEdit{Name: name}})
	if err != nil {
		t.Fatalf("PostUserGroup: %v", err)
	}
	t.Cleanup(func() {
		_ = c.DeleteUserGroup(context.Background(), DeleteUserGroupOptions{UserGroupID: uuidFromString(t, string(created.ID))})
	})
	got, err := c.GetUserGroup(ctx, GetUserGroupOptions{UserGroupID: uuidFromString(t, string(created.ID))})
	if err != nil {
		t.Fatalf("GetUserGroup: %v", err)
	}
	requireSchema(t, got)
}

func TestDeleteUserGroup(t *testing.T) {
	c := testSpaceClient(t)
	ctx := testCtx(t)
	name := uniqueName(t, "test-res")
	created, err := c.PostUserGroup(ctx, PostUserGroupOptions{Body: vampy.UserGroupCreateEdit{Name: name}})
	if err != nil {
		t.Fatalf("PostUserGroup: %v", err)
	}
	if err := c.DeleteUserGroup(ctx, DeleteUserGroupOptions{UserGroupID: uuidFromString(t, string(created.ID))}); err != nil {
		t.Fatalf("DeleteUserGroup: %v", err)
	}
}

func TestPatchUserGroup(t *testing.T) {
	c := testSpaceClient(t)
	ctx := testCtx(t)
	name := uniqueName(t, "test-res")
	name2 := uniqueName(t, "test-res-renamed")
	created, err := c.PostUserGroup(ctx, PostUserGroupOptions{Body: vampy.UserGroupCreateEdit{Name: name}})
	if err != nil {
		t.Fatalf("PostUserGroup: %v", err)
	}
	t.Cleanup(func() {
		_ = c.DeleteUserGroup(context.Background(), DeleteUserGroupOptions{UserGroupID: uuidFromString(t, string(created.ID))})
	})
	got, err := c.PatchUserGroup(ctx, PatchUserGroupOptions{UserGroupID: uuidFromString(t, string(created.ID)), Body: vampy.UserGroupCreateEdit{Name: name2}})
	if err != nil {
		t.Fatalf("PatchUserGroup: %v", err)
	}
	requireSchema(t, got)
}

func TestGetUserGroupRelations(t *testing.T) {
	// requires an existing resource
	t.Skip("requires an existing resource")
}

func TestUserGroupsSetRole(t *testing.T) {
	// mutates global/shared state; skipped for safety
	t.Skip("mutates global/shared state; skipped for safety")
}

func TestUserGroupMoveUsers(t *testing.T) {
	// mutates global/shared state; skipped for safety
	t.Skip("mutates global/shared state; skipped for safety")
}
