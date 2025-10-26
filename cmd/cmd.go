package cmd

import (
	"ipahg/internal"

	"github.com/spf13/cobra"
)

var fetchInfo = &cobra.Command{
	Use:   "fetch []",
	Short: "basic info of any github user",
	Long:  "This command gives some basic info of the username entered",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		user := "skarekroe666"
		if len(args) == 1 {
			user = args[0]
		}
		internal.FetchApi(user)
	},
}

func init() {
	rootCmd.AddCommand(fetchInfo)
}
