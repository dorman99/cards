package main

func main() {
	cards := newdeck()
	cards.shuffle()
	cards.print()
}
