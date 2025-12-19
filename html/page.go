package html

import (
	"io"
	"net/http"

	"github.com/invertedbit/gocook/viewmodels"
	"maragu.dev/gomponents"
)

type PageRenderer interface {
	Render(w io.Writer) error
}

// type Layout interface {
// 	Render(w io.Writer, pageContent gomponents.Node) error
// }

type Page struct {
	Title           string
	PageContent     gomponents.Node
	LayoutViewModel *viewmodels.LayoutViewModel
}

func (p *Page) Render(w io.Writer) error {
	layout := Layout{
		LayoutViewModel: p.LayoutViewModel,
	}
	return layout.Render(w, p.PageContent)
}

func (p Page) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.Render(w)
}
