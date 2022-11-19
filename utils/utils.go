package utils

import (
	"bufio"
	"log"
	"os"
)

func ReadFile(fPath string, line func(line string)) {
	f, err := os.Open(fPath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
