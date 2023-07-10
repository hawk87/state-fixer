/*
Copyright © 2023 hawk87
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "state-fixer",
	Short: "A state fixer for Edit State",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		upperLimit, err := cmd.Flags().GetInt("upper-limit")
		if err != nil {
			return
		}

		obj := make(map[string]int)

		for i := 1; i <= upperLimit; i++ {
			key := fmt.Sprintf("i%d", i)
			value := 1
			obj[key] = value
		}
		jsonObj, err := json.Marshal(obj)
		if err != nil {
			fmt.Printf("error in marshal")
			return
		}
		fmt.Println(string(jsonObj))
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

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.state-fixer.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().IntP("upper-limit", "l", 0, "Set upper-limit for collections")
}
