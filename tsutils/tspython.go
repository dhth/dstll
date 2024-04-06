package tsutils

import (
	"context"
	"fmt"

	ts "github.com/smacker/go-tree-sitter"
	tspy "github.com/smacker/go-tree-sitter/python"
)

func getPyData(fContent []byte) ([]string, error) {
	parser := ts.NewParser()
	parser.SetLanguage(tspy.GetLanguage())

	tree, err := parser.ParseCtx(context.Background(), nil, fContent)
	if err != nil {
		return nil, err
	}

	rootNode := tree.RootNode()
	q, err := ts.NewQuery([]byte(`
(function_definition
  name: (identifier) @name
  parameters: (parameters)? @params
  return_type: (_)? @return-type
  )
`), tspy.GetLanguage())

	if err != nil {
		return nil, err
	}

	qc := ts.NewQueryCursor()

	qc.Exec(q, rootNode)

	var elements []string

	var fName string
	var fParams string
	var fReturnT string
	var fMatchedNode *ts.Node
	for {
		fMatch, cOk := qc.NextMatch()
		if !cOk {
			break
		}
		for _, capture := range fMatch.Captures {
			fMatchedNode = capture.Node

			switch fMatchedNode.Type() {
			case "identifier":
				fName = fMatchedNode.Content(fContent)
			case "parameters":
				fParams = fMatchedNode.Content(fContent)
			default:
				// TODO: This is not the best way to get the return type; find a better way
				fReturnT = " -> " + fMatchedNode.Content(fContent)
			}
		}

		elem := fmt.Sprintf("def %s%s%s", fName, fParams, fReturnT)

		elements = append(elements, elem)
	}
	return elements, nil

}
