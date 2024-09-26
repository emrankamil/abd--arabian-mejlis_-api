package controller

import (
	"abduselam-arabianmejlis/bootstrap"
	"abduselam-arabianmejlis/domain"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)


type UploadController struct {
	productUseCase domain.ProductUseCase
	Env 		*bootstrap.Env
}

func NewUploadController(u domain.ProductUseCase, env *bootstrap.Env) *UploadController {
	return &UploadController{
		productUseCase: u,
		Env: env,
	}
}

func (uc *UploadController) HandleUpload(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(400, gin.H{"message": "Unable to parse form"})
		return
	}

	files := form.File["files"]
	if len(files) == 0 {
		c.JSON(400, gin.H{"message": "No files uploaded"})
		return
	}

	fileMap := make(map[string]io.Reader)

	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to open the file"})
			return
		}
		defer file.Close()
		
		fileMap[fileHeader.Filename] = file
	}

	paths, err := uc.productUseCase.UploadProductImages(c, fileMap, uc.Env.HostAddress)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to upload images"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"images": paths,
	})
}