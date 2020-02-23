package timeutils

import "time"

func Every(d time.Duration, f func()) {
	go (func() {
		f()

		for range time.Tick(d) {
			f()
		}
	})()
}
