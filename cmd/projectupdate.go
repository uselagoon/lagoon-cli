package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/amazeeio/lagoon-cli/graphql"
	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

var projectUpdateCmd = &cobra.Command{
	Use:   "update [project name]",
	Short: "Update a project",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Not enough arguments. Requires: project name")
			cmd.Help()
			os.Exit(1)
		}
		//projectName := args[0]
		var patchData []string
		if projectGitURL != "" {
			patchData = append(patchData, `gitUrl:"`+projectGitURL+`"`)
		}
		if projectBranches != "" {
			patchData = append(patchData, `branches:"`+projectBranches+`"`)
		}
		if projectProductionEnvironment != "" {
			patchData = append(patchData, `productionEnvironment:"`+projectProductionEnvironment+`"`)
		}
		if projectPullRequests != "" {
			patchData = append(patchData, `pullRequests:"`+projectPullRequests+`"`)
		}
		if projectAutoIdle == 1 || projectAutoIdle == 0 {
			patchData = append(patchData, `autoIdle:`+strconv.Itoa(projectAutoIdle))
		}
		var responseData UpdateProject
		err := graphql.GraphQLRequest(fmt.Sprintf(`mutation {
  updateProject(input:{
    id:23
    patch:{
      %s
    }
  }){
    id
    name
  }
}`, strings.Join(patchData[:], ",")), &responseData)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(fmt.Sprintf("Result: %s\n", aurora.Green("success")))
			fmt.Println(fmt.Sprintf("%s: %s", aurora.Yellow("Project"), responseData.UpdateProject.Name))
		}
	},
}

func init() {
	projectCmd.AddCommand(projectUpdateCmd)
	projectUpdateCmd.Flags().StringVarP(&projectGitURL, "giturl", "g", "", "GitURL of the project")
	projectUpdateCmd.Flags().StringVarP(&projectBranches, "branches", "b", "", "Branches of the project")
	projectUpdateCmd.Flags().StringVarP(&projectProductionEnvironment, "prod-env", "P", "", "Production environment of the project")
	projectUpdateCmd.Flags().StringVarP(&projectPullRequests, "pull-requests", "r", "", "Pull requests of the project")
	projectUpdateCmd.Flags().IntVarP(&projectAutoIdle, "auto-idle", "a", 1, "Auto idle setting of the project")
}
