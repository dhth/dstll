package tsutils

import (
	"context"
	"fmt"

	ts "github.com/smacker/go-tree-sitter"
	tsgo "github.com/smacker/go-tree-sitter/golang"
)

func getGoFuncs(resultChan chan<- Result, fContent []byte) {
	parser := ts.NewParser()
	parser.SetLanguage(tsgo.GetLanguage())

	tree, err := parser.ParseCtx(context.Background(), nil, fContent)
	if err != nil {
		resultChan <- Result{Err: err}
		return
	}

	rootNode := tree.RootNode()

	q, err := ts.NewQuery([]byte(`
(function_declaration
  name: (identifier) @name
  parameters: (_)? @params
  result: (_)? @return-type
  )
`), tsgo.GetLanguage())

	if err != nil {
		resultChan <- Result{Err: err}
		return
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
			case "parameter_list":
				fParams = fMatchedNode.Content(fContent)
			default:
				// TODO: This is not the best way to get the return type; find a better way
				fReturnT = " " + fMatchedNode.Content(fContent)
			}
		}

		elem := fmt.Sprintf("func %s%s%s", fName, fParams, fReturnT)

		elements = append(elements, elem)
	}
	resultChan <- Result{Results: elements}

}

func getGoMethods(resultChan chan<- Result, fContent []byte) {
	parser := ts.NewParser()
	parser.SetLanguage(tsgo.GetLanguage())

	tree, err := parser.ParseCtx(context.Background(), nil, fContent)
	if err != nil {
		resultChan <- Result{Err: err}
		return
	}

	rootNode := tree.RootNode()

	q, err := ts.NewQuery([]byte(`
(method_declaration
  receiver: (parameter_list) @rec
  name: (field_identifier) @name
  parameters: (_)? @params
  result: (_)? @return-type
  )
`), tsgo.GetLanguage())

	if err != nil {
		resultChan <- Result{Err: err}
		return
	}

	qc := ts.NewQueryCursor()

	qc.Exec(q, rootNode)

	var elements []string

	var fRec string
	var recQueried bool
	var fName string
	var fParams string
	var fReturnT string
	var fMatchedNode *ts.Node

	for {
		fMatch, cOk := qc.NextMatch()
		if !cOk {
			break
		}

		recQueried = false
		for _, capture := range fMatch.Captures {
			fMatchedNode = capture.Node

			switch fMatchedNode.Type() {
			case "field_identifier":
				fName = fMatchedNode.Content(fContent)
			case "parameter_list":
				if recQueried {
					fParams = fMatchedNode.Content(fContent)
				} else {
					fRec = fMatchedNode.Content(fContent)
					recQueried = true
				}
			default:
				// TODO: This is not the best way to get the return type; find a better way
				fReturnT = " " + fMatchedNode.Content(fContent)
			}
		}

		elem := fmt.Sprintf("func %s %s%s%s", fRec, fName, fParams, fReturnT)

		elements = append(elements, elem)
	}
	resultChan <- Result{Results: elements}

}
