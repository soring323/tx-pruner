package cli

import (
	app "github.com/soring323/tx-pruner/package"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start",
	Long:  `This command starts prune transactions`,
	Run: func(cmd *cobra.Command, args []string) {
		app.PruneTransactions()
	},
}

func Execute() {
	startCmd.Execute()
}
