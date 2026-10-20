package vampy_go

import (
	"fmt"
	"log/slog"
	"os"
	"uuid"
)

var testClient *Client

func init() {
	slog.SetDefault(slog.New(slog.NewTextHandler(
		os.Stdout,
		&slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	)))

	var (
		err error

		vampyURL         = os.Getenv("VAMPY_BASE_API_URL")
		vampySpaceUUID   = os.Getenv("VAMPY_SPACE_UUID")
		vampyAccessToken = os.Getenv("VAMPY_ACCESS_TOKEN")
	)

	slog.Info("initializing vampy test client...",
		slog.String("api_url", vampyURL),
		slog.String("space_uuid", vampySpaceUUID),
		slog.Bool("access_token_set", vampyAccessToken != ""),
	)

	if vampyURL == "" {
		panic("VAMPY_BASE_API_URL environment variable is not set")
	}

	if vampySpaceUUID == "" {
		panic("VAMPY_SPACE_UUID environment variable is not set")
	}

	if vampyAccessToken == "" {
		panic("VAMPY_ACCESS_TOKEN environment variable is not set")
	}

	testClient, err = NewClient(vampyURL, []Option{
		WithAccessToken(vampyAccessToken),
		WithSpaceUUID(uuid.MustParse(vampySpaceUUID)),
	}...)

	if err != nil {
		panic(fmt.Sprintf("failed to create vampy test client: %s", err.Error()))
	}
}
