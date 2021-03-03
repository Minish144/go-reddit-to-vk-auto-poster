package main

import (
	"fmt"
	"reddit-to-vk-auto-poster/src/reddit"
)

func main() {
	structure := reddit.Initialize()
	fmt.Print(structure)
}
