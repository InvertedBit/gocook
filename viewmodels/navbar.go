package viewmodels

type NavbarMenuItem struct {
	Label    string
	Link     string
	Children []*NavbarMenuItem
}

type NavbarViewModel struct {
	Items []*NavbarMenuItem
}

func NewNavbarViewModel() *NavbarViewModel {
	return &NavbarViewModel{
		Items: []*NavbarMenuItem{},
	}
}

func (n *NavbarViewModel) AddItem(item *NavbarMenuItem) {
	n.Items = append(n.Items, item)
}
