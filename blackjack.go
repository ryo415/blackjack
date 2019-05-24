package main

import (
	"fmt"
	"strconv"
	"math/rand"
	"time"
)

type card struct {
	mark string
	num int
}

func NumToCard(num int) string {

	switch num {
	case 1:
		return "A"
	case 11:
		return "J"
	case 12:
		return "Q"
	case 13:
		return "K"
	default:
		return strconv.Itoa(num)
	}
}

func NumToPoint(num int) int {

	switch num {
	case 11,12,13:
		return 10
	default:
		return num
	}
}

func Initialize_Deck() []card {
	marks := []string{"ハート", "スペード", "ダイヤ", "クラブ"}
	var deck []card
	var count int = 0

	deck = make([]card, 52)

	for _, mark := range marks {
		for i := 1;i <= 13;i++ {
			deck[count].mark = mark
			deck[count].num = i
			count++
		}
	}

	return deck
}

func Print_Deck(deck []card) {
	for i := range deck {
		fmt.Printf("%vの%v\n",deck[i].mark, deck[i].num)
	}
}

func Print_Card(print_card card, secret int) {
	if secret == 0 {
		fmt.Printf("%vの%v\n", print_card.mark, NumToCard(print_card.num))
	}else{
		fmt.Printf("****\n")
	}
}

func Drow_Card(deck []card) card {
	var drow_card card
	var drow_num int

	rand.Seed(time.Now().UnixNano())
	drow_num = rand.Intn(len(deck))
	drow_card.mark = deck[drow_num].mark
	drow_card.num = deck[drow_num].num

	return drow_card
}

func Delete_Card(deck []card, delete_card card) []card {
	result_deck := []card{}

	for _, cards := range deck {
		if cards.mark != delete_card.mark || cards.num != delete_card.num {
			result_deck = append(result_deck, cards)
		}
	}

	return result_deck
}

func Drow_Phase(deck []card, point int, turn string, secret int) ([]card, int) {
	var drow_card = Drow_Card(deck)
	deck = Delete_Card(deck, drow_card)
	point += NumToPoint(drow_card.num)

	if turn == "player" {
		fmt.Print("あなた: ")
	}else if turn == "croupier" {
		fmt.Print("ディーラー: ")
	}
	Print_Card(drow_card, secret)

	return deck, point
}

func Initialize_game(deck []card, user_point int, cpu_point int) ([]card, int, int) {
	var turn string = "player"

	fmt.Println("---ブラックジャック---")
	fmt.Println("ゲーム開始")

	deck, user_point = Drow_Phase(deck, user_point, turn, 0)
	deck, user_point = Drow_Phase(deck, user_point, turn, 0)

	turn = "croupier"
	deck, cpu_point = Drow_Phase(deck, cpu_point, turn, 0)
	deck, cpu_point = Drow_Phase(deck, cpu_point, turn, 1)

	return deck, user_point, cpu_point

}

func Player_Turn(deck []card, user_point int) ([]card, int) {
	var cont string
	var turn string = "player"

	for {
		fmt.Printf("あなたの現在の得点は%v\n", user_point)
		fmt.Println("カードを引きますか? Y/N")
		fmt.Scan(&cont)

		if cont == "Y" || cont == "y"{
			deck, user_point = Drow_Phase(deck, user_point, turn, 0)
			if user_point > 21 {
				fmt.Println("バーストしました")
				user_point = -1
				return deck, user_point
			}
		} else if cont == "N" || cont == "n" {
			return deck, user_point
		} else {
			fmt.Println("YかNを入力")
		}
	}
}

func Croupier_Turn(deck []card, cpu_point int) ([]card, int) {
	var turn string = "croupir"

	for {
		fmt.Printf("ディーラーの現在の得点は%v\n", cpu_point)
		if cpu_point < 17 {
			deck, cpu_point = Drow_Phase(deck, cpu_point, turn, 0)
		} else {
			return deck, cpu_point
		}
	}
}

func Print_Result(user_point, cpu_point int) {
	if user_point > cpu_point || cpu_point > 21 {
		fmt.Println("あなたの勝ち")
	} else if user_point == cpu_point {
		fmt.Println("引き分け")
	} else {
		fmt.Println("あなたの負け")
	}
}

func main() {
	var cpu_point int = 0
	var user_point int = 0
	var deck = Initialize_Deck()

	deck, user_point, cpu_point = Initialize_game(deck, user_point, cpu_point)

	deck, user_point = Player_Turn(deck, user_point)
	if user_point < 0 {

	} else {
		deck, cpu_point = Croupier_Turn(deck, cpu_point)

		fmt.Printf("あなたのポイント: %v\n", user_point)
		fmt.Printf("ディーラーのポイント: %v\n", cpu_point)
		Print_Result(user_point, cpu_point)
	}

}
