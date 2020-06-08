# Nightswatch

This repo contains code to monitor websites/Rss Feeds etc at a given cadence and send any updates as a slack message.

# Includes trackers for 
- Arxiv feed {CS (databases and distributed systems)}
- Trending Stocks (gainers)
- Tech/pharma News 

To implement your own tracker, follow the examples. You will have to extend `Watcher` interface; will need to implement

```go
Check() //Called at regular interval by main which returns array of strings; each will be posted to slack

Interval() // return time.Duration object which tells main how long to sleep between successive calls to Check()

SlackChannel() // return which slack channel to post; note for this to work, you will need to install "Incoming webhooks" app to your workplace
```

# How to run it

- Insert api keys and webhook urls in Dockerfile_template
- `mv Dockerfile_template Dockerfile`
- Build and deploy docker image

```bash
docker build . -t nightswatch:latest
docker run   -d --name nightswatch nightswatch
```
