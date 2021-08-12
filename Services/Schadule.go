package Services

import "time"


// Schedule function in interval
func Schedule(process func(), interval time.Duration) *time.Ticker {
	ticker := time.NewTicker(interval)
	go func() {
		for range ticker.C {
			process()
		}
	}()
	return ticker
}