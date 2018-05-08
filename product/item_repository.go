package product

// ItemsRepository наше хранилище товаров в памяти
type ItemsRepository struct {
	items  map[ItemID]Item
	logger ...
}

// Add добавить версию объекта в хранилище.
// Повторное добавление одного товара - обновление - не приводит к ошибкам,
// мы просто обновляем в репозитории сохраненную копию
func (r *ItemsRepository) Add(item ...) error {
	panic("Реализуй меня")
}

// FindByID найти необходимый нам товар по его идентификатору
// Если ничего не найдено - то возвращаем nil
// Т.к. у нас может быть только один обхект с уникальным ID - то и результат работы этого метода - всегда один объект
func (r *ItemsRepository) FindByID(id ItemID) (..., error) {
	panic("Реализуй меня")
}

// FindByText поиск по подстроке всех товаров
// Люди обычно ищут товары не по идентификаторам, а по какому то текстовому описанию
func (r *ItemsRepository) FindByText(text string) []... {
	panic("Реализуй меня")
}
