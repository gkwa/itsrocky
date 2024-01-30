/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/taylormonacelli/itsrocky/data"
	"github.com/taylormonacelli/itsrocky/report"
)

// report1Cmd represents the report1 command
var report1Cmd = &cobra.Command{
	Use:   "report1",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := data.RunFetch()
		if err != nil {
			fmt.Println(err)
		}

		repos, err := data.LoadFromFile()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error loading from file: %v\n", err)
			os.Exit(1)
		}

		cRepos, err := data.BuildCustomizedRepositoryInfoSlice(repos)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error building customized repository info list: %v\n", err)
			os.Exit(1)
		}

		err = report.RunReport1(cRepos)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error running report: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(report1Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// report1Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// report1Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
