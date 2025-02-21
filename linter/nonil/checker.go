package nonil

import (
	"go/ast"
	"regexp"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "nonil",
	Doc:  "reports instances of nil values",
	Run:  run,
}

var skipComment = regexp.MustCompile(`^\s+nonil:unsafe`)

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(node ast.Node) bool {
			switch exp := node.(type) {
			case *ast.Ident:
				if exp.Name == "nil" {
					pass.Reportf(exp.Pos(), "nil used")
				}
			case *ast.ValueSpec:
				if len(exp.Values) == 0 && zeroHasNil(exp.Type) {
					pass.Reportf(exp.Pos(), "referenced types not initialized")
				}
			case *ast.Comment:
				// TODO: return false if file contains an ignore comment
				return !skipComment.MatchString(exp.Text)
			}
			return true
		})
	}

	return nil, nil
}

func zeroHasNil(node ast.Node) bool {
	switch t := node.(type) {
	case *ast.ArrayType:
		if t.Len == nil {
			return true
		}
		return zeroHasNil(t.Elt)
	case *ast.InterfaceType, *ast.MapType, *ast.FuncType, *ast.ChanType, *ast.StarExpr:
		return true
	case *ast.StructType:
		for _, f := range t.Fields.List {
			if zeroHasNil(f.Type) {
				return true
			}
		}
	}
	return false
}
