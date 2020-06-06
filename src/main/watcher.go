package main

import "time"

type Watcher interface {
	Check() []string
	Interval() time.Duration
	WebHookUrl() string
}
