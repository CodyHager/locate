/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"locate/cmd"
	"locate/util"
)

func main() {
	if err := util.ReadInConfig(); err != nil {
		util.CreateConfig()
	}
	cmd.Execute()
}
