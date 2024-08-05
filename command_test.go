/*
@Time : 2024/8/4 下午9:15
@Author : ljn
@File : command_test
@Software: GoLand
*/

package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestCommand(t *testing.T) {
	if strings.Contains("java -jar back.jar", "java") {
		fmt.Println(ExecuteCommand("docker", "logs", "-f", "mysql"))

	}
}
