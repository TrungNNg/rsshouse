package api

import (
	"net/http"

	"github.com/TrungNNg/rsshouse/internal/database"
	"github.com/google/uuid"
)

// return a list of feeds of input topics, filter by lang and feed type
// sort by follow count and created time of feed
func (c *ApiConfig) SearchFeed(w http.ResponseWriter, r *http.Request) {
	input := struct {
		Topics   []string
		Lang     string
		FeedType string
		Offset   int
		Limit    int
	}{
		Topics:   []string{"space"}, // this will be topics id -> convert to []uuid.UUID
		Lang:     "EN",
		FeedType: "Rss",
		Offset:   0,
		Limit:    10,
	}

	topicIDs := []uuid.UUID{}
	for _, tid := range input.Topics {
		id, _ := uuid.Parse(tid)
		topicIDs = append(topicIDs, id)
	}

	// get feeds with given topic ids
	feeds, err := c.DB.GetFeedsByTopicID(r.Context(), database.GetFeedsByTopicIDParams{
		Column1:  topicIDs,
		Lang:     input.Lang,
		FeedType: input.FeedType,
		Offset:   int32(input.Offset),
		Limit:    int32(input.Limit),
	})
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Couldn't load feed", err)
		return
	}

	for _, feed := range feeds {
		println(feed.Title)
	}
}
