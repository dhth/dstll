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
  type_parameters: (_)? @type-params
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

	for {
		fMatch, cOk := qc.NextMatch()
		if !cOk {
			break
		}

		var fName string
		var fTParams string
		var fParams string
		var fReturnT string
		var fMatchedNode *ts.Node

		parametersSeen := false
		for _, capture := range fMatch.Captures {
			fMatchedNode = capture.Node

			switch fMatchedNode.Type() {
			case nodeTypeIdentifier:
				fName = fMatchedNode.Content(fContent)
			case "type_parameter_list":
				fTParams = fMatchedNode.Content(fContent)
			case "parameter_list":
				if parametersSeen {
					fReturnT = " " + fMatchedNode.Content(fContent)
				} else {
					fParams = fMatchedNode.Content(fContent)
					parametersSeen = true
				}
			default:
				// TODO: This is not the best way to get the return type; find a better way
				fReturnT = " " + fMatchedNode.Content(fContent)
			}
		}

		elem := fmt.Sprintf("func %s%s%s%s", fName, fTParams, fParams, fReturnT)

		elements = append(elements, elem)
	}
	resultChan <- Result{Results: elements}
}

func getGoTypes(resultChan chan<- Result, fContent []byte) {
	parser := ts.NewParser()
	parser.SetLanguage(tsgo.GetLanguage())

	tree, err := parser.ParseCtx(context.Background(), nil, fContent)
	if err != nil {
		resultChan <- Result{Err: err}
		return
	}

	rootNode := tree.RootNode()

	q, err := ts.NewQuery([]byte(`
(type_declaration) @type-dec
`), tsgo.GetLanguage())
	if err != nil {
		resultChan <- Result{Err: err}
		return
	}

	qc := ts.NewQueryCursor()

	qc.Exec(q, rootNode)

	var elements []string

	var typeDec string
	for {
		tMatch, cOk := qc.NextMatch()
		if !cOk {
			break
		}
		if len(tMatch.Captures) != 1 {
			continue
		}
		typeDec = tMatch.Captures[0].Node.Content(fContent)

		elements = append(elements, typeDec)
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

	for {
		fMatch, cOk := qc.NextMatch()
		if !cOk {
			break
		}

		var fRec string
		var fName string
		var fParams string
		var fReturnT string
		var fMatchedNode *ts.Node

		receiverQueried := false
		parametersSeen := false
		for _, capture := range fMatch.Captures {
			fMatchedNode = capture.Node

			switch fMatchedNode.Type() {
			case "field_identifier":
				fName = fMatchedNode.Content(fContent)
			case "parameter_list":
				if !receiverQueried {
					fRec = fMatchedNode.Content(fContent)
					receiverQueried = true
				} else if !parametersSeen {
					fParams = fMatchedNode.Content(fContent)
					parametersSeen = true
				} else {
					fReturnT = " " + fMatchedNode.Content(fContent)
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
