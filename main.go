package main

import (
	"fmt"

	"github.com/hpcloud/tail"
)

func main() {
	fmt.Println("Starting---")
	t, err := tail.TailFile("/var/log/apache2/access.log", tail.Config{Follow: true, ReOpen: true})
	if err != nil {
		fmt.Println(err)
		panic(1)
	}
	for line := range t.Lines {
		fmt.Println(line.Text)
	}
	fmt.Println("Finished---")
}
