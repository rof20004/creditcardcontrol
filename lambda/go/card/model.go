package card

// Card model
type Card struct {
	ID     int64  `fauna:"_id"`
	Number int64  `fauna:"number"`
	Brand  string `fauna:"brand"`
}
