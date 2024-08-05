/*
@Time : 2024/8/4 下午9:12
@Author : ljn
@File : buildDocker
@Software: GoLand
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func BuilderDockerImage() {
	const usageInfo = `构建之前你需要了解你deploy目录结构:
ljn@ljndembp deploy % tree
.
├── back
│   ├── back # 可执行程序
│   └── config.ini # 配置文件
├── front
│   └── dist # vue编译目录
├── nginx.conf # nginx代理文件
└── sql
    └── database.sql # 与数据库名称相同的sql文件
`

	fmt.Printf("%s\n", usageInfo)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter docker image name:👇")
	imageName, _ := reader.ReadString('\n')
	imageName = strings.TrimSpace(imageName)
	if imageName == "" {
		fmt.Println("imageName cannot be empty")
		os.Exit(1)
	}
	err := ExecuteCommand("docker", "build", "-t", imageName, ".")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Println("Building docker image success 🥳")
	fmt.Println("start container example: 👇")
	fmt.Println("docker run -it --name=container_name -p 8080:8080 -p 8081:8081  -d " + imageName)
}
