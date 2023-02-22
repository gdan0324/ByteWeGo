package consts

const (
	UserTableName   = "user"
	FollowTableName = "follow"
	SecretKey       = "secret key"
	IdentityKey     = "id"
	Total           = "total"
	// TODO: change to current service name
	ApiServiceName       = "apiservice"
	UserServiceName      = "userservice"
	CommunityServiceName = "communityservice"
	CommentServiceName   = "commentservice"
	MySQLDefaultDSN      = "root:Ab123456@tcp(47.115.210.15:3306)/bytewego?charset=utf8&parseTime=True&loc=Local"
	TCP                  = "tcp"
	// service address
	UserServiceAddr      = ":9000"
	CommunityServiceAddr = ":9001"
	CommentServiceAddr   = ":9002"
	ExportEndpoint       = ":4317"
	ETCDAddress          = "127.0.0.1:2379"
	DefaultLimit         = 10

	// minio
	MinioEndpoint   = "localhost:8999"
	AccessKeyId     = "minioadmin"
	SecretAccessKey = "minioadmin"
	UseSSL          = false
	VideoBucketName = "video"
)
