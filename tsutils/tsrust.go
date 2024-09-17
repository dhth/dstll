package tsutils

import (
	"context"
	"fmt"

	ts "github.com/smacker/go-tree-sitter"
	tsrust "github.com/smacker/go-tree-sitter/rust"
)

const (
	rustNodeTypeVisibilityModifier = "visibility_modifier"
	rustNodeTypeTypeParameters     = "type_parameters"
	rustNodeTypeParameters         = "parameters"
)

func getRustTypes(resultChan chan<- Result, fContent []byte) {
	query := `[
  (struct_item) @struct
  (enum_item) @enum
  (type_item) @type_item
  (trait_item) @trait
  (union_item) @union
]`
	results, err := getGenericResult(fContent, query, tsrust.GetLanguage())
	resultChan <- Result{Results: results, Err: err}
}

func getRustFuncs(resultChan chan<- Result, fContent []byte) {
	parser := ts.NewParser()
	parser.SetLanguage(tsrust.GetLanguage())

	tree, err := parser.ParseCtx(context.Background(), nil, fContent)
	if err != nil {
		resultChan <- Result{Err: err}
		return
	}

	rootNode := tree.RootNode()

	q, err := ts.NewQuery([]byte(`
(function_item
    (visibility_modifier)? @visibility
    name: (_) @identifier
    type_parameters: (_)? @type_parameters
    parameters: (_)? @parameter_list
    return_type: (_)? @return_type
)
`), tsrust.GetLanguage())
	if err != nil {
		resultChan <- Result{Err: err}
		return
	}

	qc := ts.NewQueryCursor()

	qc.Exec(q, rootNode)

	var elements []string

	for {
		fMatch, cOk := qc.NextMatch()
		if !cOk {
			break
		}

		var visibilityModifier string
		var fName string
		var fTParams string
		var fParams string
		var fReturnT string
		var fMatchedNode *ts.Node

		for _, capture := range fMatch.Captures {
			fMatchedNode = capture.Node

			switch fMatchedNode.Type() {
			case rustNodeTypeVisibilityModifier:
				visibilityModifier = fMatchedNode.Content(fContent) + " "
			case nodeTypeIdentifier:
				fName = fMatchedNode.Content(fContent)
			case rustNodeTypeTypeParameters:
				fTParams = fMatchedNode.Content(fContent)
			case rustNodeTypeParameters:
				fParams = fMatchedNode.Content(fContent)
			default:
				// TODO: This is not the best way to get the return type; find a better way
				fReturnT = " -> " + fMatchedNode.Content(fContent)
			}
		}

		elem := fmt.Sprintf("%sfn %s%s%s%s", visibilityModifier, fName, fTParams, fParams, fReturnT)

		elements = append(elements, elem)
	}
	resultChan <- Result{Results: elements}
}
