package service

import (
	"bytes"
	"context"
	"fmt"
	"github.com/gdan0324/ByteWeGo/api/pkg/jwt"
	"github.com/gdan0324/ByteWeGo/api/pkg/minio"
	"github.com/gdan0324/ByteWeGo/video/dal/db"
	"github.com/gdan0324/ByteWeGo/video/kitex_gen/videoservice"
	"github.com/gofrs/uuid"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"image"
	"image/jpeg"
	"os"
)

type CreateVideoService struct {
	ctx context.Context
}

// NewCreateVideoService new CreateVideoService
func NewCreateVideoService(ctx context.Context) *CreateVideoService {
	return &CreateVideoService{ctx: ctx}
}

// CreateVideo create video info.
func (s *CreateVideoService) CreateVideo(req *videoservice.CreateVideoRequest) error {
	claim, err := jwt.ParseToken(req.Token)
	userId := claim["userId"].(int64)
	if err != nil {
		return err
	}
	//userId := int64(1)
	MinioBucketName := minio.VideoBucketName
	videoData := req.Data
	reader := bytes.NewReader(videoData)
	u2, err := uuid.NewV4()
	if err != nil {
		return nil
	}
	filename := u2.String() + ".mp4"
	// 上传视频
	err = minio.UploadFile(MinioBucketName, filename, reader, int64(len(videoData)))
	if err != nil {
		return err
	}
	// 获取视频链接
	url, err := minio.GetFileUrl(MinioBucketName, filename, 0)
	if err != nil {
		return nil
	}
	u3, err := uuid.NewV4()
	if err != nil {
		return err
	}

	// 获取封面
	coverPath := u3.String() + ".jpg"
	coverData, err := readFrameAsJpeg(url.String())
	if err != nil {
		return err
	}

	// 上传封面
	coverReader := bytes.NewReader(coverData)
	err = minio.UploadFile(MinioBucketName, coverPath, coverReader, int64(len(coverData)))
	if err != nil {
		return err
	}

	// 获取封面链接
	coverUrl, err := minio.GetFileUrl(MinioBucketName, coverPath, 0)
	if err != nil {
		return err
	}

	video := &db.Video{
		UserId:   userId,
		PlayUrl:  url.String(),
		CoverUrl: coverUrl.String(),
		Title:    req.Title,
	}
	return db.CreateVideo(s.ctx, video)
}

// ReadFrameAsJpeg
// 从视频流中截取一帧并返回 需要在本地环境中安装ffmpeg并将bin添加到环境变量
func readFrameAsJpeg(filePath string) ([]byte, error) {
	reader := bytes.NewBuffer(nil)
	err := ffmpeg.Input(filePath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", 1)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(reader, os.Stdout).
		Run()
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	jpeg.Encode(buf, img, nil)

	return buf.Bytes(), err
}
