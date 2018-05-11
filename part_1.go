package main

import (
	"log"
	"os"

	"gitlab.2gis.ru/i.yatsevich/go-course-1/product"
)

func part1() {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	repo := product.NewRepository(logger)
	repo.Add(product.Item{
		ID:          "1",
		Name:        "Вальсакор таб. п/о 80мг №5",
		Description: "",
	})
	repo.Add(product.Item{
		ID:   "2",
		Name: "Iphone 8 256Gb silver",
	})
	repo.Add(product.Item{
		ID:          "3",
		Name:        "IRhone 12 c антеной ",
		Description: "Новый IRhone 12 c антеной для просмотра телевизора и 5 sim картами",
	})

	showProductItem(repo.FindByID(product.ItemID("1")))
	showProductItem(repo.FindByID(product.ItemID("666")))

	showProductList(repo.FindByText("iphone"))
	showProductList(repo.FindByText("5"))
	showProductList(repo.FindByText("Айфон"))
}
