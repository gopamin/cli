package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gopamin",
	Short: "Gopamin CLI Tool",
	Long: `
	   ______                            _
	  / ____/___  ____  ____ _____ ___  (_)___
	 / / __/ __ \/ __ \/ __  / __  __ \/ / __ \
	/ /_/ / /_/ / /_/ / /_/ / / / / / / / / / /
	\____/\____/ .___/\__,_/_/ /_/ /_/_/_/ /_/
	          /_/

The CLI tool to scaffold Go projects`,
}

func init() {
	// mark completion hidden
	completion := completionCommand()
	completion.Hidden = true
	rootCmd.AddCommand(completion)
}

func completionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "completion",
		Short: "Generate the autocompletion script for the specified shell",
	}
}

func Execute() {
	isValid, newVersion := versionChecker()
	if !isValid {
		fmt.Printf("The newest version of the Gopamin CLI is %v but the installed version on your system is %v. To get the latest features and likely bugfixes, please first install the latest version by running 'go install github.com/gopamin/gopamin@%v'", newVersion, VERSION, newVersion)
		return
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
