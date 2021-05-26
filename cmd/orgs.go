package cmd

import (
	"github.com/spf13/cobra"
	"github.com/xunull/helper-gitea/pkg/helper_api"
)

var orgCmd = &cobra.Command{
	Use:   "org",
	Short: "org",
	Long:  `org`,
	PreRun: func(cmd *cobra.Command, args []string) {
		helper_api.InitDefaultApi()
	},
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var orgReposCmd = &cobra.Command{
	Use:   "repo",
	Short: "repo",
	Long:  `repo`,
	PreRun: func(cmd *cobra.Command, args []string) {
		helper_api.InitDefaultApi()
	},
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var orgReposCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create",
	Long:  `create`,
	Args:  cobra.ExactArgs(2),
	PreRun: func(cmd *cobra.Command, args []string) {
		helper_api.InitDefaultApi()
	},
	Run: func(cmd *cobra.Command, args []string) {
		org := args[0]
		repo := args[1]
		helper_api.Api.CreateOrgRepo(org, repo)
	},
}

var orgReposListCmd = &cobra.Command{
	Use:   "list",
	Short: "list",
	Long:  `list`,
	Args:  cobra.ExactArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) {
		helper_api.InitDefaultApi()
	},
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		helper_api.Api.ListOrgRepos(name)
	},
}

func init() {
	rootCmd.AddCommand(orgCmd)
	orgCmd.AddCommand(orgReposCmd)
	orgReposCmd.AddCommand(orgReposListCmd)
	orgReposCmd.AddCommand(orgReposCreateCmd)
}
