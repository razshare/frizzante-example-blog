//go:build types

package types

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"main/lib/core/files"
)

func Generate[T any]() (err error) {
	var value T

	type_ := reflect.TypeOf(value)

	var packages = map[string][]string{}
	var definitions = map[string]map[string][]string{}
	if _, err = Define(type_, packages, definitions); err != nil {
		return
	}

	if !files.IsDirectory(filepath.Join(".gen", "types")) {
		if err = os.MkdirAll(filepath.Join(".gen", "types"), os.ModePerm); err != nil {
			return
		}
	}

	befores := []string{
		"main",
		"main",
	}
	after := "main"
	packagePath := type_.PkgPath()

	for _, before := range befores {
		packagePath = strings.ReplaceAll(packagePath, before, after)
	}

	dname := filepath.Join(".gen", "types", strings.ReplaceAll(packagePath, "/", string(filepath.Separator)))
	if files.IsDirectory(dname) {
		if err = os.RemoveAll(dname); err != nil {
			return
		}
	}

	if err = os.MkdirAll(dname, os.ModePerm); err != nil {
		return
	}

	parts := strings.Split(type_.PkgPath(), "/")
	count := len(parts)
	package_ := parts[count-1]

	var globalBuilder strings.Builder
	var namespaceBuilder strings.Builder
	globalBuilder.WriteString(fmt.Sprintf("export type %s = %s.%s\n\n", type_.Name(), package_, type_.Name()))
	for namespace, definition := range definitions {
		namespaceBuilder.Reset()
		namespaceBuilder.WriteString(fmt.Sprintf("export declare namespace %s {\n", namespace))
		for name, lines := range definition {
			namespaceBuilder.WriteString(fmt.Sprintf("    export type %s = {\n", name))
			for _, line := range lines {
				namespaceBuilder.WriteString(fmt.Sprintf("        %s\n", line))
			}
			namespaceBuilder.WriteString("    }\n")
		}
		namespaceBuilder.WriteString("}\n\n")
		globalBuilder.WriteString(strings.TrimSpace(namespaceBuilder.String()))
		globalBuilder.WriteString("\n\n")
	}

	fname := filepath.Join(dname, type_.Name()+".d.ts")
	if err = os.WriteFile(fname, []byte(strings.TrimSpace(globalBuilder.String())), os.ModePerm); err != nil {
		return
	}

	return
}
