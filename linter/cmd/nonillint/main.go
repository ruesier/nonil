package main

import (
	"github.com/ruesier/nonil/linter/nonil"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(nonil.Analyzer)
}
