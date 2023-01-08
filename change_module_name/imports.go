package change_module_name

import (
	"fmt"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

const goExt = ".go"

type (
	Imports	struct {
		currentModuleName	ModuleName
		newModuleName		ModuleName
	}

	Dir	struct {
		Path string
	}

	File	struct {
		Path string
	}
)

func NewImports(currentModuleName ModuleName, newModuleName ModuleName) *Imports {
	return &Imports{currentModuleName: currentModuleName, newModuleName: newModuleName}
}

func (i *Imports) Walker() error {
	var queue Queue[Dir]
	elems, err := os.ReadDir(".")
	if err != nil {
		return err
	}

	for _, elem := range elems {
		if elem.IsDir() {
			queue.Push(Dir{Path: elem.Name()})
			continue
		}
		if filepath.Ext(elem.Name()) != goExt {
			continue
		}

		if err = i.fixImports(File{Path: elem.Name()}); err != nil {
			return err
		}
	}

	for !queue.Empty() {
		head := queue.Pop()
		elems, err = os.ReadDir(head.Path)
		if err != nil {
			return err
		}

		for _, elem := range elems {
			if elem.IsDir() {
				queue.Push(Dir{Path: fmt.Sprintf("%s/%s", head.Path, elem.Name())})
				continue
			}

			if err := i.fixImports(
				File{
					Path: fmt.Sprintf("%s/%s", head.Path, elem.Name()),
				}); err != nil {
				return err
			}
		}
	}

	return nil
}

func (i *Imports) fixImports(file File) error {
	fSet := token.NewFileSet()
	astFile, err := parser.ParseFile(fSet, file.Path, nil, parser.ParseComments)
	if err != nil {
		return err
	}

	for _, e := range astFile.Imports {
		e.Path.Value = strings.Replace(
			e.Path.Value,
			i.currentModuleName.ToString(),
			i.newModuleName.ToString(),
			-1,
		)
	}

	f, err := os.OpenFile(file.Path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)

	err = printer.Fprint(f, fSet, astFile)
	if err != nil {
		return err
	}

	return nil
}
