package tsutils

import (
	"context"
	"fmt"

	ts "github.com/smacker/go-tree-sitter"
	tsscala "github.com/smacker/go-tree-sitter/scala"
)

func getScalaData(fContent []byte) ([]string, error) {

	parser := ts.NewParser()
	parser.SetLanguage(tsscala.GetLanguage())

	tree, err := parser.ParseCtx(context.Background(), nil, fContent)
	if err != nil {
		return nil, err
	}

	rootNode := tree.RootNode()

	classQuery, err := ts.NewQuery([]byte(`
(class_definition
  name: (identifier) @class-name
  class_parameters: (class_parameters)? @cparams
  extend: (extends_clause)? @extends-clause
  ) @class
`), tsscala.GetLanguage())

	if err != nil {
		return nil, err
	}

	classQueryCur := ts.NewQueryCursor()

	funcQuery, err := ts.NewQuery([]byte(`
	(function_definition
     (modifiers
      (access_modifier) @access-modifier
     )?
	 name: (identifier) @fname
	 parameters: (parameters)? @fparams
	 return_type: (_)? @return-type
	)
	    `), tsscala.GetLanguage())
	if err != nil {
		return nil, err
	}

	funcQueryCur := ts.NewQueryCursor()

	classQueryCur.Exec(classQuery, rootNode)

	var elements []string

	for {
		classMatch, cOk := classQueryCur.NextMatch()

		if !cOk {
			break
		}

		cNode := classMatch.Captures[0].Node

		var cName string
		var cParams string
		var cExtendsCl string
		var cMatchedNode *ts.Node
		for _, capture := range classMatch.Captures[1:] {
			cMatchedNode = capture.Node

			switch cMatchedNode.Type() {
			case "identifier":
				cName = cMatchedNode.Content(fContent)
			case "class_parameters":
				cParams = cMatchedNode.Content(fContent)
			case "extends_clause":
				cExtendsCl = " " + cMatchedNode.Content(fContent)
			}
		}

		funcQueryCur.Exec(funcQuery, cNode)

		elem := fmt.Sprintf("class %s%s%s", cName, cParams, cExtendsCl)

		var fMatchedNode *ts.Node
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
				case "access_modifier":
					fAccessModifier = fMatchedNode.Content(fContent) + " "
				case "identifier":
					fIdentifer = fMatchedNode.Content(fContent)
				case "parameters":
					fParams = fMatchedNode.Content(fContent)
				default:
					// TODO: This is not the best way to get the return type; find a better way
					fReturnType = ": " + fMatchedNode.Content(fContent)
				}
			}

			elem += fmt.Sprintf("\n\t%sdef %s%s%s", fAccessModifier, fIdentifer, fParams, fReturnType)
		}

		elements = append(elements, elem)
	}
	return elements, nil

}
