package nonil

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "nonil",
	Doc:  "reports instances of nil values",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(ast.Node) bool {
			return false // TODO: implement this
		})
	}

	return nil, nil
}
