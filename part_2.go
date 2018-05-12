package main

import (
	"fmt"

	"gitlab.2gis.ru/i.yatsevich/go-course-1/nestedset"
	"gitlab.2gis.ru/i.yatsevich/go-course-1/product"
)

func part2() {
	categoryRoot := nestedset.NewNode(
		product.Category{
			ID:   product.CategoryID("1"),
			Name: "Каталог товаров",
		},
	)

	goods := categoryRoot.Add(nestedset.NewNode(
		product.Category{
			ID:   product.CategoryID("2"),
			Name: "Товары",
		},
	))
	tires := goods.Add(nestedset.NewNode(
		product.Category{
			ID:   product.CategoryID("3"),
			Name: "Шины",
		},
	))
	tires.Add(nestedset.NewNode(
		product.Item{
			ID:   product.ItemID("100"),
			Name: "Nokian Hakkapelitta-5 15 195X65",
		},
	))
	goods.Add(nestedset.NewNode(
		product.Category{
			ID:   product.CategoryID("4"),
			Name: "Диски",
		},
	))

	services := categoryRoot.Add(nestedset.NewNode(
		product.Category{
			ID:   product.CategoryID("3"),
			Name: "Услуги",
		},
	))
	services.Add(nestedset.NewNode(
		product.Item{
			ID:   product.ItemID("1"),
			Name: "Эпиляция зоны бикини",
		},
	))

	// Обойти дерево и выбрать только товары
	fmt.Println("Выбираем все товары в нашем каталоге")
	productsInCategories := findProducts(*categoryRoot)
	showProductList(productsInCategories)
}

func findProducts(node nestedset.Node) []product.Item {
	panic("Реализуй меня")
}
