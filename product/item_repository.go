package product

// ItemsRepository наше хранилище товаров в памяти
type ItemsRepository struct {
	items  map[ItemID]Item
	logger ... // TODO нужно как то указать тип компонента
}

// Add добавить версию объекта в хранилище.
// Повторное добавление одного товара - обновление - не приводит к ошибкам,
// мы просто обновляем в репозитории сохраненную копию
func (r *ItemsRepository) Add(item ...) error {
	// TODO для дебага полезно знать что была замена существующего товара
	if ... {
		r.logger.Printf("Заменяем уже существующий товар ID:%s - %q", item)
	}

	r.logger.Printf("Успешно добавили новый товар %q", item)

	panic("Реализуй меня")
}

// FindByID найти необходимый нам товар по его идентификатору
// Если ничего не найдено - то возвращаем ошибку
// Т.к. у нас может быть только один обхект с уникальным ID - то и результат работы этого метода - всегда один объект
func (r *ItemsRepository) FindByID(id ItemID) (..., error) {
	panic("Реализуй меня")
}

// FindByText поиск по подстроке всех товаров
// Люди обычно ищут товары не по идентификаторам, а по какому то текстовому описанию
func (r *ItemsRepository) FindByText(text string) []... {

	// TODO например лог сообщение в данном случае может быть таким
	r.logger.Printf(
		"Товар '%s' совпал в названии при поиске по фразе '%s'",
		item.
		text,
	)
	// TODO Или другим, или их может быть несоклько - на ваше усмотрение - можете проявить фантазию

	panic("Реализуй меня")
}
