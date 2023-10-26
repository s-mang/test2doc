package parse

import (
	"go/doc"
	"go/parser"
	"go/token"
	"os"
	"regexp"
	"strings"
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
	funcsMap = make(map[string]*doc.Func, len(pkgDoc.Funcs))
	for _, tp := range pkgDoc.Types {
		for _, fn := range tp.Methods {
			fn.Doc = strings.Replace(fn.Doc, "\t", "    ", -1)
			funcsMap["("+fn.Recv+")"+fn.Name] = fn
		}
	}
	for _, fn := range pkgDoc.Funcs {
		// replace any \t character in comment with spaces
		// because apiary parser cannot handle tabs
		fn.Doc = strings.Replace(fn.Doc, "\t", "    ", -1)
		funcsMap[fn.Name] = fn
	}
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
