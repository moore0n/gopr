package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"os/exec"
	"strings"

	"github.com/urfave/cli/v2"
)

const prtemplate = "https://%s/compare/%s...%s"

func main() {
	app := cli.NewApp()
	app.Name = "gopr"
	app.Version = "1.0.0"

	app.Usage = "Open the link for creating a pull request in Github"

	app.UsageText = "gopr [options...] <target>"

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "title",
			Usage: "The title for your pull request",
		},
		&cli.StringFlag{
			Name:  "body",
			Usage: "The body for your pull request",
		},
		&cli.StringFlag{
			Name:  "remote",
			Usage: "The name of the remote to target",
			Value: "origin",
		},
	}

	app.Action = action

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func action(c *cli.Context) error {
	targetBranch := c.Args().Get(0)

	if targetBranch == "" {
		cli.ShowAppHelpAndExit(c, 0)
	}

	// Get any provided options
	title := c.String("title")
	body := c.String("body")
	remote := c.String("remote")

	// Determine the origin url from the current git remote url.
	remoteURL := getOriginURL(remote)

	// Get the name of the current branch.
	branch := getCurrentBranch()

	// Compile the base template url without query params
	targetURL := fmt.Sprintf(prtemplate, remoteURL, "master", branch)

	// Start building the query params.
	vals := url.Values{}

	if title != "" {
		vals.Add("title", title)
	}

	if body != "" {
		vals.Add("body", body)
	}

	// Encode the query params so that we can escaped data.
	encode := vals.Encode()

	// Add the query params, if there are any, to the base url
	if encode != "" {
		targetURL = fmt.Sprintf("%s?%s", targetURL, encode)
	}

	// Trigger the default browser to open the url.
	cmd := exec.Command("open", targetURL)

	if err := cmd.Run(); err != nil {
		log.Fatalf("error opening pr url: %s", err)
	}

	return nil
}

func getOriginURL(remote string) string {
	rawURL, err := exec.Command("git", "ls-remote", "--get-url", remote).Output()

	if err != nil {
		log.Fatalf("error: unable to get the remote url: %s", err)
	}

	if strings.Index(string(rawURL), "git@") == -1 {
		log.Fatalf("error: unable to get url for remote: %s\n", remote)
	}

	url := string(rawURL)
	// Strip the git
	url = strings.Replace(url, "git@", "", 1)
	// Strip the colon
	url = strings.Replace(url, ":", "/", 1)
	// Strip the suffix
	url = strings.Replace(url, ".git", "", 1)

	url = strings.Trim(url, "\n")

	return url
}

func getCurrentBranch() string {
	rawBranch, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()

	if err != nil {
		log.Fatalf("error: unable to get the current branch name: %s", err)
	}

	branch := strings.Trim(string(rawBranch), "\n")

	return branch
}
