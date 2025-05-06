package pages

import f "github.com/razshare/frizzante"

func Default(context f.PageContext) {
	path, view, base, action := context()
	path("/")
	view(f.ViewReference("Login"))
	base(loginBaseFunction)
	action(loginActionFunction)
}
