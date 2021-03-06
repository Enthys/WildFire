package group

import (
	"errors"
	"github.com/kyokomi/emoji/v2"
	"github.com/spf13/cobra"
	"wildfire/pkg"
)

func NewCreateGroupCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "create <name> [project...]",
		Short: "Create project group",
		Long: `Create new project group.

ProjectConfig groups are groups which contain projects existing in the configuration.
Groups are primarily used when commands are to be executed on specific projects.
`,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("invalid number of arguments provided")
			}

			return nil
		},
		RunE: pkg.ProjectFunc(func(config *pkg.WildFireConfig, cmd *cobra.Command, args []string) (*pkg.WildFireConfig, bool, error) {
			projectService := pkg.NewProjectService(config)
			groupService := pkg.NewGroupService(config)

			groupName := args[0]

			group, err := groupService.CreateGroup(groupName)
			if err != nil {
				return config, false, err
			}

			projectNames := args[1:]

			emoji.Printf(":star: Created new group '%s'\n", groupName)
			if len(projectNames) == 0 {
				return config, true, nil
			}

			success := true
			for _, name := range projectNames {
				exists := projectService.HasProject(name)

				if exists == false {
					emoji.Println(":prohibited: ProjectConfig ", name, " does not exist in this configuration.")
					success = false
					continue
				}

				if success == true {
					emoji.Printf(":ocean: ProjectConfig '%s' has been added to group '%s'\n", name, groupName)
					group, _ = groupService.AddProject(group, name)
				}
			}

			if success == false {
				emoji.Println(":error: Reverting configuration. Resolve issues and try again.")
				return config, false, nil
			}

			return config, true, nil
		}),
		SilenceUsage: true,
		SilenceErrors: true,
	}
}
