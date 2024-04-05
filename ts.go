package main

import (
	"context"
	"fmt"

	ts "github.com/smacker/go-tree-sitter"
	tsgo "github.com/smacker/go-tree-sitter/golang"
	tspy "github.com/smacker/go-tree-sitter/python"
	tsscala "github.com/smacker/go-tree-sitter/scala"
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
  parameters: (parameters) @params
  return_type: (_)? @return-type
  )
`), tspy.GetLanguage())

	if err != nil {
		return nil, err
	}

	qc := ts.NewQueryCursor()

	qc.Exec(q, rootNode)

	var elements []string

	for {
		fMatch, cOk := qc.NextMatch()
		if !cOk {
			break
		}

		fName := fMatch.Captures[0].Node
		fParams := fMatch.Captures[1].Node
		var fReturnT string
		if len(fMatch.Captures) == 3 {
			fReturnT = fMatch.Captures[2].Node.Content(fContent)
		}

		var elem string
		if fReturnT == "" {
			elem = fmt.Sprintf("def %s%s", fName.Content(fContent), fParams.Content(fContent))
		} else {
			elem = fmt.Sprintf("def %s%s -> %s", fName.Content(fContent), fParams.Content(fContent), fReturnT)
		}

		elements = append(elements, elem)
	}
	return elements, nil

}

func getGoData(fContent []byte) ([]string, error) {
	parser := ts.NewParser()
	parser.SetLanguage(tsgo.GetLanguage())

	tree, err := parser.ParseCtx(context.Background(), nil, fContent)
	if err != nil {
		return nil, err
	}

	rootNode := tree.RootNode()

	q, err := ts.NewQuery([]byte(`
(function_declaration
  name: (identifier) @name
  parameters: (_) @params
  result: (_) @return-type
  )
`), tsgo.GetLanguage())

	if err != nil {
		return nil, err
	}

	qc := ts.NewQueryCursor()

	qc.Exec(q, rootNode)

	var elements []string

	for {
		fMatch, cOk := qc.NextMatch()
		if !cOk {
			break
		}

		fName := fMatch.Captures[0].Node
		fParams := fMatch.Captures[1].Node
		fReturnT := fMatch.Captures[2].Node

		elem := fmt.Sprintf("func %s%s %s", fName.Content(fContent), fParams.Content(fContent), fReturnT.Content(fContent))

		elements = append(elements, elem)
	}
	return elements, nil

}

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
  class_parameters: (class_parameters) @params
  ) @class
`), tsscala.GetLanguage())

	if err != nil {
		return nil, err
	}

	classQueryCur := ts.NewQueryCursor()

	funcQuery, err := ts.NewQuery([]byte(`
	(function_definition
	  name: (identifier) @name
	  parameters: (parameters) @params
	  return_type: (_) @return-type
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
		cName := classMatch.Captures[1].Node
		cParams := classMatch.Captures[2].Node

		funcQueryCur.Exec(funcQuery, cNode)

		elem := fmt.Sprintf("class %s%s", cName.Content(fContent), cParams.Content(fContent))

		for {
			funcMatch, fOk := funcQueryCur.NextMatch()

			if !fOk {
				break
			}
			fName := funcMatch.Captures[0].Node.Content(fContent)
			fParams := funcMatch.Captures[1].Node.Content(fContent)
			fReturnT := funcMatch.Captures[2].Node.Content(fContent)
			elem += fmt.Sprintf("\n\tdef %s%s: %s", fName, fParams, fReturnT)
		}

		elements = append(elements, elem)
	}
	return elements, nil

}
