package config

// Config custom config struct
type Config struct {
	CfgBase `yaml:"base"`
	SSH     SSH    `yaml:"ssh"`
	Dir     string `yaml:"dir"`
}

// SSH 远程地址端口
type SSH struct {
	IP       string
	Port     int
	Username string `validate:"required"` // Username 用户名
	Password string // Password 密码
}

// GetSSH 获取ssh信息
func GetSSH() SSH {
	return _info.SSH
}
