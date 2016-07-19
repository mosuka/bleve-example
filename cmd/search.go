package cmd

import (
	utils "bleve-example/utils"
	"encoding/json"
	"fmt"
	"github.com/blevesearch/bleve"
	"github.com/spf13/cobra"
	"os"
)

var searchCmdIndexPath string
var searchCmdQueryStr string

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search data of index.",
	Long:  "Search data of index.",
	Run: func(cmd *cobra.Command, args []string) {
		var index bleve.Index
		var err error
		var indexMapping *bleve.IndexMapping

		_, err = os.Stat(searchCmdIndexPath)
		if err == nil {
			index, err = bleve.Open(searchCmdIndexPath)
		} else {
			indexMapping, _ = utils.CreateMapping()
			index, err = bleve.New(searchCmdIndexPath, indexMapping)
		}
		defer index.Close()

		if err != nil {
			fmt.Printf(`{"status":"NG", "message":"%s"}\n`, err.Error())
			return
		}

		query := bleve.NewQueryStringQuery(searchCmdQueryStr)
		searchRequest := bleve.NewSearchRequest(query)
		searchResult, _ := index.Search(searchRequest)

		bytesSearchResult, err := json.Marshal(searchResult)
		fmt.Println(string(bytesSearchResult))
	},
}

func init() {
	searchCmd.Flags().StringVarP(&searchCmdIndexPath, "index", "i", "", "Index directory path.")
	searchCmd.Flags().StringVarP(&searchCmdQueryStr, "query", "q", "", "Query to search index.")

	RootCmd.AddCommand(searchCmd)
}
