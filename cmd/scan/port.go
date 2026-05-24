/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package scan

import (
	"fmt"

	"github.com/Adam-Belkadi/pogo/internal/scanner"
	"github.com/spf13/cobra"
)

var (
	urlPath string
)

// portCmd represents the port command
var portCmd = &cobra.Command{
	Use:   "port",
	Short: "This ports a remote url",
	Long:  `	`,
	Run: func(cmd *cobra.Command, args []string) {
		scanner.Scan(urlPath)
	},
}

func init() {
	portCmd.Flags().StringVarP(&urlPath, "url", "u", "", "The url to port")

	if err := portCmd.MarkFlagRequired("url"); err != nil {
		fmt.Println(err)
	}

	ScanCmd.AddCommand(portCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// portCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// portCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
