package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func getNextZipFilename(folder string, prefix string) string {
	counter := 1
	for {
		filename := fmt.Sprintf("%s%d.zip", prefix, counter)
		if _, err := os.Stat(filepath.Join(folder, filename)); os.IsNotExist(err) {
			return filename
		}
		counter++
	}
}

func zipLogFiles(source string, destination string, excludedDirs []string) error {
	zipfile, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	err = filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip excluded directories
		if info.IsDir() {
			for _, dir := range excludedDirs {
				if info.Name() == dir {
					return filepath.SkipDir
				}
			}
		}

		// Only include files that start with "log" and end with ".txt"
		if !info.IsDir() && strings.HasPrefix(info.Name(), "log") && strings.HasSuffix(info.Name(), ".txt") {
			header, err := zip.FileInfoHeader(info)
			if err != nil {
				return err
			}

			header.Name = strings.TrimPrefix(path, source+string(filepath.Separator))
			header.Method = zip.Deflate

			writer, err := archive.CreateHeader(header)
			if err != nil {
				return err
			}

			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = io.Copy(writer, file)
		}
		return err
	})

	return err
}
