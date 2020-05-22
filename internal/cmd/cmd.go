package cmd

import "github.com/spf13/cobra"

var (
	rootCmd = &cobra.Command{
		Use:   "cobra",
		Short: "remote monitoring",
		Long: `rmon is a remote monitoring base on ssh , 
		support ptrace to monitor runing programs cmd info .`,
	}
)

// Execute exec init
func Execute() error {
	return rootCmd.Execute()
}

func init() {

}
