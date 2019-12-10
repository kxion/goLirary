package main

import (
	"fmt"
	"go/ast"
	"go/build"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"log"
	"math"
	"os"
	"path/filepath"
	"strings"
)

type structDir string

func main() {
	a := 1000 / 2000
	log.Printf("%T", a)
	fmt.Println(math.Ceil(float64(a)))
	fmt.Println(math.Ceil(float64(1000) / float64(2000)))
	fmt.Println(len(strings.Split("", "")))
	fmt.Println(len(strings.Split("", ",")))
}
func main1() {
	path := "goLirary/process"
	fset := token.NewFileSet()
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}
	aPkgs, _ := parser.ParseDir(fset, filepath.Join(gopath, "src", path), nil, parser.Mode(0))
	conf := types.Config{Importer: importer.Default()}
	for _, aPkg := range aPkgs {
		files := []*ast.File{}
		for _, f := range aPkg.Files {
			files = append(files, f)
		}
		pkg, _ := conf.Check(path, fset, files, nil)
		for _, declName := range pkg.Scope().Names() {
			if !strings.Contains(declName, "New") {
			}
		}
	}
}
