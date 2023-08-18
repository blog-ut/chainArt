// Package util
// Time : 2023/8/4 9:29
// Author : PTJ
// File : File
// Software: GoLand
package util

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"io"
)

func FilePush(byteData io.Reader, size int64) (string, error) {
	bucket := "micro-ut"
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	UUID, _ := uuid.NewUUID()
	key := fmt.Sprintf("%v", UUID)

	mac := qbox.NewMac("a7__nphz-KdZSFsGyBc4oDcDeve8cixj4DwcbSLS", "pUjlZChW2Oq9K7CXibYSp84yK4IjCyWmBirodGgf")
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Region = &storage.ZoneHuanan
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false

	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	// 可选配置
	putExtra := storage.PutExtra{}
	err := formUploader.Put(context.Background(), &ret, upToken, key, byteData, size, &putExtra)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("https://micro.xpit.top/%v", UUID), nil
}
