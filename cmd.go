/*
@Time : 2024/8/4 下午4:34
@Author : ljn
@File : cmd
@Software: GoLand
*/

package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "BlockChainDAPP",
	Short: "一键部署区块链毕设app🥳🔧",
	Long: fmt.Sprintf(
		"%s\n%s\n%s\n%s",
		"@Author: ljn",
		"@Description: BlockChainDAPP🥳🔧\n是一个用于快速部署基于区块链的Web应用程序的工具。\n它能够帮助开发者快速生成Dockerfile，并将应用容器化，简化了部署流程。",
		"@UpdateTime: 2024-08-04 21:09:02",
		"@Contact: <mailto:1521505611@qq.com>",
	),
}

var GenerateDockerFileCmd = &cobra.Command{
	Use:   "dockerfile",
	Short: "生成 Dockerfile",
	Long: `GenerateDockerFile 命令用于生成 Dockerfile 文件。
通过此命令，您可以快速创建一个适用于您的 Go 项目 Dockerfile 文件，方便项目的容器化和部署。`,
	Run: func(cmd *cobra.Command, args []string) {
		GenerateDockerFile()
	},
}

var BuildDockerCmd = &cobra.Command{
	Use:   "build",
	Short: "构建Docker镜像",
	Long:  "构建Docker镜像",
	Run: func(cmd *cobra.Command, args []string) {
		BuilderDockerImage()
	},
}

var BuildNginxCmd = &cobra.Command{
	Use:   "nginx",
	Short: "生成 Nginx 配置文件",
	Long:  "生成 Nginx 配置文件",
	Run: func(cmd *cobra.Command, args []string) {
		GenerateNginxConf()
	},
}

func init() {
	printLogo()
	rootCmd.AddCommand(
		GenerateDockerFileCmd,
		BuildDockerCmd,
		BuildNginxCmd,
	)
}

func printLogo() {
	const (
		Reset  = "\033[0m"
		Red    = "\033[31m"
		Green  = "\033[32m"
		Yellow = "\033[33m"
		Blue   = "\033[34m"
		Purple = "\033[35m"
		Cyan   = "\033[36m"
		White  = "\033[37m"
	)

	fmt.Println(Cyan + `
  ____  _            _        _____ _                 _ 
 |  _ \| |          | |      |  __ (_)               | |
 | |_) | | ___   ___| | _____| |__) | _ __ __ _ _ __ | |
 |  _ <| |/ _ \ / __| |/ / _ \  ___/ | '__/ _` + "`" + ` | '_ \| |
 | |_) | | (_) | (__|   <  __/ |   | | | | (_| | | | |_|
 |____/|_|\___/ \___|_|\_\___|_|   |_|_|  \__,_|_| |_(_)
` + Reset)
}
