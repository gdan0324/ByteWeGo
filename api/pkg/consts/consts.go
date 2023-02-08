package consts

const (
	UserTableName = "user"
	SecretKey     = "secret key"
	IdentityKey   = "id"
	Total         = "total"
	Notes         = "notes"
	// TODO: change to current service name
	ApiServiceName  = "demoapi"
	NoteServiceName = "demonote"
	UserServiceName = "demouser"
	MySQLDefaultDSN = "root:n7Zs3usIM15HlkvQ@tcp(120.46.190.10:3306)/bytewego?charset=utf8&parseTime=True&loc=Local"
	TCP             = "tcp"
	// service address
	UserServiceAddr = ":9000"
	NoteServiceAddr = ":10000"
	ExportEndpoint  = ":4317"
	ETCDAddress     = "127.0.0.1:2379"
	DefaultLimit    = 10
)
