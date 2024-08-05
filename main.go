/*
@Time : 2024/8/4 下午4:31
@Author : ljn
@File : main
@Software: GoLand
*/

package main

import (
	"github.com/spf13/cobra"
)

func main() {
	Execute()
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
