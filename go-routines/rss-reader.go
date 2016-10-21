package main

import(
	"github.com/ajstarks/svgo"
	rss "github.com/jteeuwen/go-pkg-rss"
	"net/http"
	"log"
	"fmt"
	"strconv"
	"time"
	"os"
	"sync"
	"runtime"
)

type Feed struct {
	url string
	status int
	itemCount int
	itemsComplete bool
	index int
}

type FeedItem struct { 
	feedIndex int
	complete bool
	url string
}

var feeds []Feed
var height int
var width int
var colors []string
var startTime int64
var timeout int
var feedSpace int

var wg sync.WaitGroup

func grabFeed(feed *Feed, feedChain chan bool, osvg *svg.SVG) {
	startGrab := time.Now().Unix()
	startGrabSeconds := startGrab - startTime

	fmt.Println("Grabbing feed", feed.url, " at", startGrabSeconds, " second mark")

	if feed.status == 0 {
		fmt.Println("feed not yet read")
		feed.status = 1

		startX := int(startGrabSeconds * 33)
		startY := feedSpace * feed.feedIndex

		fmt.Println(startY)
		wg.Add(1)


		rssFeed := rss.New(timeout, true, channelHandler, itemsHandler)

		if err := rssFeed.Fetch(feed.url, nil); err != nil {
			fmt.Fprintf(os.Stderr, "[e] %s: %s", feed.url, err)
			return
		} else {
			endSec := time.Now().Unix()
			endX := int(endSec - startGrab)
			if endX == 0 {
				endX = 1
			}
			fmt.Println("Read feed in", endX, " seconds")
			osvg.Rect(startX, startY, endX, feedSpace, "fill:#000000;opacity:.4")
			wg.Wait()

			endGrab := time.Now().Unix()
			endGrabSeconds := endGrab - startTime
			feedEndX := int(endGrabSeconds * 33)

			osvg.Rect(feedEndX, startY, 1, feedSpace, "fill:#ff0000;opacity:.9")

			feedChain <- true
		}

	} else if feed.status == 1 {
		fmt.Println("feed already in progress")
	}
}






















