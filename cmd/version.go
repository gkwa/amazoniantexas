package cmd

import (
	"fmt"
	
	"github.com/gkwa/amazoniantexas/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of amazoniantexas",
	Long:  `All software has versions. This is amazoniantexas's`,
	Run: func(cmd *cobra.Command, args []string) {
		buildInfo := version.GetBuildInfo()
		fmt.Println(buildInfo)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
