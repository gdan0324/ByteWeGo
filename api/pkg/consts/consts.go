package consts

const (
	UserTableName   = "user"
	FollowTableName = "follow"
	VideoTableName  = "video"

	// TODO: change to current service name
	ApiServiceName       = "apiservice"
	UserServiceName      = "userservice"
	CommunityServiceName = "communityservice"
	CommentServiceName   = "commentservice"
	VideoServiceName     = "videoservice"
	FavoriteServiceName  = "favoriteservice"
	MySQLDefaultDSN      = "root:Ab123456@tcp(47.115.210.15:3306)/bytewego?charset=utf8&parseTime=True&loc=Local"
	TCP                  = "tcp"
	// service address
	UserServiceAddr      = ":9000"
	CommunityServiceAddr = ":9001"
	CommentServiceAddr   = ":9002"
	VideoServiceAddr     = ":9003"
	FavoriteServiceAddr  = ":9004"
	ExportEndpoint       = ":4317"
	ETCDAddress          = "127.0.0.1:2379"

	// minio
	MinioEndpoint   = "localhost:8999"
	AccessKeyId     = "minioadmin"
	SecretAccessKey = "minioadmin"
	UseSSL          = false
	VideoBucketName = "video"
)
