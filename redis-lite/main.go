package main

import "redis-lite/core"

func main() {
	core.NewServer(core.ServerOpts{Port: 8080}).Start()
}
