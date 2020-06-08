package main

import (
	"log"
	"math/rand"
	"reflect"
	"time"

	"github.com/vvksh/amigo"
)

var watchers []Watcher

func register(watcher Watcher) {
	watchers = append(watchers, watcher)
}

func main() {
	for _, watcher := range watchers {
		go watch(watcher)
	}
	select {}
}

func watch(watcher Watcher) {
	log.Printf("watching now %s", reflect.TypeOf(watcher).String())
	for {
		updates := watcher.Check()
		handleUpdates(updates, watcher.SlackChannel())
		r := rand.Intn(500)
		time.Sleep(watcher.Interval() + time.Duration(r)*time.Millisecond)
	}
}

func handleUpdates(updates []string, channel string) {
	if len(updates) == 0 {
		log.Printf("\nNo new updates\n")
		return
	}
	for _, update := range updates {
		err := amigo.SendSlackNotification(update, channel)
		if err != nil {
			log.Printf(err.Error())
		}
	}

}
