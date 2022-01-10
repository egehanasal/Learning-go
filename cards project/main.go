package main

//import "fmt"

func main() {
	//var card string = "Ace of Spades"
	// var card = "asdsd"
	// this means card is a variable can be any type
	// var card string = "blah blah"
	// this means card ia variable that wlil always contain a string
	// var card = " " == card:=" "
	// if you are reassigning a new value to a variable, you **can't** write card:= " "

	// for loop
	//for i := 1; i<= 5; i++{
	//	fmt.Println(i)
	//}

	// array: fixed length
	// slice: not fixed length
	//cards := []string{newCard(), "Ace of Diamonds"}
	//cards = append(cards, "Six of Spades")
	//for i, card := range cards{
	//	fmt.Println(i, card)
	//}
	//type deck []string
	//greeting := "Hi there"
	//fmt.Println([]byte(greeting)) //[72 105 32 116 104 101 114 101]
	
	cards := newDeckFromFile("my_cards")
	cards.shuffle()
	cards.print()
	//cards.saveToFile("my_cards")
	//cards := newDeckFromFile("my_cards")


	//hand, remainingCards := deal(cards, 5)
	//hand.print()
	//fmt.Println()
	//remainingCards.print()
}
