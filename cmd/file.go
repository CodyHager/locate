/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"locate/engine"
	"os"

	"github.com/spf13/cobra"
)

var path string

// fileCmd represents the file command
var fileCmd = &cobra.Command{
	Use:   "file",
	Short: "Search for a file",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
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
		if paths, err := engine.SearchForFile(args[0], os.DirFS(path)); err != nil {
			fmt.Printf("Error when searching for file: %v\n", err)
		} else {
			for _, path := range paths{
				fmt.Println(path)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(fileCmd)
	fileCmd.Flags().StringVarP(&path, "path", "p", "", `Root directory to start searching at, 
	defaults to your working directory`)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
