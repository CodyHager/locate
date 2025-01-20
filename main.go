/*
Copyright © 2025 Cody Hager chager004@gmail.com
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
