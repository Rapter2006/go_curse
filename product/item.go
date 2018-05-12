package product

package product

// Item товар описывается набором полей из структуры, и пкоа не имеет своего поведения из бизнес логики
type Item struct {
	ID          ItemID
	Name        string
	Description string
}

// ItemID свой тип для идентификатора товаров,
// чтобы его всегда можно было отличить от идентификатора другого типа объектов
type ItemID string

// String приведение к примитивным видам данных лучше делать на стороне кастомного типа данных
func (i ItemID) String() string {
	return string(i)
}
