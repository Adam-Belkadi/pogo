/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/Adam-Belkadi/pogo/cmd/scan"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pogo",
	Short: "Blazingly fast port scanner written in Go",
	Long: `My first serious (like really serious) open source project`,
	// Run: func(cmd *cobra.Command, args []string) {
    //     if len(args) == 0 {
    //         log.Fatal("please provide a pet")
    //     }
    //     switch args[0] {
	// 	case "dog":
    //         fmt.Println("wooooof!")
    //     case "cat":
    //         fmt.Println("meoooow!")
    //     default:
    //         fmt.Println("sorry your pet is currently not supported :(")
    //     }
    // },
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func addSubcommandPalettes() {
	rootCmd.AddCommand(scan.ScanCmd)
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.deeper.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	addSubcommandPalettes()
}


