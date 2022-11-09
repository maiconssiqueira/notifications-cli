package cmd

import (
	"fmt"
	"regexp"

	"github.com/maiconssiqueira/ci-notifications/github"
	"github.com/spf13/cobra"
)

var releasesCmd = &cobra.Command{
	Use:   "releases",
	Short: "Set a new release to a Github repository",
	RunE: func(_ *cobra.Command, _ []string) error {

		valid, _ := regexp.MatchString("^(v[0-9]+)(\\.[0-9]+)(\\.[0-9])(|\\-rc)(|\\-beta)(|\\-alpha)([0-9])$", tagName)
		if !valid {
			return fmt.Errorf(`this organization uses the semantic version pattern. You sent %v and the allowed is [v0.0.0, v0.0.0-rc0, v0.0.0-beta0]`, tagName)
		}
		//TODO
		github := github.Github{}
		init := github.ReleasesInit(tagName, targetCommitish, name, body, draft, prerelease, generateReleaseNotes, *repoConf)
		res, err := github.SetRelease(init)
		if err != nil {
			return err
		}
		fmt.Println(res)
		return nil
	},
}

var tagName string
var targetCommitish string
var name string
var body string
var draft bool
var prerelease bool
var generateReleaseNotes bool

func init() {
	rootCmd.AddCommand(releasesCmd)
	releasesCmd.Flags().StringVarP(&tagName, "tagName", "t", "", `The name of the tag. Example: v1.0.2`)
	releasesCmd.Flags().StringVarP(&targetCommitish, "targetCommitish", "T", "", `Specifies the commitish value that determines where the Git tag is created from. 
	Can be any branch or commit SHA. Unused if the Git tag already exists. Default: the repository's default branch (usually master)`)
	releasesCmd.Flags().StringVarP(&name, "name", "n", "", "The name of the release")
	releasesCmd.Flags().StringVarP(&body, "body", "b", "", "Text describing the contents of the tag. Markdown style")
	releasesCmd.Flags().BoolVarP(&draft, "draft", "d", false, "True to create a draft (unpublished) release, false to create a published one")
	releasesCmd.Flags().BoolVarP(&prerelease, "prerelease", "p", false, "True to identify the release as a prerelease. false to identify the release as a full release")
	releasesCmd.Flags().BoolVarP(&generateReleaseNotes, "generateReleaseNotes", "g", true, `Whether to automatically generate the name and body for this release. 
	If name is specified, the specified name will be used; otherwise, a name will be automatically generated. 
	If body is specified, the body will be pre-pended to the automatically generated notes`)
	releasesCmd.MarkFlagRequired("tagName")
	releasesCmd.MarkFlagRequired("targetCommitish")
	releasesCmd.MarkFlagRequired("name")
}
