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
