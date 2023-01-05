package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
)

func main() {
    // Get the YouTube video URL
    var videoURL string
    fmt.Print("Enter the URL of the YouTube video: ")
    fmt.Scanln(&videoURL)

    // Send an HTTP request to the URL
    resp, err := http.Get(videoURL)
    if err != nil {
        fmt.Println("Error downloading video:", err)
        return
    }
    defer resp.Body.Close()

    // Create a file to save the video
    file, err := os.Create("video.mp4")
    if err != nil {
        fmt.Println("Error creating file:", err)
        return
    }
    defer file.Close()

    // Copy the video data from the HTTP response to the file
    _, err = io.Copy(file, resp.Body)
    if err != nil {
        fmt.Println("Error downloading video:", err)
        return
    }

    fmt.Println("Video downloaded successfully!")
}
