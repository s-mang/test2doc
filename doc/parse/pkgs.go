package parse

import (
	"go/ast"
	"go/doc"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"regexp"
)

var (
	files            = ".go"
	docMode doc.Mode = 0
)

func GetPackageDocs(dir string) []*doc.Package {
	pkgs := getPackages(dir)

	var pkgDocs = make([]*doc.Package, 0, len(pkgs))
	for pkgName, pkgAST := range pkgs {
		importPath := dir + "/" + pkgName
		pkgDoc := doc.New(pkgAST, importPath, docMode)

		pkgDocs = append(pkgDocs, pkgDoc)
	}

	return pkgDocs
}

func FilterByFnName(fnName string, pkgDoc *doc.Package) {
	pkgDoc.Filter(func(name string) bool {
		return regexp.MustCompile("^" + name + "$").MatchString(fnName)
	})
}

func getPackages(dir string) (pkgs map[string]*ast.Package) {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, ".", isGoFile, parser.ParseComments)
	if err != nil {
		panic("parser.ParseDir: " + err.Error())
	}

	return pkgs
}
func isGoFile(fi os.FileInfo) bool {
	name := fi.Name()

	// clean up
	return !fi.IsDir() &&
		len(name) > 0 && name[0] != '.' && // ignore .files
		filepath.Ext(name) == ".go"
}

// func printDocFuncs(d *doc.Package) {
// 	fmt.Printf("\nFuncs:\n")
// 	for _, fn := range d.Funcs {
// 		fmt.Printf(" %s:\n", fn.Name)
// 		indent := "        "
// 		indentedDoc := strings.Replace(fn.Doc, "\n", "\n"+indent, -1)
// 		fmt.Printf("  - Doc: \n%s%s\n", indent, indentedDoc)
// 	}
// }

// func printDocNotes(d *doc.Package) {
// 	fmt.Println("\nNotes:")
// 	for k, noteSet := range d.Notes {
// 		fmt.Printf(" %s:\n", k)
// 		for _, note := range noteSet {
// 			fmt.Printf("    - UID: %s\n    - Body: %s\n", note.UID, note.Body)
// 		}

// 	}
// }
