package in

import (
	"fmt"

	"github.com/xxjwxc/public/mysignal"

	"github.com/spf13/cobra"
	"github.com/xxjwxc/public/mylog"
	"github.com/xxjwxc/public/myssh"
	"github.com/xxjwxdc/rmon/internal/config"
	"github.com/xxjwxdc/rmon/internal/view/conf"
)

// InitIn init in from cmd tags
func InitIn(cmd *cobra.Command) {
	conf.ReadConfig(cmd)
	readIn(cmd)
	run() // runing
}

// InitFlag 初始化flag
func InitFlag(cmd *cobra.Command) {
	conf.InitFlag(cmd)
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

	w := myssh.NewWriter()
	r := myssh.NewReader()

	dir := config.GetDir()
	if dir != "." {
		go r.Push("cd " + dir) // 执行一个命令
	}

	go func() {
		for {
			b := make([]byte, 1024)
			o, err := r.Read(b)
			fmt.Println(o, err)
		}
	}()
	r.ListenStdin() // 监听stdin 并push 到ssh远程
	w.Run()         // 远程内容写入终端

	slg := mysignal.New()
	go func() {
		err = sshclient.Enter(w, r) // start inter the ssh
		if err != nil {
			mylog.ErrorString(fmt.Sprintf("init session error:%v", err))
		}
		mylog.Info("logout...")
		slg.NotifyStop()
	}()

	slg.Wait() // 等待结束命令
	mylog.Info("logout.......")
}
