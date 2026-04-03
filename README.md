# Sortify CLI

A lightweight and efficient CLI tool built in Go to automatically organize files into categorized folders based on their type.

## Features

* Automatically sorts files into categories (images, documents, videos, etc.)
* Creates folders dynamically if they don’t exist
* Clean and fast file operations
* Simple CLI-based interaction

## Categories Supported

* Images (.jpg, .png, .jpeg)
* Documents (.pdf, .docx, .txt)
* Videos (.mp4, .mkv)
* PSDs (.psd)
* Audio (.mp3,.wav)
* Others (fallback category)

## Usage

```bash
go run .
```

Enter the folder path when prompted.

## 🛠 Tech Stack

* Go (Golang)
* Standard Library (os, filepath, strings)

## Note

This tool moves files from the original directory into categorized folders. Use it on test directories before applying to important data.

---

> Built as part of my Go learning journey, evolving into a real utility tool.
