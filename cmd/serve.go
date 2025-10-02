/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"ee/tech-test-cjc/internal/github"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	port int
	url  string
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the webserver on the default port",
	Long: `Start a simple HTTP web server API. That interacts with the GitHub API and responds to requests on
/<USER> with a list of the user’s publicly available Gists
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Serving API on port: %d", port)
		client := github.New(url)
		client.Serve(port)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	serveCmd.Flags().IntVarP(&port, "port", "p", 8080, "Port to run the webserver on")
	serveCmd.Flags().StringVarP(&url, "url", "u", "https://api.github.com/users/%s/gists", "URL to fetch gists from")
}
