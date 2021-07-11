package cmd

import (
	"github.com/jptalukdar/DocsSphere/config"
	"github.com/jptalukdar/DocsSphere/repo"
	"github.com/spf13/cobra"
)

var repoCmd = &cobra.Command{
	Use:   "test",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) {
	// 	c := config.ReadConfig()
	// 	repo.GetRepository(c.RepoURL, c.Path)
	// },
}

func init() {

	configList := config.ReadConfig()
	for _, c := range configList {
		baseMethod := func(cmd *cobra.Command, args []string) {
			repo.GetRepository(c.RepoURL, c.Path)
		}
		var cmd = &cobra.Command{}
		cmd.Run = baseMethod
		cmd.Use = c.Command
		rootCmd.AddCommand(cmd)
	}

	// rootCmd.AddCommand(repoCmd)
}
