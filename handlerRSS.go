package main


import ("net/http"
		"io"
		"encoding/xml"
		"html"
		"context"
		"fmt"
		"time"
		)

type RSSFeed struct {
	Channel struct {
		Title string `xml:"title"`
		Link string `xml:"link"`
		Description string `xml:"description"`
		Item []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title string `xml:"title"`
	Link string `xml:"link"`
	Description string `xml:"description"`
	PubDate string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error){
	rss := &RSSFeed{}
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil) 
	if err != nil {
		return &RSSFeed{}, err
	}
	req.Header.Set("User-Agent","gator")
	client := &http.Client{}
	
	res, err := client.Do(req)
	if err != nil {
		return &RSSFeed{}, err
	}
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return &RSSFeed{}, err
	}
	defer res.Body.Close()

	err = xml.Unmarshal(b,rss)

	if err != nil {
		return &RSSFeed{}, err
	}

	
	
	

	for i, de := range rss.Channel.Item{
		de.Description = html.UnescapeString(de.Description)
		de.Title = html.UnescapeString(de.Title)
		rss.Channel.Item[i] = de
	  		
	}

 
	


	return rss, nil

	
}


func Agg(s *state, cmd command) error{
	if len(cmd.args) != 1 {
		return fmt.Errorf("THis is empty")
	}

	timeBetweenReqs, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return fmt.Errorf("invalided duration: %w", err) 
	}
		

ticker := time.NewTicker(timeBetweenReqs)

fmt.Printf("Collecting feeds every %s\n", timeBetweenReqs)
for ; ; <-ticker.C {
	scrapeFeeds(s)
	}

 return nil
}  
