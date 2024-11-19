package helpers

import (
	"mime/multipart"
	"net/http"
)

func IsImageFile(file *multipart.FileHeader) bool {
    openedFile, err := file.Open()
    if err != nil {
        return false
    }
    defer openedFile.Close()

    
    buffer := make([]byte, 512)
    _, err = openedFile.Read(buffer)
    if err != nil {
        return false
    }

    mimeType := http.DetectContentType(buffer)
    return mimeType == "image/jpeg" || mimeType == "image/png" || mimeType == "image/jpg"
}