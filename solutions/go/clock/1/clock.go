package clock

import "fmt"

type Clock struct {
	minutes int
}

func New(h, m int) Clock {
	minutes := (((h*60 + m) % 1440) + 1440) % 1440
	return Clock{minutes: minutes}
}

func (c Clock) Add(m int) Clock {
	minutes := (((c.minutes + m) % 1440) + 1440) % 1440
	return Clock{minutes: minutes}
}

func (c Clock) Subtract(m int) Clock {
	minutes := (((c.minutes - m) % 1440) + 1440) % 1440
	return Clock{minutes: minutes}
}

func (c Clock) String() string {
	minutes := ((c.minutes % 1440) + 1440) % 1440
	hours := minutes / 60
	displayMinutes := minutes % 60
	return fmt.Sprintf("%02d:%02d", hours, displayMinutes)
}
