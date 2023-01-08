package change_module_name

import (
	_ "embed"
	"errors"
	"log"
)

var (
	errEmptyModuleName = errors.New("moduleName is missing")
)

func ChangeModuleName() {
	currentModuleName, err := getCurrentModuleName()
	if err != nil {
		log.Fatal(err)
	}

	newModuleName, err := getNewModuleName()
	if err != nil {
		log.Fatal(err)
	}

	if err := NewImports(currentModuleName, newModuleName).Walker(); err != nil {
		log.Fatal(err)
	}
}
