package borrow_model

type Book struct {
	id      string
	title   string
	special bool
}

func (b *Book) IsSpecial() bool {
	return b.special
}

func (b *Book) MarkAsSpecial() {
	b.special = true
}
