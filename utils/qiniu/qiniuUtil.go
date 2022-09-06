package qiniu

import (
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"webapp01/utils/config"
)

func GetQiniuToken() (qiniuToken string) {
	putPolicy := storage.PutPolicy{Scope: config.Qiniu.Bucket}
	mac := qbox.NewMac(config.Qiniu.AccessKey, config.Qiniu.SecretKey)
	qiniuToken = putPolicy.UploadToken(mac)
	return qiniuToken
}
