package conf

import (
	"fmt"
	"net"

	"github.com/spf13/cobra"
	"github.com/xxjwxc/public/mycobra"
	"github.com/xxjwxdc/rmon/internal/config"
)

func init() {

}

// InitConfig init config from cmd tags
func InitConfig(cmd *cobra.Command) {
	readConfig(cmd)
	config.SaveToFile() // 保存配置
}

func readConfig(cmd *cobra.Command) {
	_info := config.GetInfoXXX()

	ip := net.ParseIP(_info.SSH.Host) // ip host init
	err := mycobra.IfReplace(cmd, "host", &ip)
	if err != nil {
		fmt.Println(err)
	}
	_info.SSH.Host = ip.String()

}
