package parse

import (
	"go/doc"
	"go/parser"
	"go/token"
)

var (
	docMode doc.Mode = 0
)

func getAllPackageDocs(dir string) ([]*doc.Package, error) {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, ".", nil, parser.ParseComments)
	if err != nil {
		return nil, err
	} else if len(pkgs) == 0 {
		return nil, nil
	}

	var pkgDocs = make([]*doc.Package, 0, len(pkgs))
	for pkgName, pkgAST := range pkgs {
		importPath := dir + "/" + pkgName
		pkgDoc := doc.New(pkgAST, importPath, docMode)

		pkgDocs = append(pkgDocs, pkgDoc)
	}

	return pkgDocs, nil
}
