package product

type Category struct {
	ID    CategoryID
	Name  string
	Items []*Item
}

func (c Category) GetName() string {
	return c.Name
}

type CategoryID string

func (c CategoryID) String() string {
	return string(c)
}
