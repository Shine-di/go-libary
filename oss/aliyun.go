/**
 * @author: D-S
 * @date: 2020/9/2 10:21 下午
 */

package oss

import (
	"bytes"
	"fmt"
	"github.com/Shine-di/go-libary/log"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/h2non/filetype"
)

type Config struct {
	EndPoint        string `json:"end_point"`
	Bucket          string `json:"bucket"`
	Timeout         int32  `json:"timeout"`
	AccessKeyId     string `json:"access_key_id"`
	SecretAccessKey string `json:"secret_access_key"`
	BaseUrl         string `json:"base_url"`
}

var aliyunOss *AliyunOss

func GetOss() *AliyunOss {
	if aliyunOss == nil {
		panic("aliyun oss 连接为空")
	}
	return aliyunOss
}

type AliyunOss struct {
	clint      *oss.Client
	bucket     *oss.Bucket
	BaseUrl    string
	bucketName string
}

func InitAliyunOss(config *Config) {
	client, err := oss.New(config.EndPoint, config.AccessKeyId, config.SecretAccessKey)
	if err != nil {
		panic("oss error" + err.Error())
	}
	bucket, err := client.Bucket(config.Bucket)
	if err != nil {
		panic("oss bucket error" + err.Error())
	}

	aliyunOss = &AliyunOss{
		clint:      client,
		bucket:     bucket,
		BaseUrl:    config.BaseUrl,
		bucketName: config.Bucket,
	}
	log.Info("load aliyun oss success" + config.EndPoint)
}

func (a *AliyunOss) Upload(key string, data []byte) (string, error) {
	r := bytes.NewReader(data)
	filename := a.getName(key, data)
	if err := a.bucket.PutObject(filename, r); err != nil {
		return "", err
	}
	return a.GetUrl(filename), nil
}

func (a *AliyunOss) GetUrl(filename string) string {
	return fmt.Sprintf("http://%v.%v/%v", a.bucketName, a.BaseUrl, filename)
}

func (a *AliyunOss) getName(key string, data []byte) string {
	ext := ""
	mime := ""
	var fileType, err = filetype.Match(data)
	if err == nil {
		mime = fileType.MIME.Value
		ext = fileType.Extension
	}
	if ext == "unknown" {
		ext = ""
	}
	_ = mime
	//svg格式处理
	if fmt.Sprintf("%x", data[:8]) == "3c3f786d6c207665" {
		ext = "svg"
	}
	var filename string
	if len(ext) > 0 {
		filename = fmt.Sprintf("%v", key) + "." + ext
	} else {
		filename = fmt.Sprintf("%v", key)
	}
	return filename
}
