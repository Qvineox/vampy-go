package vampy_go

import (
	"context"
	"testing"

	"github.com/qvineox/vampy-go/vampy"
)

func TestProductListGet(t *testing.T) {
	c := testSpaceClient(t)
	ctx := testCtx(t)
	got, err := c.ProductListGet(ctx, ProductListGetOptions{Paging: vampy.Paging{Limit: 1}})
	if err != nil {
		t.Fatalf("ProductListGet: %v", err)
	}
	requireSlice(t, got.Items)
}

func TestProductPost(t *testing.T) {
	c := testSpaceClient(t)
	ctx := testCtx(t)
	name := uniqueName(t, "test-res")
	got, err := c.ProductPost(ctx, ProductPostOptions{Body: vampy.Product{Name: name}})
	if err != nil {
		t.Fatalf("ProductPost: %v", err)
	}
	requireSchema(t, got)
	t.Cleanup(func() {
		_ = c.ProductDelete(context.Background(), ProductDeleteOptions{ProductID: uuidFromString(t, string(got.ID))})
	})
}

func TestProductDetailsGet(t *testing.T) {
	c := testSpaceClient(t)
	ctx := testCtx(t)
	name := uniqueName(t, "test-res")
	created, err := c.ProductPost(ctx, ProductPostOptions{Body: vampy.Product{Name: name}})
	if err != nil {
		t.Fatalf("ProductPost: %v", err)
	}
	t.Cleanup(func() {
		_ = c.ProductDelete(context.Background(), ProductDeleteOptions{ProductID: uuidFromString(t, string(created.ID))})
	})
	got, err := c.ProductDetailsGet(ctx, ProductDetailsGetOptions{ProductID: uuidFromString(t, string(created.ID))})
	if err != nil {
		t.Fatalf("ProductDetailsGet: %v", err)
	}
	requireSchema(t, got)
}

func TestProductDelete(t *testing.T) {
	c := testSpaceClient(t)
	ctx := testCtx(t)
	name := uniqueName(t, "test-res")
	created, err := c.ProductPost(ctx, ProductPostOptions{Body: vampy.Product{Name: name}})
	if err != nil {
		t.Fatalf("ProductPost: %v", err)
	}
	if err := c.ProductDelete(ctx, ProductDeleteOptions{ProductID: uuidFromString(t, string(created.ID))}); err != nil {
		t.Fatalf("ProductDelete: %v", err)
	}
}

func TestProductPatch(t *testing.T) {
	c := testSpaceClient(t)
	ctx := testCtx(t)
	name := uniqueName(t, "test-res")
	name2 := uniqueName(t, "test-res-renamed")
	created, err := c.ProductPost(ctx, ProductPostOptions{Body: vampy.Product{Name: name}})
	if err != nil {
		t.Fatalf("ProductPost: %v", err)
	}
	t.Cleanup(func() {
		_ = c.ProductDelete(context.Background(), ProductDeleteOptions{ProductID: uuidFromString(t, string(created.ID))})
	})
	got, err := c.ProductPatch(ctx, ProductPatchOptions{ProductID: uuidFromString(t, string(created.ID)), Body: vampy.ProductEdit{Name: name2}})
	if err != nil {
		t.Fatalf("ProductPatch: %v", err)
	}
	requireSchema(t, got)
}

func TestProductUserGroupsGet(t *testing.T) {
	// requires an existing resource
	t.Skip("requires an existing resource")
}

func TestProductUsersSet(t *testing.T) {
	// requires an existing resource or external service
	t.Skip("requires an existing resource or external service")
}

func TestProductUsersUnset(t *testing.T) {
	// requires an existing resource or external service
	t.Skip("requires an existing resource or external service")
}

func TestProductUserGroupsSet(t *testing.T) {
	// requires an existing resource or external service
	t.Skip("requires an existing resource or external service")
}

func TestProductUserGroupsUnset(t *testing.T) {
	// requires an existing resource or external service
	t.Skip("requires an existing resource or external service")
}

func TestProductRelationsCreate(t *testing.T) {
	// requires an existing resource or external service
	t.Skip("requires an existing resource or external service")
}

func TestProductRelationsDelete(t *testing.T) {
	// requires an existing resource or external service
	t.Skip("requires an existing resource or external service")
}
