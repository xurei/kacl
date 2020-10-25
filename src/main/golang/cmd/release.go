package cmd

import (
	"github.com/helstern/kacl/src/main/golang/release"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// releaseCmd represents the release command
var releaseCmd = &cobra.Command{
	Use:   "release [tag]",
	Short: "Create a new release",
	Long: `Create a new release by moving the current Unreleased changes into 
a new change with the given tag.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			color.Red("no tag specified!")
			return
		}

		tag := args[0]

		contents, ok := getContents()
		if !ok {
			return
		}

		newContents, err := release.Create(contents, tag, time.Now())
		if err != nil {
			return
		}

		writeContents(newContents)
	},
}

func init() {
	RootCmd.AddCommand(releaseCmd)
}
