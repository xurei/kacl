package cmd

import (
	"github.com/fatih/color"
	"github.com/helstern/kacl/src/main/golang/changelog"
	"github.com/spf13/cobra"
	"os"
	"strings"

	"github.com/helstern/kacl/src/main/golang/prompt"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a CHANGELOG.md file",
	Long: `Initialize will create a new CHANGLE.md file in the 
http://keepachangelog.com/en/1.0.0/ format.`,
	Run: func(cmd *cobra.Command, args []string) {

		force, _ := cmd.LocalFlags().GetBool("force")

		if !force {
			if _, err := os.Stat("./CHANGELOG.md"); !os.IsNotExist(err) {
				color.Red("CHANGELOG.md already exists!")
				return
			}
		}

		var cfg changelog.InitTemplateData

		prompt.For("Project URL", &cfg.ProjectURL)
		prompt.ForWithDefault("Initial commit", "0.0.1", &cfg.InitialTag)

		cfg.ProjectURL = strings.TrimRight(cfg.ProjectURL, " /")

		f, err := os.Create("./CHANGELOG.md")
		if err != nil {
			color.Red(err.Error())
			return
		}

		err = changelog.Init(f, cfg)

		if err != nil {
			color.Red(err.Error())
			return
		}

		f.Close()
	},
}

func init() {
	RootCmd.AddCommand(initCmd)
	initCmd.Flags().BoolP("force", "f", false, "Will force the creation of the CHANGELOG.md file even if it already exists")
}
