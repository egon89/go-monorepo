package v1

import "time"

type WebsiteChecker func(string) bool

// race condition issue: map is not thread safe
// we can use the go Race Detector to identify race conditions
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		// anonymous functions
		go func() {
			results[url] = wc(url)
		}()
	}

	time.Sleep(2 * time.Second)

	return results
}
