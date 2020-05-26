package conf

import (
	"fmt"
	"net"

	"github.com/xxjwxc/public/tools"

	"github.com/xxjwxc/public/mylog"

	"github.com/spf13/cobra"
	"github.com/xxjwxc/public/mycobra"
	"github.com/xxjwxdc/rmon/internal/config"
)

func init() {

}

// InitFlag 初始化flag
func InitFlag(cmd *cobra.Command) {
	cmd.Flags().IPP("ip", "i", nil, "host of ssh (127.0.0.1). 主机ip")
	cmd.Flags().IntP("port", "P", 22, "port of ssh (22). 主机端口")
	cmd.Flags().StringP("pwd", "p", "", "password of ssh. 主机密码")
	cmd.Flags().StringP("user", "u", "", "user of ssh. 主机用户名")
	cmd.Flags().StringP("dir", "d", "", "dir of ssh. 主机主目录")
	cmd.Flags().StringSliceP("cmd", "c", nil, "commond of shell. run 模式下进入主目录执行的命令")
	// viper.BindPFlag("host", cmd.Flags().Lookup("host"))
	// cmd.MarkFlagRequired("host")

}

// InitConfig init config from cmd tags
func InitConfig(cmd *cobra.Command) {
	ReadConfig(cmd)
	config.SaveToFile() // 保存配置
	mylog.Info(fmt.Sprintf("config done:%v", tools.JSONDecode(config.GetInfoXXX())))
}

func ReadConfig(cmd *cobra.Command) {
	_info := config.GetInfoXXX()

	ce := func(err error, msg string) {
		if err != nil {
			mylog.Fatalf("%v error: %v", msg, err)
		}
	}

	ip := net.ParseIP(_info.SSH.IP) // ip host init
	err := mycobra.IfReplace(cmd, "ip", &ip)
	ce(err, "ip")

	_info.SSH.IP = ip.String()

	if _info.SSH.Port == 0 {
		_info.SSH.Port = 22
	}
	err = mycobra.IfReplace(cmd, "port", &_info.SSH.Port)
	ce(err, "port")

	err = mycobra.IfReplace(cmd, "pwd", &_info.SSH.Password)
	ce(err, "pwd")

	err = mycobra.IfReplace(cmd, "user", &_info.SSH.Username)
	ce(err, "user")

	// err = mycobra.IfReplace(cmd, "dir", &_info.Dir)
	// ce(err, "dir")
	err = mycobra.IfReplace(cmd, "cmd", &_info.Cmd)
	ce(err, "cmd")
}
