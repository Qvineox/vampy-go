package vampy_go

import (
	"context"
	"testing"

	"github.com/qvineox/vampy-go/vampy"
)

func TestQualityGateListGet(t *testing.T) {
	c := testSpaceClient(t)
	ctx := testCtx(t)
	got, err := c.QualityGateListGet(ctx, QualityGateListGetOptions{})
	if err != nil {
		t.Fatalf("QualityGateListGet: %v", err)
	}
	requireSlice(t, got)
}

func TestQualityGatePost(t *testing.T) {
	c := testSpaceClient(t)
	ctx := testCtx(t)
	name := uniqueName(t, "test-res")
	slug := uniqueName(t, "test-slug")
	got, err := c.QualityGatePost(ctx, QualityGatePostOptions{Body: vampy.QualityGate{Name: name, Slug: slug}})
	if err != nil {
		t.Fatalf("QualityGatePost: %v", err)
	}
	requireSchema(t, got)
	t.Cleanup(func() {
		_ = c.QualityGateDelete(context.Background(), QualityGateDeleteOptions{QualityGateID: uuidFromString(t, string(got.ID))})
	})
}

func TestQualityGateDetailsGet(t *testing.T) {
	c := testSpaceClient(t)
	ctx := testCtx(t)
	name := uniqueName(t, "test-res")
	slug := uniqueName(t, "test-slug")
	created, err := c.QualityGatePost(ctx, QualityGatePostOptions{Body: vampy.QualityGate{Name: name, Slug: slug}})
	if err != nil {
		t.Fatalf("QualityGatePost: %v", err)
	}
	t.Cleanup(func() {
		_ = c.QualityGateDelete(context.Background(), QualityGateDeleteOptions{QualityGateID: uuidFromString(t, string(created.ID))})
	})
	got, err := c.QualityGateDetailsGet(ctx, QualityGateDetailsGetOptions{QualityGateID: uuidFromString(t, string(created.ID))})
	if err != nil {
		t.Fatalf("QualityGateDetailsGet: %v", err)
	}
	requireSchema(t, got)
}

func TestQualityGateDelete(t *testing.T) {
	c := testSpaceClient(t)
	ctx := testCtx(t)
	name := uniqueName(t, "test-res")
	slug := uniqueName(t, "test-slug")
	created, err := c.QualityGatePost(ctx, QualityGatePostOptions{Body: vampy.QualityGate{Name: name, Slug: slug}})
	if err != nil {
		t.Fatalf("QualityGatePost: %v", err)
	}
	if err := c.QualityGateDelete(ctx, QualityGateDeleteOptions{QualityGateID: uuidFromString(t, string(created.ID))}); err != nil {
		t.Fatalf("QualityGateDelete: %v", err)
	}
}

func TestQualityGatePatch(t *testing.T) {
	c := testSpaceClient(t)
	ctx := testCtx(t)
	name := uniqueName(t, "test-res")
	slug := uniqueName(t, "test-slug")
	name2 := uniqueName(t, "test-res-renamed")
	created, err := c.QualityGatePost(ctx, QualityGatePostOptions{Body: vampy.QualityGate{Name: name, Slug: slug}})
	if err != nil {
		t.Fatalf("QualityGatePost: %v", err)
	}
	t.Cleanup(func() {
		_ = c.QualityGateDelete(context.Background(), QualityGateDeleteOptions{QualityGateID: uuidFromString(t, string(created.ID))})
	})
	got, err := c.QualityGatePatch(ctx, QualityGatePatchOptions{QualityGateID: uuidFromString(t, string(created.ID)), Body: vampy.QualityGateEdit{Name: name2}})
	if err != nil {
		t.Fatalf("QualityGatePatch: %v", err)
	}
	requireSchema(t, got)
}

func TestQualityGateResultListGet(t *testing.T) {
	// requires an existing resource
	t.Skip("requires an existing resource")
}

func TestQualityGateResultGetDetails(t *testing.T) {
	// requires an existing resource
	t.Skip("requires an existing resource")
}
