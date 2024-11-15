# YouTube Video Downloader in Go

A simple command-line application written in Go that allows users to download YouTube videos. The script prompts the user for a YouTube video URL, retrieves the video details, and downloads it in MP4 format, saving it with the video’s title as the filename.

## Table of Contents

- [Features](#features)
- [Requirements](#requirements)
- [Installation](#installation)
- [Usage](#usage)


## Features

- Prompts the user to enter a YouTube video URL.
- Retrieves the video title and uses it as the filename.
- Downloads the video in MP4 format (chooses the first available quality).
- Handles special characters in video titles to create valid filenames.

## Requirements

- [Go](https://golang.org/doc/install) 1.16 or higher

## Installation

1. Clone the repository or create a new directory for the script:

   ```sh
   git clone https://github.com/yourusername/yt-downloader-go.git
   cd yt-downloader-go
2. Initialize a Go module in the directory (if not already set up):

   ```sh
   go mod init yt_downloader

3. Install the required YouTube download package:
   ```sh
   go get github.com/kkdai/youtube

## Usage
1. Run the script:
   ```sh
   go run main.go

2. When prompted, paste in the YouTube video URL and press Enter.

3. The script will:

Retrieve the video title.
Sanitize the title to create a valid filename.
Download the video in MP4 format, saving it as <video_title>.mp4 in the current directory.

4. The download progress and completion message will be displayed in the terminal.

