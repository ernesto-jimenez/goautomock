package parseutil

import (
	"errors"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

// ErrTooManyPackages is returned when there is more than one package in a
// directory where there should only be one Go package.
var ErrTooManyPackages = errors.New("more than one package found in a directory")

// PackageAST returns the AST of the package at the given path.
func PackageAST(path string) (pkg *ast.Package, err error) {
	return parseAndFilterPackages(path, func(k string, v *ast.Package) bool {
		return !strings.HasSuffix(k, "_test")
	})
}

// PackageTestAST returns the AST of the test package at the given path.
func PackageTestAST(path string) (pkg *ast.Package, err error) {
	return parseAndFilterPackages(path, func(k string, v *ast.Package) bool {
		return strings.HasSuffix(k, "_test")
	})
}

type packageFilter func(string, *ast.Package) bool

// filteredPackages filters the parsed packages and then makes sure there is only
// one left.
func parseAndFilterPackages(path string, filter packageFilter) (pkg *ast.Package, err error) {
	fset := token.NewFileSet()
	srcDir, err := DefaultGoPath.Abs(path)
	if err != nil {
		return nil, err
	}

	pkgs, err := parser.ParseDir(fset, srcDir, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	pkgs = filterPkgs(pkgs, filter)

	if len(pkgs) != 1 {
		return nil, ErrTooManyPackages
	}

	for _, p := range pkgs {
		pkg = p
	}

	return
}

func filterPkgs(pkgs map[string]*ast.Package, filter packageFilter) map[string]*ast.Package {
	filtered := make(map[string]*ast.Package)
	for k, v := range pkgs {
		if filter(k, v) {
			filtered[k] = v
		}
	}

	return filtered
}
