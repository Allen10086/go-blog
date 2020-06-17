package routers

import (
	"blog/controller"
	"github.com/gin-gonic/gin"
)

func ImgUpload(e *gin.Engine) {
	UploadControl := &controller.UploadController{}
	// 上传图片
	Upload := e.Group("/api/v1")
	{
		Upload.POST("/upload", UploadControl.UploadImg)
	}
}
