package utils

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/minio/minio-go/v7"
)

func UploadHandler(w http.ResponseWriter, r *http.Request, client *minio.Client) {
	// Parse form multipart
	err := r.ParseMultipartForm(5 << 20) // 5 MB max
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "File is required", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Get optional folder name from query (default "general")
	folder := r.FormValue("folder")
	if folder == "" {
		folder = "general"
	}

	// Prepare object name with folder
	objectName := folder + "/" + handler.Filename
	bucketName := "belajar"

	// Upload to MinIO
	uploadInfo, err := client.PutObject(context.Background(), bucketName, objectName, file, handler.Size, minio.PutObjectOptions{
		ContentType: handler.Header.Get("Content-Type"),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := fmt.Sprintf("✅ Uploaded %s (%d bytes)", uploadInfo.Key, uploadInfo.Size)
	log.Println(resp)
	w.Write([]byte(resp))
}
