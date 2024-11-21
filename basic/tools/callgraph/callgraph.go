package callgraph

import (
	"log"

	"golang.org/x/tools/go/callgraph"
	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/pointer"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
)

func Analysis(packageDir string) {

	analysisProject(packageDir)

}

func analysisProject(packageDir string) (*callgraph.Graph, error) {
	pgs, err := packages.Load(&packages.Config{
		Mode:  packages.LoadAllSyntax,
		Dir:   packageDir,
		Tests: false,
		Fset:  nil,
		Env:   nil,
	})

	if err != nil {
		return nil, err
	}

	program, ssaPkgs := ssautil.Packages(pgs, ssa.BuilderMode(0))
	program.Build()
	log.Panicf("ssaPkgs: %v", ssaPkgs)

	result, err := pointer.Analyze(&pointer.Config{
		Mains:          ssaPkgs,
		BuildCallGraph: true,
	})
	if err != nil {
		return nil, err
	}

	return result.CallGraph, nil
}
