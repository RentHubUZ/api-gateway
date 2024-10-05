package handler

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// UploadMedia godoc
// @Summary Fayl yuklash
// @Description Fayl yuklash uchun endpoint
// @Tags Media
// @Accept multipart/form-data
// @Produce json
// @Security     ApiKeyAuth
// @Param file formData file true "Yuklanadigan fayl"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/upload/imagesandvedio [post]
func (h Handler) UploadMedia(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Faylni olishda xatolik: " + err.Error(),
		})
		return
	}
	fmt.Println("1")

	fileExt := filepath.Ext(file.Filename)

	newFile := uuid.NewString() + fileExt

	mediaDir := "./media"
	err = os.MkdirAll(mediaDir, os.ModePerm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Serverda katalog yaratishda xatolik: " + err.Error(),
		})
		return
	}
	fmt.Println("2")

	filePath := filepath.Join(mediaDir, newFile)
	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Faylni saqlashda xatolik: " + err.Error(),
		})
		return
	}
	fmt.Println("3")

	minioClient, err := minio.New("minio:9000", &minio.Options{
		Creds:  credentials.NewStaticV4("minio", "minioadmin", ""),
		Secure: false,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "MinIO bilan bog'lanishda xatolik: " + err.Error(),
		})
		return
	}
	fmt.Println("4")

	var bucketName string
	contentType := "application/octet-stream"

	switch fileExt {
	case ".jpg", ".jpeg", ".png":
		bucketName = "images"
		if fileExt == ".jpg" || fileExt == ".jpeg" {
			contentType = "image/jpeg"
		} else if fileExt == ".png" {
			contentType = "image/png"
		}
	case ".mp4":
		bucketName = "videos"
		contentType = "video/mp4"
	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Yaroqsiz fayl turi",
		})
		return
	}
	fmt.Println("5")

	exists, err := minioClient.BucketExists(context.Background(), bucketName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "MinIO bilan bucketni tekshirishda xatolik: " + err.Error(),
		})
		return
	}
	fmt.Println("6")
	if !exists {
		err = minioClient.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Bucket yaratishda xatolik: " + err.Error(),
			})
			return
		}
	}
	fmt.Println("7")

	_, err = minioClient.FPutObject(context.Background(), bucketName, newFile, filePath, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "MinIO ga faylni yuklashda xatolik: " + err.Error(),
		})
		return
	}
	fmt.Println("8")

	objUrl := fmt.Sprintf("http://minio:9000/%s/%s", bucketName, newFile)
	c.JSON(http.StatusOK, gin.H{
		"file_url": objUrl,
	})
}
