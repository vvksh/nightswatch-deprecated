package main

import "time"

type Watcher interface {
	Check() []Update
	Interval() time.Duration
}

type Update struct {
	Date  string
	Title string
	URL   string
}
