/*
@Time : 2024/8/4 下午4:32
@Author : ljn
@File : generate
@Software: GoLand
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GenerateDockerFile() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter ports to expose (comma separated and default open ports is 8080,8081):👇")
	portsInput, _ := reader.ReadString('\n')
	portsInput = strings.TrimSpace(portsInput)
	ports := strings.Split(portsInput, ",")
	if portsInput == "" {
		ports = []string{"8080", "8081"}
	}
	var exposedPorts string
	for _, port := range ports {
		exposedPorts += fmt.Sprintf("EXPOSE %s\n", strings.TrimSpace(port))
	}

	fmt.Println("Enter MySQL database name:👇")
	databaseName, _ := reader.ReadString('\n')
	databaseName = strings.TrimSpace(databaseName)
	if databaseName == "" {
		fmt.Println("Database name cannot be empty")
		os.Exit(1)
	}

	fmt.Println("Enter start command:👇")
	startCmd, _ := reader.ReadString('\n')
	startCmd = strings.TrimSpace(startCmd)
	if startCmd == "" {
		fmt.Println("Start command cannot be empty")
		os.Exit(1)
	}
	dockerFileContent := strings.Replace(DockerFileTemplate, "{{EXPOSE_PORTS}}", exposedPorts, 1)
	dockerFileContent = strings.ReplaceAll(dockerFileContent, "{{DB_NAME}}", databaseName)
	dockerFileContent = strings.Replace(dockerFileContent, "{{START_CMD}}", startCmd, 1)
	if strings.Contains(startCmd, "java") {
		dockerFileContent = strings.Replace(dockerFileContent, "#    {{INSTALL_JAVA}}",
			"    apt install -y default-jdk ", 1)
	}
	WriteToFile("./Dockerfile", dockerFileContent)
}

func GenerateNginxConf() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter nginx proxy url:👇")
	proxyURL, _ := reader.ReadString('\n')
	proxyURL = strings.TrimSpace(proxyURL)
	if proxyURL == "" {
		proxyURL = "http://localhost:8080/api"
	}
	nginxFileContent := strings.Replace(NginxTemplate, "{{PROXY_URL}}", proxyURL, 1)
	WriteToFile("./nginx.conf", nginxFileContent)
	fmt.Println("Copy the build to the specified directory (deploy) 🥳")
}
