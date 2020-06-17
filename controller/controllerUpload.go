package controller

import (
	"blog/oss"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type UploadController struct {
}
// 分割文件名 获取后缀名
func splitString(fileName string) (suffix string) {
	// 取得文件后缀名
	i := strings.LastIndex(fileName, ".")
	if i > 0 {
		suffix = fileName[i+1:]
	}
	return
}
// 获取前端传过来的图片对象
func (u *UploadController) UploadImg(c *gin.Context) {
	// https://my.oschina.net/solate/blog/741039
	// 因为上传文件的类型是multipart/form-data 所以不能使用 r.ParseForm(), 这个只能获得普通post
	file, header, err := c.Request.FormFile("upload")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    1001,
			"message": err,
		})
		return
	}
	// 获取前端传过来文件的后缀名
	suffix := splitString(header.Filename)
	//fmt.Println(header.Filename, header.Size, err)
	// 上传图片到aliyun  oss
	imgUrl := oss.UploadImg(file, suffix)
	c.JSON(http.StatusOK, gin.H{
		"code":    1000,
		"message": "success",
		"imgurl":  imgUrl,
	})
	return
}
