package config

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/joho/godotenv"
)

type BinPath struct {
	Paths []string
	// A string path and its entries
	Entries map[string][]string
}

func NewBinPath() (*BinPath, error) {

	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	osPaths := os.Getenv("PATH")
	fmt.Println(osPaths)

	pathList := strings.Split(osPaths, ":")

	fmt.Println(pathList)

	newBinPath := BinPath{
		Paths:   pathList,
		Entries: make(map[string][]string),
	}

	for _, path := range pathList {
		entries, err := os.ReadDir(path)
		if err != nil {
			continue
		}
		for _, entry := range entries {
			newBinPath.Entries[path] = append(newBinPath.Entries[path], entry.Name())
		}
	}
	return &newBinPath, nil

}

// Returns the path to an executable in specified paths
func (b *BinPath) GetPath(cmd string) (string, error) {
	for _, path := range b.Paths {
		entries := b.Entries[path]
		if slices.Contains(entries, cmd) {
			return fmt.Sprintf("%s/%s", path, cmd), nil
		}
	}
	return "", fmt.Errorf("%s: not found", cmd)
}
