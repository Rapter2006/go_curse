package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	"gitlab.2gis.ru/i.yatsevich/go-course-1/product"
)

func main() {
	clientId := os.Args[1]
	clientSecret := os.Args[2]
	//idParam := os.Args[3]
	timestamp := fmt.Sprintf("%d", time.Now().Unix()) //Nano()/1000000000)

	signatureParts := map[string]string{
		"client_id": clientId,
		//"id":               idParam,
		"x-auth-timestamp": timestamp,
	}

	fmt.Printf("timestamp: %s \nsignature: %s\n", timestamp, makeSignature(clientSecret, signatureParts))

	fmt.Println(generateSignature(timestamp, "user", "secretKey"))
	// Часть 1. Хранилище товаров
	//part1()
	//
	// часть 2. Дерево категорий
	// part2()
}

func generateSignature(ts, id, secretKey string) (string, error) {
	message := fmt.Sprintf("x-auth-timestamp:%s;x-client-id:%s", ts, id)
	h := hmac.New(sha256.New, []byte(secretKey))
	_, err := h.Write([]byte(message))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(h.Sum(nil)), nil
}

//1525938080
//1525938086035

// makeSignature формирует подпись для работы с трастовыми методами
// В качестве параметров должны быть переданы все query параметры и заголовки.
func makeSignature(secretKey string, params map[string]string) string {
	parts := make([]string, len(params))
	i := 0
	for name, value := range params {
		parts[i] = strings.ToLower(name) + ":" + strings.ToLower(value)
		i++
	}

	sort.SliceStable(parts, func(i, j int) bool {
		return strings.Compare(parts[i], parts[j]) < 0
	})

	payload := strings.Join(parts, ";")

	mac := hmac.New(sha256.New, []byte(secretKey))
	mac.Write([]byte(payload))
	sign := mac.Sum(nil)

	return base64.StdEncoding.EncodeToString(sign)
}

func showProductList(items []product.Item) {
	fmt.Printf("Всего товаров: %d", len(items))
	fmt.Println()

	for _, item := range items {
		showProductItem(item, nil)
	}
}

func showProductItem(item product.Item, err error) {
	if err != nil {
		fmt.Printf("Ошибка при отображении товара: %s", err.Error())
	} else {
		fmt.Printf("Товар: %q", item)
	}
	fmt.Println()
}
