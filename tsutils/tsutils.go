package tsutils

import (
	"context"

	ts "github.com/smacker/go-tree-sitter"
)

func getGenericResult(fContent []byte, query string, language *ts.Language) ([]string, error) {
	parser := ts.NewParser()
	parser.SetLanguage(language)

	tree, err := parser.ParseCtx(context.Background(), nil, fContent)
	if err != nil {
		return nil, err
	}

	rootNode := tree.RootNode()

	q, err := ts.NewQuery([]byte(query), language)
	if err != nil {
		return nil, err
	}

	qc := ts.NewQueryCursor()

	qc.Exec(q, rootNode)

	var elements []string

	var result string
	for {
		tMatch, cOk := qc.NextMatch()
		if !cOk {
			break
		}
		if len(tMatch.Captures) != 1 {
			continue
		}
		result = tMatch.Captures[0].Node.Content(fContent)

		elements = append(elements, result)
	}
	return elements, nil
}
