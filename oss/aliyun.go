/**
 * @author: D-S
 * @date: 2020/9/2 10:21 下午
 */

package oss

import (
	"bytes"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type Config struct {
	EndPoint        string `json:"end_point"`
	Bucket          string `json:"bucket"`
	Timeout         int32  `json:"timeout"`
	AccessKeyId     string `json:"access_key_id"`
	SecretAccessKey string `json:"secret_access_key"`
}

var aliyunOss *AliyunOss

func GetOss() *AliyunOss {
	if aliyunOss == nil {
		panic("aliyun oss 连接为空")
	}
	return aliyunOss
}

type AliyunOss struct {
	clint  *oss.Client
	bucket *oss.Bucket
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
		clint:  client,
		bucket: bucket,
	}
}

func (a *AliyunOss) Upload(key string, data []byte) (string, error) {
	r := bytes.NewReader(data)

	if err := a.bucket.PutObject(key, r); err != nil {
		return "", err
	}
	return a.getName(key), nil
}

func (a *AliyunOss) getName(key string) string {
	return key
}
