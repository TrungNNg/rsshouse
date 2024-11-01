package api

import (
	"github.com/TrungNNg/rsshouse/internal/database"
	"github.com/mmcdole/gofeed"
)

type ApiConfig struct {
	DB     *database.Queries
	Secret string
	Parser *gofeed.Parser
}
