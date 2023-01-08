package change_module_name

import (
	_ "embed"
	"golang.org/x/mod/modfile"
	"os"
)

//go:embed go.mod
var gomod []byte

type (
	ModuleName string
)

func (m *ModuleName) ToString() string {
	return string(*m)
}

func (m *ModuleName) Empty() bool {
	return m.ToString() == ""
}

func getNewModuleName() (ModuleName, error) {
	if len(os.Args) == 1 {
		return "", errEmptyModuleName
	}
	moduleName := ModuleName(os.Args[1])
	if moduleName.Empty() {
		return "", errEmptyModuleName
	}
	return moduleName, nil
}

func getCurrentModuleName() (ModuleName, error) {
	f, err := modfile.Parse("go.mod", gomod, nil)
	if err != nil {
		return "", nil
	}
	return ModuleName(f.Module.Mod.Path), nil
}
