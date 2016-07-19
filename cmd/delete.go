package cmd

import (
	utils "bleve-example/utils"
	"fmt"
	"github.com/blevesearch/bleve"
	"github.com/spf13/cobra"
	"os"
)

var deleteCmdIndexPath string
var deleteCmdDataStr string

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete data from index.",
	Long:  "Delete data from index.",
	Run: func(cmd *cobra.Command, args []string) {
		var index bleve.Index
		var err error
		var indexMapping *bleve.IndexMapping

		_, err = os.Stat(deleteCmdIndexPath)
		if err == nil {
			index, err = bleve.Open(deleteCmdIndexPath)
		} else {
			indexMapping, _ = utils.CreateMapping()
			index, err = bleve.New(deleteCmdIndexPath, indexMapping)
		}
		defer index.Close()

		if err != nil {
			fmt.Printf(`{"status":"NG", "message":"%s"}\n`, err.Error())
			return
		}

		index.Delete(deleteCmdDataStr)
		fmt.Println(`{"status":"OK"}`)
	},
}

func init() {
	deleteCmd.Flags().StringVarP(&deleteCmdIndexPath, "index", "i", "", "Index directory path.")
	deleteCmd.Flags().StringVarP(&deleteCmdDataStr, "data", "d", "", "Document ID to be deleted.")

	RootCmd.AddCommand(deleteCmd)
}
