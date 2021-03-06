package group

import "github.com/spf13/cobra"

var GroupCmd = &cobra.Command{
	Use: "group",
}

func init() {
	GroupCmd.AddCommand(NewListGroupsCommand())
	GroupCmd.AddCommand(NewCreateGroupCmd())
	GroupCmd.AddCommand(NewDeleteGroupCmd())
	GroupCmd.AddCommand(NewAddProjectToGroupCmd())
	GroupCmd.AddCommand(NewRemoveProjectFromGroupCmd())
}
