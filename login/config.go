package login

type HttpConfig struct {
	Scheme string
	Addr   string
}

var (
	Database HttpConfig
	Login    HttpConfig
)
