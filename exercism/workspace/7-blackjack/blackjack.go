package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
    var value int;
    switch card {
        case "ace":
            value = 11;
        case "two":
            value = 2;
        case "three":
            value = 3;
        case "four":
            value = 4;
        case "five":
            value = 5;
        case "six":
            value = 6;
        case "seven":
            value = 7;
        case "eight":
            value = 8;
        case "nine":
            value = 9;
        case "ten":
            value = 10;
        case "jack":
            value = 10;
        case "queen":
            value = 10;
        case "king":
            value = 10;
        default:
            value = 0;
    }
    return value;
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
    switch {
        case card1 == "ace" && card1 == card2:
            return "P";
        case card1 == "ace" && card1 == card2:
            return "P";
        case ParseCard(card1) + ParseCard(card2) == 21 && ParseCard(dealerCard) < 10:
            return "W";
        case ParseCard(card1) + ParseCard(card2) <= 20 && ParseCard(card1) + ParseCard(card2) >= 17:
            return "S";
        case ParseCard(card1) + ParseCard(card2) <= 16 && ParseCard(card1) + ParseCard(card2) >= 12:
            switch {
                case ParseCard(dealerCard) >= 7:
                    return "H";
                default:
                    return "S";
            }
        case ParseCard(card1) + ParseCard(card2) <= 11:
            return "H";
    }
    return "S"
}

// the Ideal Solution is:
/*
// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten", "jack", "queen", "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	handScore := ParseCard(card1) + ParseCard(card2)
	dealerScore := ParseCard(dealerCard)
	switch {
	case handScore == 22:
		return "P"
	case handScore == 21:
		if dealerScore < 10 {
			return "W"
		} else {
			return "S"
		}
	case handScore >= 17 || handScore >= 12 && dealerScore < 7:
		return "S"
	default:
		return "H"
	}
}
*/