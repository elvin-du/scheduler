package main

import (
	"scheduler/task"
)

func main() {
	go task.StartAll()
	select {}
}
