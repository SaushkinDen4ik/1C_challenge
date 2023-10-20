package main

import (
	"flag"
	"os"
)

type FileAndContent struct {
	Name    string
	Content string
}

func GetFilesNamesInDirectory(dirPath string) ([]string, error) {
	files, err := os.ReadDir(dirPath)
	filesNames := make([]string, 0)

	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if file.Name() != "" {
			filesNames = append(filesNames, file.Name())
		}
	}
	return filesNames, nil
}

var (
	precision = flag.Float64("precision", 0.5, "Coincidence precision")
)

func main() {
	flag.Parse()

	dir := "./files"
	filesNames, _ := GetFilesNamesInDirectory(dir)
	filesContent := make([]FileAndContent, 0)

	for _, name := range filesNames {
		copy_dir := dir
		copy_dir += "/"
		copy_dir += name
		content, _ := os.ReadFile(copy_dir)
		filesContent = append(filesContent, FileAndContent{Name: name, Content: string(content)})
	}

	for i := 0; i < len(filesContent); i++ {
		for j := i; j < len(filesContent); j++ {
			if j == i {
				continue
			}
			if len(filesContent[i].Content) >= len(filesContent[j].Content) {
				CompareFiles(*precision, filesContent[i], filesContent[j])
			} else {
				CompareFiles(*precision, filesContent[j], filesContent[i])
			}
		}
	}
}
