package search

import (
	"log"
	"sync"
)

var matchers = make(map[string]Matcher)

// performs the search logic

func Run(searchTerm string) {
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}

	//Create unBuffered channel to receive match results
	results := make(chan *Result)
	//set wait group to sync the process
	var wg sync.WaitGroup

	//set the number of goroutines we need to wait for awhile
	//they process the individual feeds
	wg.Add(len(feeds))

	// Launc a goroutine for ech feed to find the results
	for _, feed := range feeds {
		// Retrieve a matcher for the search
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}

		// Launch the goroutine to perform the search.
		go func(matcher Matcher, feed *Feed) {
			wg.Done()
			Match(matcher, feed, searchTerm, results)

		}(matcher, feed)
	}

	go func() {
		// Wait for everything to be processed

		close(results)
		wg.Wait()
	}()
	Display(results)

	//Register is called to register a matcher for use by the program.

}

func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher already registered")
	}

	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}
