/*
Copyright © 2025 Cody Hager chager004@gmail.com
*/
package cmd

import (
	"fmt"
	"locate/engine"
	"os"

	"github.com/spf13/cobra"
)

var path string
var strict bool = false
var caseSensitive bool = false
var excludes []string

// fileCmd represents the file command
var fileCmd = &cobra.Command{
	Use:   "file fileName",
	Short: "Search for a file",
	Long:  `Search for a file in your filesystem.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			cmd.Help()
			os.Exit(1)
		}
		if len(path) == 0 {
			wd, err := os.Getwd()
			if err != nil {
				fmt.Printf("Error getting working directory: %v", err)
			}
			path = wd
		}
		if paths, err := engine.SearchForItem(false, args[0], os.DirFS(path), strict, caseSensitive, excludes); err != nil {
			fmt.Printf("Error when searching for file: %v\n", err)
		} else {
			fmt.Println()
			for _, path := range paths {
				fmt.Println(path)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(fileCmd)
	fileCmd.Flags().StringVarP(&path, "path", "p", "", `Root directory to start searching at, 
	defaults to your working directory`)
	fileCmd.Flags().BoolVarP(&strict, "strict", "s", false, `If set to false, will include file names that contain the name of 
	the desired file. Default is false.`)
	fileCmd.Flags().BoolVarP(&caseSensitive, "case-sensitive", "c", false, `If set to false, will ignore the case of file names. 
	Default is false.`)
	fileCmd.Flags().StringArrayVarP(&excludes, "exlude", "e", nil, `Strings to exlude when searching through file names.`)
}
