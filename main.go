package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// This function shows an example of how you would use this package
var rootCmd = &cobra.Command{
	Use:   "edit-cli",
	Short: "A test CLI to showcase the editor package",
	Run: func(cmd *cobra.Command, args []string) {
		contents, err := os.ReadFile(filePathToEdit)
		if err != nil {
			fmt.Println("File reading error: ", err)
			os.Exit(1)
		}

		// Run the editor to get the edited contents
		edited, _, err := Run(contents, filePathToEdit)
		if err != nil {
			fmt.Println("File editing error: ", err)
			os.Exit(1)
		}

		// TODO(briancain): This is where you can run your parsers and validators
		// on the edited text so you can show any parse or lint errors prior to
		// saving the file

		// If changes, overwrite the existing file
		// Open the file for writing, creating it if it doesn't exist
		file, err := os.OpenFile(filePathToEdit, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			fmt.Println("Failed to open original file to update: ", err)
			os.Exit(1)
		}
		defer file.Close()

		// Write the content to the file
		estr := string(edited)
		_, err = file.WriteString(estr)
		if err != nil {
			fmt.Println("Failed to write content to file: ", err)
			os.Exit(1)
		}

		fmt.Println(fmt.Sprintf("Successfully updated %q!", filePathToEdit))
		os.Exit(0)
	},
}

var filePathToEdit string

func init() {
	rootCmd.PersistentFlags().StringVar(&filePathToEdit, "file", "", "the file to edit")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
