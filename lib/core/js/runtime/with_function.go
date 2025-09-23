package runtime

import (
	"github.com/dop251/goja"
	"main/lib/core/js"
)

// WithFunction sets a function.
func WithFunction(runtime *goja.Runtime, name string, call js.Function) error {
	return runtime.Set(name, call)
}
