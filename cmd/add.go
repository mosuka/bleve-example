package cmd

import (
	utils "bleve-example/utils"
	"fmt"
	"github.com/blevesearch/bleve"
	"github.com/spf13/cobra"
	"os"
)

var addCmdIndexPath string
var addCmdDataStr string

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add data to index.",
	Long:  "Add data to index.",
	Run: func(cmd *cobra.Command, args []string) {
		var index bleve.Index
		var err error
		var indexMapping *bleve.IndexMapping

		_, err = os.Stat(addCmdIndexPath)
		if err == nil {
			index, err = bleve.Open(addCmdIndexPath)
		} else {
			indexMapping, _ = utils.CreateMapping()
			index, err = bleve.New(addCmdIndexPath, indexMapping)
		}
		defer index.Close()

		if err != nil {
			fmt.Printf(`{"status":"NG", "message":"%s"}\n`, err.Error())
			return
		}

		var docs interface{}
		docs, err = utils.CreateDocument(&addCmdDataStr)
		if err != nil {
			fmt.Printf(`{"status":"NG", "message":"%s"}\n`, err.Error())
			return
		}
		for ID, doc := range docs.(map[string]interface{}) {
			index.Index(ID, doc)
		}

		fmt.Println(`{"status":"OK"}`)
	},
}

func init() {
	addCmd.Flags().StringVarP(&addCmdIndexPath, "index", "i", "", "Index directory path.")
	addCmd.Flags().StringVarP(&addCmdDataStr, "data", "d", "", "Document data formatted using JSON.")

	RootCmd.AddCommand(addCmd)
}
