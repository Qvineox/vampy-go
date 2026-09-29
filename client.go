package vampy_go

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

// Client is a Vampy ASPM API client.
type Client struct {
	httpClient *http.Client
	baseURL    *url.URL

	accessToken string
	userAgent   string

	spaceID uuid.UUID
}

// Option configures the Client.
type Option func(*Client) error

// NewClient creates a Vampy API client for the given base URL (for example
// "https://vampy.example.com/api").
func NewClient(baseURL string, opts ...Option) (*Client, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("vampy: invalid base URL: %w", err)
	}

	c := &Client{
		baseURL:    u,
		httpClient: &http.Client{},
		userAgent:  "vampy-go",
	}

	for _, o := range opts {
		if err := o(c); err != nil {
			return nil, err
		}
	}
	return c, nil
}

// WithAccessToken sets the bearer access token sent with every request.
func WithAccessToken(token string) Option {
	return func(c *Client) error { c.accessToken = token; return nil }
}

// WithSpaceUUID sets the default Vampy space UUID for all interactions.
func WithSpaceUUID(id uuid.UUID) Option {
	return func(c *Client) error { c.spaceID = id; return nil }
}

// WithHTTPClient sets a custom http.Client.
func WithHTTPClient(client *http.Client) Option {
	return func(c *Client) error {
		if client == nil {
			return fmt.Errorf("vampy: nil http client")
		}
		c.httpClient = client
		return nil
	}
}

// WithHTTPProxy routes requests through the given HTTP proxy URL.
func WithHTTPProxy(proxyURL string) Option {
	return func(c *Client) error {
		u, err := url.Parse(proxyURL)
		if err != nil {
			return fmt.Errorf("vampy: invalid proxy URL: %w", err)
		}
		tr := http.DefaultTransport.(*http.Transport).Clone()
		tr.Proxy = http.ProxyURL(u)
		c.httpClient.Transport = tr
		return nil
	}
}

// WithUserAgent sets the User-Agent header.
func WithUserAgent(ua string) Option {
	return func(c *Client) error { c.userAgent = ua; return nil }
}

// WithTimeout sets the client request timeout.
func WithTimeout(d time.Duration) Option {
	return func(c *Client) error { c.httpClient.Timeout = d; return nil }
}

// Error is an API error response.
type Error struct {
	StatusCode int
	Status     string
	Body       []byte
}

func (e *Error) Error() string {
	return fmt.Sprintf("vampy: %s (status %d): %s", e.Status, e.StatusCode, string(e.Body))
}

func (c *Client) newRequest(ctx context.Context, method, path string, query url.Values, body io.Reader, contentType string) (*http.Request, error) {
	base := strings.TrimRight(c.baseURL.String(), "/")
	full := base + "/" + strings.TrimLeft(path, "/")
	u, err := url.Parse(full)
	if err != nil {
		return nil, err
	}
	if query != nil {
		u.RawQuery = query.Encode()
	}
	req, err := http.NewRequestWithContext(ctx, method, u.String(), body)
	if err != nil {
		return nil, err
	}
	if c.accessToken != "" {
		req.Header.Set("Authorization", "Bearer "+c.accessToken)
	}
	if c.userAgent != "" {
		req.Header.Set("User-Agent", c.userAgent)
	}
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}
	return req, nil
}

func (c *Client) do(req *http.Request) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 400 {
		b, _ := io.ReadAll(io.LimitReader(resp.Body, 1<<20))
		resp.Body.Close()
		return nil, &Error{StatusCode: resp.StatusCode, Status: resp.Status, Body: b}
	}
	return resp, nil
}

func (c *Client) doJSON(ctx context.Context, method, path string, query url.Values, body, out interface{}) error {
	var rdr io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return err
		}
		rdr = bytes.NewReader(b)
	}
	req, err := c.newRequest(ctx, method, path, query, rdr, "application/json")
	if err != nil {
		return err
	}
	resp, err := c.do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if out == nil || resp.StatusCode == 204 {
		return nil
	}
	return json.NewDecoder(resp.Body).Decode(out)
}

func (c *Client) doBinary(ctx context.Context, method, path string, query url.Values, body interface{}) ([]byte, error) {
	var rdr io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		rdr = bytes.NewReader(b)
	}
	req, err := c.newRequest(ctx, method, path, query, rdr, "application/json")
	if err != nil {
		return nil, err
	}
	resp, err := c.do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func (c *Client) doMultipart(ctx context.Context, method, path string, query url.Values, body, out interface{}) error {
	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)
	if body != nil {
		if err := writeMultipart(w, body); err != nil {
			return err
		}
	}
	w.Close()
	req, err := c.newRequest(ctx, method, path, query, &buf, w.FormDataContentType())
	if err != nil {
		return err
	}
	resp, err := c.do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if out == nil || resp.StatusCode == 204 {
		return nil
	}
	return json.NewDecoder(resp.Body).Decode(out)
}

func writeMultipart(w *multipart.Writer, body interface{}) error {
	v := reflect.ValueOf(body)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		ft := t.Field(i)
		if !f.CanInterface() {
			continue
		}
		tag := ft.Tag.Get("json")
		name := strings.Split(tag, ",")[0]
		if name == "" || name == "-" {
			continue
		}
		if f.Kind() == reflect.Ptr {
			if f.IsNil() {
				continue
			}
			f = f.Elem()
		}
		if f.Kind() == reflect.Slice && f.Type().Elem().Kind() == reflect.Uint8 {
			part, err := w.CreateFormFile(name, name)
			if err != nil {
				return err
			}
			if _, err := part.Write(f.Bytes()); err != nil {
				return err
			}
			continue
		}
		var s string
		switch f.Kind() {
		case reflect.String:
			s = f.String()
		case reflect.Bool:
			s = strconv.FormatBool(f.Bool())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			s = strconv.FormatInt(f.Int(), 10)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			s = strconv.FormatUint(f.Uint(), 10)
		case reflect.Float32, reflect.Float64:
			s = strconv.FormatFloat(f.Float(), 'g', -1, 64)
		default:
			b, err := json.Marshal(f.Interface())
			if err != nil {
				return err
			}
			s = string(b)
		}
		w.WriteField(name, s)
	}
	return nil
}
