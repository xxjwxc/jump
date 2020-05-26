package run

import (
	"fmt"
	"time"

	"github.com/xxjwxc/public/mycobra"

	"github.com/xxjwxc/public/mysignal"

	"github.com/spf13/cobra"
	"github.com/xxjwxc/public/mylog"
	"github.com/xxjwxc/public/myssh"
	"github.com/xxjwxdc/rmon/internal/config"
)

type runner struct {
	file string
	port int
}

// InitFlag 初始化flag
func InitFlag(cmd *cobra.Command) {
	// conf.InitFlag(cmd)
	cmd.Flags().StringP("file", "f", "", "Monitoring file log output. 监控文件内容")
	cmd.Flags().IntP("port", "p", 0, "Monitoring port. 监控端口")
}

// InitRun init in from cmd tags
func InitRun(cmd *cobra.Command) {
	// conf.ReadConfig(cmd)
	rr := &runner{}
	rr.readRun(cmd)
	rr.run() // runing
}

func (rr *runner) readRun(cmd *cobra.Command) {
	mycobra.IfReplace(cmd, "file", &rr.file)
	mycobra.IfReplace(cmd, "port", &rr.port)
}

func (rr *runner) run() {
	sshclient, err := myssh.New(config.GetSSH().IP, config.GetSSH().Username,
		config.GetSSH().Password, config.GetSSH().Port)
	if err != nil {
		mylog.ErrorString(fmt.Sprintf("init ssh error:%v", err))
		return
	}

	w := myssh.NewWriter()
	defer w.Close()

	r := myssh.NewReader()
	defer r.Close()

	r.ListenStdin() // 监听stdin 并push 到ssh远程
	w.Run()         // 远程内容写入终端

	slg := mysignal.New()
	go func() {
		err = sshclient.Enter(w, r) // start inter the ssh
		if err != nil {
			mylog.ErrorString(fmt.Sprintf("init session error:%v", err))
		}
		slg.NotifyStop()
	}()

	for _, c := range config.GetCmd() {
		r.Push(c) // 执行一个命令
		time.Sleep(1 * time.Second)
	}
	
	if len(rr.file) > 0 {
		r.Push(fmt.Sprintf("tail -f %v", rr.file))
	}
	if rr.port > 0 {
		r.Push(fmt.Sprintf("netstat -ntpl|grep %v", rr.port))
	}

	slg.Wait() // 等待结束命令

	mylog.Info("exit.......")
}
