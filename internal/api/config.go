package api

import "github.com/TrungNNg/rsshouse/internal/database"

type ApiConfig struct {
	DB     *database.Queries
	Secret string
}
