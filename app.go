package main

import (
    "io"
    "log"
    "net/http"
    "os"
)

func main() {
    videoPath := "./febe802d7618d6cd2e2b834f4ce824e053429593-480p.mp4"

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        // Open the video file
        videoFile, err := os.Open(videoPath)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        defer videoFile.Close()

        // Set the content type header
        w.Header().Set("Content-Type", "video/mp4")

        // Serve the video file
        _, err = io.Copy(w, videoFile)
        if err != nil {
            log.Println(err)
        }
    })

    log.Println("Listening on :8000...")
    log.Fatal(http.ListenAndServe(":8000", nil))
}