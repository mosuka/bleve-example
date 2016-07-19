package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "bleve-example",
	Short: "Bleve example.",
	Long:  `This is a example for using Bleve.`,
	Run:   func(cmd *cobra.Command, args []string) {},
}
