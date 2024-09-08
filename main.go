package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var mediaInfoCommand string

func init() {
	if _, err := exec.LookPath("mediainfo"); err == nil {
		mediaInfoCommand = "mediainfo"
	} else if _, err := exec.LookPath("MediaInfo"); err == nil {
		mediaInfoCommand = "MediaInfo"
	} else {
		fmt.Println("Error: MediaInfo command not found. Please ensure it's installed and in your PATH.")
		os.Exit(1)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: video-renamer <folder_path>")
		os.Exit(1)
	}

	folderPath := os.Args[1]
	err := filepath.Walk(folderPath, processFile)
	if err != nil {
		fmt.Printf("Error walking through directory: %v\n", err)
		os.Exit(1)
	}
}

func processFile(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if info.IsDir() {
		return nil
	}

	if !strings.Contains(info.Name(), "tfpdl-") {
		return nil
	}

	movieName, err := getMovieName(path)
	if err != nil {
		fmt.Printf("Error getting movie name for %s: %v\n", path, err)
		return nil
	}

	if movieName == "" {
		fmt.Printf("No movie name found for %s\n", path)
		return nil
	}

	fmt.Printf("Movie name for %s: %s\n", path, movieName)

	newPath := filepath.Join(filepath.Dir(path), movieName+filepath.Ext(path))
	err = os.Rename(path, newPath)
	if err != nil {
		fmt.Printf("Error renaming %s to %s: %v\n", path, newPath, err)
		return nil
	}

	fmt.Printf("Renamed %s to %s\n", path, newPath)
	return nil
}

func getMovieName(filePath string) (string, error) {
	cmd := exec.Command(mediaInfoCommand, "--Inform=General;%Movie%", filePath)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	name := strings.TrimSpace(string(output))
	name = strings.TrimPrefix(name, "TFPDL - ")
	return strings.TrimSpace(name), nil
}
