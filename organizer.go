package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func OrganizeFiles(path string) {
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		fileName := file.Name()
		fullPath := filepath.Join(path, fileName)

		ext := GetExtension(fileName)

		category := GetCategory(ext)

		destFolder := filepath.Join(path,category)

		CreateFolderIfNotExist(destFolder)

		destPath := filepath.Join(destFolder,fileName)

		err := os.Rename(fullPath,destPath)

		if err!=nil{
			fmt.Println("Error moving file ",err)
			return
		}
		fmt.Println("Moved: ",fileName, " -> ",category)

	}
}
