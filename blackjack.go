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

func Print_Mark(mark_string string) string {
	switch mark_string {
	case "ハート":
		return "💛"
	case "スペード":
		return "♠"
	case "ダイヤ":
		return "♦"
	case "クラブ":
		return "♣"
	}

	return "Error"
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
		fmt.Printf("%v %v\n", Print_Mark(print_card.mark), NumToCard(print_card.num))
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

func Drow_Phase(deck []card, cards []card, turn string, secret int) ([]card, []card) {
	var drow_card = Drow_Card(deck)
	deck = Delete_Card(deck, drow_card)
	cards = Add_Have_Card(cards, drow_card)

	if turn == "player" {
		fmt.Print("あなた: ")
	}else if turn == "croupier" {
		fmt.Print("ディーラー: ")
	}
	Print_Card(drow_card, secret)

	return deck, cards
}

func Initialize_game(deck []card, user_cards []card, cpu_cards []card) ([]card, []card, []card) {
	var turn string = "player"

	fmt.Println("---ブラックジャック---")
	fmt.Println("ゲーム開始")

	deck, user_cards = Drow_Phase(deck, user_cards, turn, 0)
	deck, user_cards = Drow_Phase(deck, user_cards, turn, 0)

	turn = "croupier"
	deck, cpu_cards = Drow_Phase(deck, cpu_cards, turn, 0)
	deck, cpu_cards = Drow_Phase(deck, cpu_cards, turn, 1)

	return deck, user_cards, cpu_cards

}

func Player_Turn(deck []card, user_cards []card) ([]card, []card, bool) {
	var cont string
	var turn string = "player"
	var burst_flag bool

	for {
		fmt.Printf("あなたの現在の得点は%v\n", Print_Point(user_cards))
		fmt.Println("カードを引きますか? Y/N")
		fmt.Scan(&cont)

		if cont == "Y" || cont == "y"{
			deck, user_cards = Drow_Phase(deck, user_cards, turn, 0)
			if Print_Point(user_cards) > 21 {
				fmt.Println("バーストしました")
				burst_flag = true
				return deck, user_cards, burst_flag
			}
		} else if cont == "N" || cont == "n" {
			return deck, user_cards, burst_flag
		} else {
			fmt.Println("YかNを入力")
		}
	}
}

func Croupier_Turn(deck []card, cpu_cards []card) ([]card, []card) {
	var turn string = "croupier"

	for {
		fmt.Printf("ディーラーの現在の得点は%v\n", Print_Point(cpu_cards))
		if Print_Point(cpu_cards) < 17 {
			deck, cpu_cards = Drow_Phase(deck, cpu_cards, turn, 0)
		} else {
			return deck, cpu_cards
		}
	}
}

func Print_Result(user_cards, cpu_cards []card) {
	user_point := Print_Point(user_cards)
	cpu_point := Print_Point(cpu_cards)

	if user_point > cpu_point || cpu_point > 21 {
		fmt.Println("あなたの勝ち")
	} else if user_point == cpu_point {
		fmt.Println("引き分け")
	} else {
		fmt.Println("あなたの負け")
	}
}

func Add_Have_Card(cards []card, add_card card)  []card {
	cards = append(cards, add_card)

	return cards
}

func Print_Point(cards []card) int {
	var point int
	var ace_flag bool
	for i:=0;i<len(cards);i++ {
		point += NumToPoint(cards[i].num)
		if cards[i].num == 1 {
			ace_flag = true
		}
	}

	if ace_flag == true && point+10 < 21 {
		point += 10
	}

	return point
}

func main() {
	var user_cards []card
	var cpu_cards []card
	var deck = Initialize_Deck()
	var burst_flag bool

	deck, user_cards, cpu_cards = Initialize_game(deck, user_cards, cpu_cards)

	deck, user_cards, burst_flag = Player_Turn(deck, user_cards)

	if burst_flag == true {

	} else {
		deck, cpu_cards = Croupier_Turn(deck, cpu_cards)

		fmt.Printf("あなたのポイント: %v\n", Print_Point(user_cards))
		fmt.Printf("ディーラーのポイント: %v\n", Print_Point(cpu_cards))
		Print_Result(user_cards, cpu_cards)
	}

}
