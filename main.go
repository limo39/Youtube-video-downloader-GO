package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"

    "github.com/kkdai/youtube/v2"
)

func main() {
    // Initialize a YouTube client
    client := youtube.Client{}

    // Prompt user for YouTube link
    fmt.Print("Enter YouTube video link: ")
    reader := bufio.NewReader(os.Stdin)
    videoURL, _ := reader.ReadString('\n')
    videoURL = strings.TrimSpace(videoURL) // Remove any extra newline characters

    // Fetch video details
    video, err := client.GetVideo(videoURL)
    if err != nil {
        log.Fatalf("Error fetching video: %v", err)
    }

    // Format title for filename
    filename := sanitizeFilename(video.Title) + ".mp4"

    // Select a video format (e.g., the first MP4 format)
    var format *youtube.Format
    for _, f := range video.Formats {
        if strings.Contains(f.MimeType, "video/mp4") {
            format = &f
            break
        }
    }
    if format == nil {
        log.Fatalf("No compatible MP4 format found")
    }

    // Create a file to save the video
    file, err := os.Create(filename)
    if err != nil {
        log.Fatalf("Error creating file: %v", err)
    }
    defer file.Close()

    // Download the video
    fmt.Printf("Downloading %s as %s...\n", video.Title, filename)
    stream, _, err := client.GetStream(video, format)
    if err != nil {
        log.Fatalf("Error downloading video: %v", err)
    }
    _, err = file.ReadFrom(stream)
    if err != nil {
        log.Fatalf("Error saving video: %v", err)
    }

    fmt.Println("Download complete!")
}

// sanitizeFilename removes invalid characters for filenames
func sanitizeFilename(filename string) string {
    invalidChars := []string{"/", "\\", ":", "*", "?", "\"", "<", ">", "|"}
    for _, char := range invalidChars {
        filename = strings.ReplaceAll(filename, char, "_")
    }
    return filename
}
