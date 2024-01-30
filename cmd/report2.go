package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/taylormonacelli/itsrocky/data"
	"github.com/taylormonacelli/itsrocky/report"
)

// report2Cmd represents the report2 command
var report2Cmd = &cobra.Command{
	Use:   "report2",
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

		err = report.RunReport2(cRepos)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error running RunReport2: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(report2Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// report2Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// report2Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
