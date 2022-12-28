package cmd

import (
	"fmt"
	"github.com/xunull/helper-gitea/pkg/global"
	"github.com/xunull/helper-gitea/pkg/helper_api"

	"github.com/spf13/cobra"
)

// todo
// userCmd represents the user command
var userCmd = &cobra.Command{
	Use:   "user",
	Short: "user",
	Long:  `user`,
	PreRun: func(cmd *cobra.Command, args []string) {
		helper_api.InitDefaultApi()
	},
}

var userRepoCmd = &cobra.Command{
	Use:   "repo",
	Short: "repo",
	Long:  `repo`,
}

var userRepoCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create",
	Long:  `create`,
	Run: func(cmd *cobra.Command, args []string) {
		// todo
		fmt.Println("user create called")
	},
}

var userRepoListCmd = &cobra.Command{
	Use:   "list",
	Short: "list",
	Long:  `list`,
	PreRun: func(cmd *cobra.Command, args []string) {
		helper_api.InitDefaultApi()
	},
	Run: func(cmd *cobra.Command, args []string) {
		helper_api.Api.ListUserRepo()
	},
}

func init() {
	rootCmd.AddCommand(userCmd)
	userCmd.AddCommand(userRepoCmd)
	// todo
	userRepoCmd.AddCommand(userRepoCreateCmd)
	userRepoCmd.AddCommand(userRepoListCmd)

	userRepoListCmd.Flags().BoolVar(&global.CommonFlag.UserListRepoOnlyUser, "user", false, "only current user repos")
}
