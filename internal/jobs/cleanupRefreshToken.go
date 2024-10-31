package jobs

import (
	"context"
	"log"
	"time"

	"github.com/TrungNNg/rsshouse/internal/api"
)

func CleanUpRefreshToken(c *api.ApiConfig, interval time.Duration) {
	ticker := time.NewTicker(interval)
	log.Printf("CleanUpRefreshToken job running every %s\n", interval)
	for {
		<-ticker.C

		// Execute the cleanup
		err := c.DB.CleanupExpiredOrRevokedTokens(context.Background(), time.Now().UTC())
		if err != nil {
			log.Printf("Error cleaning up refresh tokens: %v", err)
		}
	}
}
