package minio

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gdan0324/ByteWeGo/api/pkg/consts"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var (
	minioClient     *minio.Client
	Endpoint        = consts.MinioEndpoint
	AccessKeyId     = consts.AccessKeyId
	SecretAccessKey = consts.SecretAccessKey
	UseSSL          = consts.UseSSL
	VideoBucketName = consts.VideoBucketName
)

// Minio 对象存储初始化
func init() {
	client, err := minio.New(Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(AccessKeyId, SecretAccessKey, ""),
		Secure: UseSSL,
	})
	if err != nil {
		klog.Errorf("minio client init failed: %v", err)
	}
	//fmt.Println(client)
	klog.Debug("minio client init successfully")
	minioClient = client
	if err := CreateBucket(VideoBucketName); err != nil {
		klog.Errorf("minio client init failed: %v", err)
	}
}
