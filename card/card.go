package card

/*
	A Basic Card.
*/
type Card struct {
	/*
		Value is an 8-bit int that stores the face value of the card. Ranges from 1 (Ace) to 13 (King).
		This value is NOT necessarily correlated with the actual game play value of the card.
	*/
	Value int8
	/*
		The suit is one of 'H' (Hearts), 'S' (Spades), 'D' (Diamonds), 'C' (Clubs)
	*/
	Suit rune
}
