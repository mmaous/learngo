package techpalace

import s "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + s.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	stars := s.Repeat("*", numStarsPerLine)

	return stars + "\n" + welcomeMsg + "\n" + stars
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	withoutStars := s.ReplaceAll(oldMsg, "*", "")
	withoutSpace := s.TrimSpace(withoutStars)
	return withoutSpace
}
