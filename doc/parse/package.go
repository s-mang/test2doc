package parse

import (
	"fmt"
	"go/doc"
	"go/parser"
	"go/token"
	"os"
	"regexp"
)

var (
	testFileRegexp = regexp.MustCompile("_test.go")

	// go/doc stores package 'Funcs' as a slice
	// - we need to look up documentation by func name

	funcsMap map[string]*doc.Func
)

// get description from go/doc for apib doc
//  - lookup by func name

// NewPackageDoc retrieves the go/doc package for the given dir
func NewPackageDoc(dir string) (*doc.Package, error) {
	pkgDoc, err := getPackageDoc(dir)
	if err != nil {
		return nil, err
	}

	setDocFuncsMap(pkgDoc)
	return pkgDoc, nil
}

func setDocFuncsMap(pkgDoc *doc.Package) {
	typeFuncsMap := getPkgTypesFunctions(pkgDoc)
	funcsMap = make(map[string]*doc.Func, len(pkgDoc.Funcs)+len(typeFuncsMap))

	for _, fn := range pkgDoc.Funcs {
		funcsMap[fn.Name] = fn
	}

	for k, fn := range typeFuncsMap {
		funcsMap[k] = fn
	}
}

func getPkgTypesFunctions(pkgDoc *doc.Package) map[string]*doc.Func {
	result := make(map[string]*doc.Func)
	for _, t := range pkgDoc.Types {
		for _, f := range t.Methods {
			if f.Doc != "" {
				result[fmt.Sprintf("%s.%s", t.Name, f.Name)] = f
			}
		}
	}

	return result
}

func getPackageDoc(dir string) (*doc.Package, error) {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, dir, isNotGoTestFile, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	// return the first pkg
	for pkgName, pkgAST := range pkgs {
		importPath := dir + "/" + pkgName
		return doc.New(pkgAST, importPath, doc.AllDecls), nil
	}

	return nil, nil
}

func isNotGoTestFile(finfo os.FileInfo) bool {
	return !testFileRegexp.MatchString(finfo.Name())
}
