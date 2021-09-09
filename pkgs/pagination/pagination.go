package pagination

import (
	"fmt"
	"github.com/slava-vishnyakov/paginator"
)

func Pagination(page, count, per_page int) string {
	p := paginator.Paginator(page, per_page, count, 2)
	res := ""
	for _, page := range p {
		var class string
		var disable string
		if page.IsActive {
			class = "active"
		} else {
			class = ""
		}
		if page.IsDisabled {
			disable = "disabled"
		} else {
			disable = ""
		}

		res += fmt.Sprintf(`<li class="page-item %s %s"><a class="page-link" href="?page=%v">%s</a></li>`, class, disable, page.Page, page.Label)
	}
	return res
}
