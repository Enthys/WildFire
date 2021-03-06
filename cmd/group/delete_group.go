package group

import (
	"errors"
	"github.com/kyokomi/emoji/v2"
	"github.com/spf13/cobra"
	"wildfire/pkg"
)

func NewDeleteGroupCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "delete <group name>",
		Short: "Delete project group",
		Long: `Delete a project group from provided configuration`,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("invalid number of arguments provided")
			}

			return nil
		},
		RunE: pkg.ProjectFunc(func(config *pkg.WildFireConfig, cmd *cobra.Command, args []string) (*pkg.WildFireConfig, bool, error) {
			groupService := pkg.NewGroupService(config)
			groupName := args[0]
			group := groupService.GetGroup(groupName)

			if group == nil {
				return config, false, emoji.Errorf("Group '%s' does not exist in configuration.", groupName)
			}

			groupService.DeleteGroup(groupName)

			return config, true, nil
		}),
		SilenceUsage: true,
		SilenceErrors: true,
	}
}
