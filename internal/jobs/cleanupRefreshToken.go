package jobs

import (
	"context"
	"log"
	"time"

	"github.com/TrungNNg/rsshouse/internal/api"
)

func CleanUpRefreshToken(c *api.ApiConfig) {
	ticker := time.NewTicker(c.CleanupRefreshTokenInterval)
	log.Printf("CleanUpRefreshToken job run every %s\n", c.CleanupRefreshTokenInterval)
	for {
		<-ticker.C

		// Execute the cleanup
		err := c.DB.CleanupExpiredOrRevokedTokens(context.Background(), time.Now().UTC())
		if err != nil {
			log.Printf("Error cleaning up refresh tokens: %v", err)
		}
	}
}
