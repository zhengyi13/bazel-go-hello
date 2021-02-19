package main

import (
	"fmt"

	"github.com/golang/glog"
	"flag"
)

func main() {
	var name = flag.String("name", "world", "to whom do you wish to say hello?")
	flag.Parse()
	fmt.Printf("Hello %s!\n", *name)
	glog.Info("Prepare to repel boarders")
}
