package main

import (
	"io/ioutil"
	"log"
	"os/exec"
)

func main() {
	file, err := ioutil.ReadFile("path.txt")
	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command("explorer.exe", string(file))
	cmd.Run()
}
