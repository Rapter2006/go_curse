package product

import (
	"log"
	"sync"
	"errors"
	"strings"
)
 var ErrNotFound = errors.New("item not found")

// ItemsRepository наше хранилище товаров в памяти
type ItemsRepository struct {
	m sync.Mutex
	items  map[ItemID]Item
	logger log.Logger //... // TODO нужно как то указать тип компонента
}

func NewRepository(log log.Logger) ItemsRepository {
		return ItemsRepository{
			items: make(map[ItemID]Item, 0),
			logger: log,
		}
}

// Add добавить версию объекта в хранилище.
// Повторное добавление одного товара - обновление - не приводит к ошибкам,
// мы просто обновляем в репозитории сохраненную копию
func (r *ItemsRepository) Add(item Item) error {
	r.m.Lock()
	defer r.m.Unlock()
	_, ok := r.items[item.ID]
	// TODO для дебага полезно знать что была замена существующего товара
	if ok {
		r.logger.Printf("Заменяем уже существующий товар ID:%s - %q", item)
	}
	r.items[item.ID] = item
	r.logger.Printf("Успешно добавили новый товар %q", item)
	return nil
}

// FindByID найти необходимый нам товар по его идентификатору
// Если ничего не найдено - то возвращаем ошибку
// Т.к. у нас может быть только один обхект с уникальным ID - то и результат работы этого метода - всегда один объект
func (r *ItemsRepository) FindByID(id ItemID) (Item, error) {
	r.m.Lock()
	defer r.m.Unlock()
	value, ok := r.items[id]
	if !ok {
		return Item{}, ErrNotFound
	}
	return value, nil
}

// FindByText поиск по подстроке всех товаров
// Люди обычно ищут товары не по идентификаторам, а по какому то текстовому описанию
func (r *ItemsRepository) FindByText(text string) []Item {
    items := make([]Item, 0)
	r.m.Lock()
	defer r.m.Unlock()

	for _, item := range r.items {
		if strings.Contains(item.Name, text) {
			r.logger.Printf(
				"Товар '%s' совпал в названии при поиске по фразе '%s'",
				text,
			)
			items = append(items, item)
		}
	}

	return items
}
