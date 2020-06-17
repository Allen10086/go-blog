package oss

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	uuid "github.com/satori/go.uuid"
	"io"
	"os"
)

const (
	Endpoint        = "http://oss-cn-beijing.aliyuncs.com"
	AccessKeyId     = "LTAI4FzqQZARLTdg7yx7Hbtc"
	AccessKeySecret = "SivI4wEna2oAmStSHerQqesc4uKiMa"
	BuckeyName      = "lz12code"
	ObjectName      = "myblog/"
	ImgPath         = "https://lz12code.oss-cn-beijing.aliyuncs.com/"
)

func handleError(err error) {
	fmt.Println("Error:", err)
	os.Exit(-1)
}

// 图片上传
func UploadImg(reader io.Reader, suffix string) (imgUrl string) {
	// 拼接上传到oss时图片的url
	uid := uuid.NewV4()
	dirName := fmt.Sprintf("%s%s.%s", ObjectName, uid, suffix)

	imgUrl = fmt.Sprintf("%s%s", ImgPath, dirName)
	fmt.Println(imgUrl)

	// Endpoint以杭州为例，其它Region请按实际情况填写。
	endpoint := Endpoint
	// 阿里云主账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM账号进行API访问或日常运维，请登录 https://ram.console.aliyun.com 创建RAM账号。
	accessKeyId := AccessKeyId
	accessKeySecret := AccessKeySecret
	bucketName := BuckeyName
	// <yourObjectName>上传文件到OSS时需要指定包含文件后缀在内的完整路径，例如abc/efg/123.jpg。
	objectName := dirName

	// 创建OSSClient实例。
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		handleError(err)
	}
	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		handleError(err)
	}

	// 上传文件流。 reader -> 前端传过来的二进制的文件流
	err = bucket.PutObject(objectName, reader)
	if err != nil {
		handleError(err)
	}
	return imgUrl
}
