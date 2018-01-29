package search

//Result contains the result of a search
type Result struct {
	Field   string
	Content string
}

// Matcher defines the behavior required by type that want
// to implement a new search type.
type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}
// Match is launched as a goroutine for each individual feed to run
// searcher concurrently
func Match(matchers Matcher, feed *Feed, searchTerm string, results chan <- *Results) {

	// Perform the search against the specified mathcer.
	searchResults, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Println(err)
		return
	}

	//Write the results to the channel.
	for _, result  := range searchResults {
		results <- result
	}

	//Display writes results to the terminal window as they 
	// are received by the individual goroutines
	func Display(results chan *Result) {
		// The channel blocks until a result is written to the channel
		// Once the channel is Closed the for loop terminates.
		for result := range results {
			fmt.Printf("%s:\n%s\n\n", result.Field, result.Content)
		}
	}
}
