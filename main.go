package main

import (
	"fmt"

	"gitlab.2gis.ru/market/go-course/product"
)

func main() {
	// Часть 1. Хранилище товаров
	//part1()
	//
	// часть 2. Дерево категорий
	part2()
}

//func showProductList(items []product.Item) {
//	fmt.Printf("Всего товаров: %d", len(items))
//	fmt.Println()
//
//	for _, item := range items {
//		showProductItem(item, nil)
//	}
//}
//
//func showProductItem(item product.Item, err error) {
//	if err != nil {
//		fmt.Printf("Ошибка при отображении товара: %s", err.Error())
//	} else {
//		fmt.Printf("Товар: %q", item)
//	}
//	fmt.Println()
//}
