package util

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func CreateConfig() {
	var username string
	fmt.Println("Enter username: ")
	fmt.Scanf("%s", &username)
	viper.Set("Username", username)
	err := viper.SafeWriteConfig()
	if err != nil {
		fmt.Printf("Error when writing to config file: %v", err)
		os.Exit(1)
	}
}

func ReadInConfig() error {
	if err := viper.ReadInConfig(); err != nil {
		return err
	} else {
		return nil
	}
}

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Error when creating config file: %v", err)
		os.Exit(1)
	}
	viper.AddConfigPath(home)
	viper.SetConfigType("json")
	viper.SetConfigName(".locate")
}
