/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate ASCII art from predefined images",
	Long:  `This command allows you to select and display predefined ASCII art images.

			Usage:
			ascii-tool generate

			Interactive Selection:
			- Provides a list of predefined ASCII art images.
			- Use arrow keys to cycle through the options.
			- Press Enter to select an image and display it.`,
	Run: func(cmd *cobra.Command, args []string) {
		displayAsciiArtOptions()
	},
}

func displayAsciiArtOptions() {
	var selectedOption string
	prompt := &survey.Select{
		Message: "Choose an ASCII art image:",
		Options: asciiArtTitles,
	}
	survey.AskOne(prompt, &selectedOption)

	for i, title := range asciiArtTitles {
		if title == selectedOption {
			fmt.Println(asciiArtOptions[i])
			break
		}
	}
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
