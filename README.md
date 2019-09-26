[![Build Status](https://circleci.com/gh/lucasdss/alieninvasion/tree/master.svg?style=svg)](https://circleci.com/gh/lucasdss/alieninvasion/tree/master)
[![Go Report Card](https://goreportcard.com/badge/github.com/lucasdss/alieninvasion)](https://goreportcard.com/report/github.com/lucasdss/alieninvasion)
[![Go Doc](https://godoc.org/github.com/lucasdss/alieninvasion?status.svg)](https://godoc.org/github.com/lucasdss/alieninvasion)

# alieninvasion

Alien Invasion is a Golang exercise.

The more detailed description of this exercise can be found at [Alien Invasion](/assets/Alien%20Invasion.pdf).

## Assumptions

It's not specified if goroutines must be used, neither the use of channels.

So there is a basic version avoiding concurrency and mutexes and the default version using goroutines and channels.


## Build the code

`go build ./cmd/...` 

or

`make build`


### Run

```
$ ./alieninvasion -h
Usage of ./alieninvasion:
  -c    version with goroutines and channels. (default true)
  -map string
        path to world map file. (default "assets/world_map.txt")
  -n int
        number os aliens invading the planet. (default 5)
```

### Tests and Build

`make`

or

`go test -v -race -cover ./...`

and then

`gp build ./cmd/...`
