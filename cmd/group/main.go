package group

import "github.com/spf13/cobra"

var GroupCmd = &cobra.Command{
	Use: "group",
}

func init() {
	GroupCmd.AddCommand(createGroupCmd)
}
