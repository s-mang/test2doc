package parse

import (
	"go/doc"
	"go/parser"
	"go/token"
	"os"
	"regexp"
)

var (
	docMode = doc.AllMethods

	goTestFileRegexp = regexp.MustCompile("_test.go")
)

func GetPackageDoc(dir string) (*doc.Package, error) {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, dir, isNotGoTestFile, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	// return the first pkg
	// TODO: support multiple packages in dir
	for pkgName, pkgAST := range pkgs {
		importPath := dir + "/" + pkgName
		return doc.New(pkgAST, importPath, docMode), nil
	}

	return nil, nil
}

func isNotGoTestFile(finfo os.FileInfo) bool {
	return !goTestFileRegexp.MatchString(finfo.Name())
}
