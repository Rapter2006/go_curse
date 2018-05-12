package product

type Category struct {
	ID    CategoryID
	Name  string
	Items []*Item
}

type CategoryID string

func (c CategoryID) String() string {
	return string(c)
}
