package main

import (
	"fmt"
	"log"
)

func main() {
	cardList1 := []string{"0213-2321", "4215-2321", "5421-2321", "2245-2321", "9865-2321"}
	cardList2 := []string{"0000-2321", "3211-2321", "8213-2321", "9213-2321", "7213-2321"}
	humoMap := make(map[string]string)
	anotherBanksMap := make(map[string]string)
	for _, v := range cardList1 {
		if v[0] >= 53 {
			humoMap[v] = "HUMO"
		} else {
			anotherBanksMap[v] = "Other Bank"
		}
	}
	for _, v := range cardList2 {
		if v[0] >= 53 {
			humoMap[v] = "HUMO"
		} else {
			anotherBanksMap[v] = "Other Bank"
		}
	}
	fmt.Println("Список карточек Хумо:")
	fmt.Println(humoMap)
	input := ""
	count := 0
	fmt.Println("Введите счёт карты!")
	_, err := fmt.Scan(&input)
	if err != nil {
		log.Fatal("Смотри что вводишь!", err)
	}
	fmt.Println("Введенное значение:", input)
	for i := 0; i < 9; i++ {
		if input[i] != '-' && input[i] >= 48 && input[i] <= 57 {
			count++
		}
	}
	val := false
	if input[4] == 45 {
		val = true
	}
	if count == 8 && val == true {
		bal := false
		fmt.Println("Correct format")
		for i := 0; i < len(cardList1); i++ {
			if input == cardList1[i] {
				fmt.Println("This card is already in our Database-1")
				bal = true
				break
			}
		}
		for i := 0; i < len(cardList2); i++ {
			if cardList2[i] == input {
				fmt.Println("This card is already in our Database-2")
				bal = true
				break
			}
		}
		if input[0] >= 53 && bal != true {
			fmt.Println("Wow, its Humo card and we add it to our Database")
			humoMap[input] = "HUMO"
			fmt.Println("HUMO DATABASE: ", humoMap)
		} else if bal != true && input[0] < 53 {
			fmt.Println("ohh damn, its bullshit but we add it to our Database")
			anotherBanksMap[input] = "Another bank"
			fmt.Println("Other banks DATABASE: ", anotherBanksMap)
		}

	} else {
		fmt.Println("Wrong format")
	}
}
