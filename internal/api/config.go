package api

import (
	"time"

	"github.com/TrungNNg/rsshouse/internal/database"
	"github.com/mmcdole/gofeed"
)

type ApiConfig struct {
	DB                          *database.Queries
	Secret                      string
	Parser                      *gofeed.Parser
	FetchFeedInterval           time.Duration
	CleanupRefreshTokenInterval time.Duration
}
