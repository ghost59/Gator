package main

import ("context"
		"fmt"
		"gator/internal/database"
		"github.com/google/uuid"
		"time"
		"log"
		"strings")


func scrapeFeeds(s *state) {

	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		fmt.Printf("could not fetch feed: %v\n",err)
		return 
	}

	_ , err = s.db.MarkFeedFetched(context.Background(), nextFeed.ID)
	if err != nil {
		fmt.Println(err)
		return 
	}
	
	feedData, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		fmt.Println(err)
		return 
	}

		
	
	for _, data := range feedData.Channel.Item {
	parsedTime, err := time.Parse(time.RFC1123Z , data.PubDate)
	if err != nil {
		fmt.Println(err)
		continue
	}	

	postParams := database.CreatePostsParams{
			ID: uuid.New(),
			CreatedAt:time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Title: data.Title,
			Url: data.Link,
			Description: data.Description,
			PublishedAt: parsedTime,
			FeedID: nextFeed.ID,
			
	}

	 _, err = s.db.CreatePosts(context.Background(), postParams)
	 if err != nil {
	 	if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
	 		continue
	 	}
		log.Println("couldn't create post:", err)
	 	continue 
	 }
	 	 
		
	}

	
	
}
