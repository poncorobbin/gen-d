package main

import (
	"fmt"
	"os"
	"path/filepath"
	// "syscall"
	"time"
)

func main() {
	currDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	year := time.Now().Year()
	month := time.Now().Month()
	day := time.Now().Day()
	path := filepath.Join(currDir, fmt.Sprintf("%d_%d/%d.txt", year, month, day))

	if _, err = os.Stat(path); os.IsNotExist(err) {
		// syscall.Umask(0)
		if err = os.MkdirAll(filepath.Dir(path), 0770); err != nil {
			fmt.Printf("%s", err)
			return
		}

		file, err := os.Create(path)
		if err != nil {
			fmt.Printf("%s", err)
			return
		}
		defer file.Close() // created file always open so we need to close it
 
		fmt.Printf("===> File successfully created %s", path)
	} else {
		fmt.Printf("%s already exists", path)
	}
}
