package pkg

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/build"
	"go/parser"
	"go/printer"
	"go/token"
	"go/types"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/astrewrite"
)

// AutoCommentDir ...
func AutoCommentDir(dir string) {
	pkg, err := build.ImportDir(dir, 0)
	autoCommentImportedPkg(pkg, err)
}

func autoCommentImportedPkg(pkg *build.Package, err error) {
	if err != nil {
		if _, nogo := err.(*build.NoGoError); nogo {
			return
		}
		_, _ = fmt.Fprintln(os.Stderr, err)
		return
	}

	files := make([]string, 0)

	files = append(files, pkg.GoFiles...)
	if pkg.Dir != "." {
		for i, f := range files {
			files[i] = filepath.Join(pkg.Dir, f)
		}
	}

	readingFiles(files...)
}

func AutoCommentFiles(files ...string) {
	readingFiles(files...)
}

func readingFiles(files ...string) {
	fileBodyMap := make(map[string][]byte)

	for _, file := range files {
		fileBody, err := ioutil.ReadFile(file)
		//log.Println(file)
		if err != nil {
			log.Println("+++++++++",err)
			continue
		}

		fileBodyMap[file] = fileBody
	}

	autoCmntr := AutoCommenter{}
	_ = autoCmntr.AutoCommentFiles(fileBodyMap)
}

// AutoCommenter ...
type AutoCommenter struct{}

type pkg struct {
	fileSet   *token.FileSet
	files     map[string]*file
	typesPkg  *types.Package
	typesInfo *types.Info
}

type file struct {
	pkg      *pkg
	f        *ast.File
	fset     *token.FileSet
	src      []byte
	filename string
}

// AutoCommentFiles ...
func (auto *AutoCommenter) AutoCommentFiles(filesMap map[string][]byte) error {
	pkg := &pkg{
		fileSet: token.NewFileSet(),
		files:   make(map[string]*file),
	}

	var packageName string

	for fileName, body := range filesMap {
		f, err := parser.ParseFile(pkg.fileSet, fileName, body, parser.ParseComments)
		if err != nil {
			return err
		}

		if packageName == "" {
			packageName = f.Name.Name
		} else if f.Name.Name != packageName {
			return fmt.Errorf("%s is in package %s, not %s", fileName, f.Name.Name, packageName)
		}

		pkg.files[fileName] = &file{
			pkg:      pkg,
			f:        f,
			fset:     pkg.fileSet,
			src:      body,
			filename: fileName,
		}
	}

	if len(pkg.files) != 0 {
		return pkg.autoComment()
	}

	return nil
}

func (pkg *pkg) autoComment() error {
	for _, file := range pkg.files {
		file.autoComment()
	}
	return nil
}

func (file *file) autoComment() {
	if strings.HasSuffix(file.filename, "_test.go") {
		return
	}

	var comments []*ast.CommentGroup
	ast.Inspect(file.f, func(node ast.Node) bool {
		c, ok := node.(*ast.CommentGroup)
		if ok {
			comments = append(comments, c)
		}

		fn, ok := node.(*ast.FuncDecl)
		if ok {
			if fn.Name.IsExported() && fn.Doc.Text() == "" {
				comment := &ast.Comment{
					Text:  "// " + fn.Name.Name + " ...",
					Slash: fn.Pos() - 1,
				}

				cg := &ast.CommentGroup{
					List: []*ast.Comment{comment},
				}
				fn.Doc = cg

				fmt.Printf("exported function declaration without documentation found on line %d: \n\t%s\n", file.fset.Position(fn.Pos()).Line, fn.Name.Name)
			}
		}

		return true
	})

	file.f.Comments = comments

	reWriteFunc := func(node ast.Node) (ast.Node, bool) {

		return node, true
	}

	newAst := astrewrite.Walk(file.f, reWriteFunc)
	var buf bytes.Buffer
	_ = printer.Fprint(&buf, file.fset, newAst)

	f, err := os.OpenFile(file.filename, os.O_RDWR|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	f.Seek(0, 0)

	f.Write(buf.Bytes())
	f.Sync()

}

type functionSpec struct {
	Name   string
	Prefix string
	Kind   string
}

func (file *file) isLintedFuncDoc(fn *ast.FuncDecl) (*functionSpec, error) {
	if !ast.IsExported(fn.Name.Name) {

		return nil, nil
	}
	kind := "function"
	name := fn.Name.Name
	prefix := fn.Name.Name + " "
	if fn.Doc == nil {
		return &functionSpec{
			Name:   name,
			Prefix: prefix,
			Kind:   kind,
		}, fmt.Errorf("exported %s %s should have comment or be unexported", kind, name)
	}
	s := fn.Doc.Text()

	if !strings.HasPrefix(s, prefix) {
		return &functionSpec{
			Name:   name,
			Prefix: prefix,
			Kind:   kind,
		}, fmt.Errorf(`comment on exported %s %s should be of the form "%s..."`, kind, name, prefix)
	}
	return nil, nil
}
