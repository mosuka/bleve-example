package utils

import (
	"encoding/json"
	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/analysis/analyzers/keyword_analyzer"
	"github.com/blevesearch/blevex/lang/ja"
)

func CreateMapping() (*bleve.IndexMapping, error) {
	japaneseTextFieldMapping := bleve.NewTextFieldMapping()
	japaneseTextFieldMapping.Analyzer = ja.AnalyzerName

	keywordFieldMapping := bleve.NewTextFieldMapping()
	keywordFieldMapping.Analyzer = keyword_analyzer.Name

	documentMapping := bleve.NewDocumentMapping()
	documentMapping.AddFieldMappingsAt("title", japaneseTextFieldMapping)
	documentMapping.AddFieldMappingsAt("description", japaneseTextFieldMapping)

	indexMapping := bleve.NewIndexMapping()
	indexMapping.AddDocumentMapping("_default", documentMapping)
	indexMapping.DefaultType = "_default"
	indexMapping.TypeField = "_type"
	indexMapping.DefaultAnalyzer = "standard"

	return indexMapping, nil
}

func CreateDocument(dataStr *string) (interface{}, error) {
	var docs interface{}

	bytesDocuments := []byte(*dataStr)
	err := json.Unmarshal(bytesDocuments, &docs)

	return docs, err
}
