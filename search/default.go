package search

// DefaultMatcher implements  the default matcher
type defaultMatcher struct{}

// init register the default mathcer with the program
func init() {
	var matcher defaultMatcher
	Register("default", matcher)
}

// Search implement the behavior for the default matcher.
func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, err) {
	return nil, nil
}
