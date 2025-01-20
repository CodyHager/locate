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

// dirCmd represents the dir command
var dirCmd = &cobra.Command{
	Use:   "dir dirName",
	Short: "Search for a directory",
	Long:  `Search for a directory in your filesystem.`,
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
		if paths, err := engine.SearchForItem(true, args[0], os.DirFS(path), strict, caseSensitive, excludes); err != nil {
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
	rootCmd.AddCommand(dirCmd)
	dirCmd.Flags().StringVarP(&path, "path", "p", "", `Root directory to start searching at, 
	defaults to your working directory`)
	dirCmd.Flags().BoolVarP(&strict, "strict", "s", false, `If set to false, will include directory names that contain the name of 
	the desired directory. Default is false.`)
	dirCmd.Flags().BoolVarP(&caseSensitive, "case-sensitive", "c", false, `If set to false, will ignore the case of directory names. 
	Default is false.`)
	dirCmd.Flags().StringArrayVarP(&excludes, "exlude", "e", nil, `Strings to exlude when searching through file names.`)
}
