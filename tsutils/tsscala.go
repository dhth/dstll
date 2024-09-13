package tsutils

import (
	"context"
	"fmt"

	ts "github.com/smacker/go-tree-sitter"
	tsscala "github.com/smacker/go-tree-sitter/scala"
)

func getScalaClasses(resultChan chan<- Result, fContent []byte) {
	parser := ts.NewParser()
	parser.SetLanguage(tsscala.GetLanguage())

	tree, err := parser.ParseCtx(context.Background(), nil, fContent)
	if err != nil {
		resultChan <- Result{Err: err}
		return
	}

	rootNode := tree.RootNode()

	classQuery, err := ts.NewQuery([]byte(`
(class_definition
  (modifiers)? @mod
  name: (identifier) @class-name
  type_parameters: (type_parameters)? @typeparams
  class_parameters: (class_parameters)? @cparams
  extend: (extends_clause)? @extends-clause
  )
`), tsscala.GetLanguage())
	if err != nil {
		resultChan <- Result{Err: err}
		return
	}

	classQueryCur := ts.NewQueryCursor()

	classQueryCur.Exec(classQuery, rootNode)

	var elements []string

	for {
		classMatch, cOk := classQueryCur.NextMatch()

		if !cOk {
			break
		}

		var cModifiers string
		var cName string
		var cTypeParams string
		var cParams string
		var cExtendsCl string
		var cMatchedNode *ts.Node
		for _, capture := range classMatch.Captures {
			cMatchedNode = capture.Node

			switch cMatchedNode.Type() {
			case nodeTypeModifiers:
				cModifiers = cMatchedNode.Content(fContent) + " "
			case nodeTypeIdentifier:
				cName = cMatchedNode.Content(fContent)
			case "type_parameters":
				cTypeParams = cMatchedNode.Content(fContent)
			case "class_parameters":
				cParams = cMatchedNode.Content(fContent)
			case "extends_clause":
				cExtendsCl = " " + cMatchedNode.Content(fContent)
			}
		}

		elem := fmt.Sprintf("%sclass %s%s%s%s", cModifiers, cName, cTypeParams, cParams, cExtendsCl)
		elements = append(elements, elem)
	}
	resultChan <- Result{Results: elements}
}

func getScalaObjects(resultChan chan<- Result, fContent []byte) {
	parser := ts.NewParser()
	parser.SetLanguage(tsscala.GetLanguage())

	tree, err := parser.ParseCtx(context.Background(), nil, fContent)
	if err != nil {
		resultChan <- Result{Err: err}
		return
	}

	rootNode := tree.RootNode()

	objectQuery, err := ts.NewQuery([]byte(`
(object_definition
  name: (identifier) @name
  extend: (extends_clause)? @extends-clause
  )
`), tsscala.GetLanguage())
	if err != nil {
		resultChan <- Result{Err: err}
		return
	}

	objectQueryCur := ts.NewQueryCursor()

	objectQueryCur.Exec(objectQuery, rootNode)

	var elements []string

	for {
		objectMatch, cOk := objectQueryCur.NextMatch()

		if !cOk {
			break
		}

		var oName string
		var oExtendsCl string
		var oMatchedNode *ts.Node
		for _, capture := range objectMatch.Captures {
			oMatchedNode = capture.Node

			switch oMatchedNode.Type() {
			case nodeTypeIdentifier:
				oName = oMatchedNode.Content(fContent)
			case "extends_clause":
				oExtendsCl = " " + oMatchedNode.Content(fContent)
			}
		}

		elem := fmt.Sprintf("object %s%s", oName, oExtendsCl)

		elements = append(elements, elem)
	}
	resultChan <- Result{Results: elements}
}

func getScalaFunctions(resultChan chan<- Result, fContent []byte) {
	parser := ts.NewParser()
	parser.SetLanguage(tsscala.GetLanguage())

	tree, err := parser.ParseCtx(context.Background(), nil, fContent)
	if err != nil {
		resultChan <- Result{Err: err}
		return
	}

	rootNode := tree.RootNode()
	funcQuery, err := ts.NewQuery([]byte(`
	(function_definition
     (modifiers)?  @access-modifier
	 name: (identifier) @fname
	 parameters: (parameters)? @fparams
	 return_type: (_)? @return-type
	)
	    `), tsscala.GetLanguage())
	if err != nil {
		resultChan <- Result{Err: err}
		return
	}
	funcQueryCur := ts.NewQueryCursor()

	funcQueryCur.Exec(funcQuery, rootNode)

	var fMatchedNode *ts.Node
	var elements []string

	for {
		funcMatch, fOk := funcQueryCur.NextMatch()

		if !fOk {
			break
		}

		var fAccessModifier string
		var fIdentifer string
		var fParams string
		var fReturnType string
		for _, capture := range funcMatch.Captures {
			fMatchedNode = capture.Node
			switch fMatchedNode.Type() {
			case nodeTypeModifiers:
				fAccessModifier = fMatchedNode.Content(fContent) + " "
			case nodeTypeIdentifier:
				fIdentifer = fMatchedNode.Content(fContent)
			case nodeTypeParameters:
				fParams = fMatchedNode.Content(fContent)
			default:
				// TODO: This is not the best way to get the return type; find a better way
				fReturnType = ": " + fMatchedNode.Content(fContent)
			}
		}

		elem := fmt.Sprintf("%sdef %s%s%s", fAccessModifier, fIdentifer, fParams, fReturnType)
		elements = append(elements, elem)
	}
	resultChan <- Result{Results: elements}
}
