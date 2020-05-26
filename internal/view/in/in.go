package in

import (
	"fmt"

	"github.com/xxjwxc/public/mysignal"

	"github.com/spf13/cobra"
	"github.com/xxjwxc/public/mylog"
	"github.com/xxjwxc/public/myssh"
	"github.com/xxjwxdc/rmon/internal/config"
)

// InitIn init in from cmd tags
func InitIn(cmd *cobra.Command) {
	// conf.ReadConfig(cmd)
	readIn(cmd)
	run() // runing
}

// InitFlag 初始化flag
func InitFlag(cmd *cobra.Command) {
	// conf.InitFlag(cmd)
}

func readIn(cmd *cobra.Command) {
}

func run() {
	sshclient, err := myssh.New(config.GetSSH().IP, config.GetSSH().Username,
		config.GetSSH().Password, config.GetSSH().Port)
	if err != nil {
		mylog.ErrorString(fmt.Sprintf("init ssh error:%v", err))
		return
	}

	slg := mysignal.New()
	go func() {
		sshclient.EnterTerminal() // start inter the ssh
		mylog.Info("exit...")
		slg.NotifyStop()
	}()

	slg.Wait() // 等待结束命令

	mylog.Info("exit logout.......")
}
