/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/xxjwxdc/rmon/internal/view/conf"

	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "set default public config info .公共配置保存",
	Long:  `Configure public default information。公共配置保存`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("config called")
		conf.InitConfig(cmd) // read and save config model
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	conf.InitFlag(configCmd)

}
