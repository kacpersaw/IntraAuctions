package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server [port]",
	Short: "Start wsb_projekt server",
	Long:  "",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		//model.DB = model.InitDB()
		//
		//logrus.Info("Starting API on port :" + args[0])
		//r := router.NewRouter()
		//logrus.Fatal(http.ListenAndServe(
		//	":"+args[0],
		//	router.CommonMiddleware(
		//		router.CorsMiddleware(r),
		//	),
		//))
	},
}
