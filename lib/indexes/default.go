package indexes

import f "github.com/razshare/frizzante"

func Default(
	route func(path string, page string),
	show func(showFunction f.PageFunction),
	action func(actionFunction f.PageFunction),
) {
	route("/", "login")
	show(loginShowFunction)
	action(loginActionFunction)
	return
}
