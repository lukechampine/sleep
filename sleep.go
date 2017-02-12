// Package sleep provides sleep-related utilities.
package sleep

import "time"

// Lightly waits until d has elapsed or wake strobes. In the former case, it
// returns false; in the latter, it returns true.
//
// Note that Lightly can be used with a Context by passing the Context's Done channel:
//
//   ctx, _ := context.WithTimeout(context.Background(), time.Second)
//   awoken := sleep.Lightly(2*time.Second, ctx.Done())
func Lightly(d time.Duration, wake <-chan struct{}) bool {
	timer := time.NewTimer(d)
	select {
	case <-timer.C:
		return false
	case <-wake:
		timer.Stop()
		return true
	}
}

// TickLightly returns a channel that strobes with the specified period until
// quit strobes.
func TickLightly(period time.Duration, quit <-chan struct{}) <-chan time.Time {
	ticker := time.NewTicker(period)
	ch := make(chan time.Time)
	go func() {
		for {
			select {
			case t := <-ticker.C:
				ch <- t
			case <-quit:
				ticker.Stop()
				close(ch)
				return
			}
		}
	}()
	return ch
}
