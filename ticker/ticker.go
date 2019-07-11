package ticker

import (
	"time"
)

const (
	STOP        = "STOP"
	PAUSE       = "PAUSE"
	POMODORO    = "POMODORO"
	SHORT_BREAK = "SHORT_BREAK"
	LONG_BREAK  = "LONG_BREAK"
)

type Ticker struct {
	status         string
	prev_status    string
	pomodoro_count int
	counter        int
	time_ticker    *time.Ticker
}

func NewTicker() *Ticker {
	t := &Ticker{
		status:      STOP,
		time_ticker: time.NewTicker(time.Second),
	}
	return t
}

func (t *Ticker) Start() {
	is_continue := t.status == PAUSE
	if is_continue && t.prev_status != "" {
		t.status = t.prev_status
	} else {
		t.Pomodoro()
	}
}

func (t *Ticker) Pause() {
	t.status = PAUSE
}

func (t *Ticker) Pomodoro() {
	t.counter = 25 * 60
	t.status = POMODORO
}

func (t *Ticker) ShortBreak() {
	t.counter = 5 * 60
	t.status = POMODORO
}

func (t *Ticker) LongBreak() {
	t.counter = 15 * 60
	t.status = POMODORO
}
