package consts

const (
	UserTableName   = "user"
	FollowTableName = "follow"
	SecretKey       = "secret key"
	IdentityKey     = "id"
	Total           = "total"
	Notes           = "notes"
	// TODO: change to current service name
	ApiServiceName  = "apiservice"
	UserServiceName = "userservice"
	MySQLDefaultDSN = "bytewego:kidNRNKefWmMNky8@tcp(112.74.41.224:3306)/bytewego?charset=utf8&parseTime=True&loc=Local"
	TCP             = "tcp"
	// service address
	UserServiceAddr = ":9000"
	NoteServiceAddr = ":10000"
	ExportEndpoint  = ":4317"
	ETCDAddress     = "127.0.0.1:2379"
	DefaultLimit    = 10
)
