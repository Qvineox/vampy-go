package vampy_go

import (
	"context"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"

	"uuid"
)

// Environment variables used to configure integration tests. Tests skip
// themselves when these are unset so the suite stays green in CI without a
// live Vampy instance.
const (
	envBaseURL = "VAMPY_BASE_URL"
	envToken   = "VAMPY_TOKEN"
	envSpaceID = "VAMPY_SPACE_ID"
)

// testCtx returns a context with a generous timeout for integration calls.
func testCtx(t *testing.T) context.Context {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	t.Cleanup(cancel)
	return ctx
}

// testClient builds a client from VAMPY_BASE_URL and VAMPY_TOKEN, skipping the
// test when either is unset.
func testClient(t *testing.T) *Client {
	t.Helper()
	base := os.Getenv(envBaseURL)
	token := os.Getenv(envToken)
	if base == "" || token == "" {
		t.Skipf("set %s and %s to run integration tests", envBaseURL, envToken)
	}
	c, err := NewClient(base, WithAccessToken(token))
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	return c
}

// testSpaceID returns the configured test space UUID, skipping when unset.
// Space-scoped tests require a pre-existing space because Vampy exposes no
// space deletion endpoint, so tests must not create one they cannot remove.
func testSpaceID(t *testing.T) uuid.UUID {
	t.Helper()
	raw := os.Getenv(envSpaceID)
	if raw == "" {
		t.Skipf("set %s to a dedicated test space UUID", envSpaceID)
	}
	id, err := uuid.Parse(raw)
	if err != nil {
		t.Fatalf("invalid %s %q: %v", envSpaceID, raw, err)
	}
	return id
}

// testSpaceClient builds a client bound to the test space.
func testSpaceClient(t *testing.T) *Client {
	t.Helper()
	c := testClient(t)
	id := testSpaceID(t)
	sc, err := NewClient(c.baseURL.String(), WithAccessToken(c.accessToken), WithSpaceUUID(id))
	if err != nil {
		t.Fatalf("NewClient(space): %v", err)
	}
	return sc
}

// requireSchema validates that every required field of v (a response entity) is
// populated. Required fields are those whose json tag has no "omitempty" option.
func requireSchema(t *testing.T, v any) {
	t.Helper()
	rv := reflect.ValueOf(v)
	for rv.IsValid() && rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			t.Error("schema: nil pointer value")
			return
		}
		rv = rv.Elem()
	}
	if !rv.IsValid() || rv.Kind() != reflect.Struct {
		return
	}
	rt := rv.Type()
	for i := 0; i < rv.NumField(); i++ {
		f := rt.Field(i)
		if f.PkgPath != "" { // unexported
			continue
		}
		name, opts, ok := strings.Cut(f.Tag.Get("json"), ",")
		if !ok || name == "-" {
			continue
		}
		if strings.Contains(opts, "omitempty") { // optional field
			continue
		}
		if rv.Field(i).IsZero() {
			t.Errorf("schema: required field %q is zero", name)
		}
	}
}

// requireSlice validates that a decoded slice response is non-nil.
func requireSlice(t *testing.T, v any) {
	t.Helper()
	rv := reflect.ValueOf(v)
	for rv.IsValid() && rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	if !rv.IsValid() || rv.Kind() != reflect.Slice {
		t.Errorf("schema: expected slice, got %T", v)
		return
	}
	if rv.IsNil() {
		t.Error("schema: nil slice")
	}
}

// uniqueName returns a name with a random suffix so tests do not collide and
// can be re-run idempotently.
func uniqueName(t *testing.T, prefix string) string {
	t.Helper()
	return prefix + "-" + uuid.New().String()[:8]
}

// uuidFromString parses a string UUID, failing the test on invalid input.
func uuidFromString(t *testing.T, s string) uuid.UUID {
	t.Helper()
	id, err := uuid.Parse(s)
	if err != nil {
		t.Fatalf("parse uuid %q: %v", s, err)
	}
	return id
}
