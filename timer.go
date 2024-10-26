package cron

import "time"

// A Timer works just like [time.Timer], and cron schedules [Jobs] by waiting
// for the time event emitted by Timer. By default, [time.Timer] is used. You
// can also customize a Timer with [WithTimer] option to control scheduling
// behavior.
type Timer interface {
	C() <-chan time.Time
	Reset(d time.Duration) bool
	Stop() bool
}

// newStdTimer returns a Timer created using [time.Timer].
func newStdTimer(d time.Duration) Timer {
	return &stdTimer{t: time.NewTimer(d)}
}

type stdTimer struct {
	t *time.Timer
}

func (self *stdTimer) C() <-chan time.Time { return self.t.C }

func (self *stdTimer) Reset(d time.Duration) bool { return self.t.Reset(d) }

func (self *stdTimer) Stop() bool { return self.t.Stop() }
