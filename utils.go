package main

import (
	"os"
	"path/filepath"
	"strings"
)

func GetExtension(fileName string) string {
	return strings.ToLower(filepath.Ext(fileName))
}

func GetCategory(ext string) string {
	categories := map[string]string{
		".jpg":  "images",
		".jpeg": "images",
		".png":  "images",
		".webp": "images",
		".avif": "images",
		".gif":  "images",

		".psd": "psds",

		".pdf":  "documents",
		".docx": "documents",
		".txt":  "documents",

		".mp4": "videos",
		".mkv": "videos",

		".mp3": "audios",
		".wav": "audios",

		".js":   "code",
		".html": "code",
		".css":  "code",
	}

	if category, exists := categories[ext]; exists {
		return category
	}
	return "others"
}

func CreateFolderIfNotExist(folderPath string){
	_,err := os.Stat(folderPath)

	if os.IsNotExist(err){
		os.Mkdir(folderPath,os.ModePerm)
	}
}