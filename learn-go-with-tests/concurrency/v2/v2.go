package v2

type WebsiteChecker func(string) bool

// anonymous within the struct
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		// anonymous functions
		go func() {
			// using channel to send the result back to the main goroutine
			resultChannel <- result{url, wc(url)}
		}()
	}

	for i := 0; i < len(urls); i++ {
		// receiving from the channel
		r := <-resultChannel

		// writing to the map (now safe because only the main goroutine writes to it)
		results[r.string] = r.bool
	}

	return results
}
