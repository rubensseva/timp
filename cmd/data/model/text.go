package model

// Text represented a text
type Text struct {
	text   string
	author string
}

func NewText(text string, author string) Text {
  return Text{text, author}
}

func NewTextCopy(t Text) Text {
  return Text{t.text, t.author}
}

func (t Text) GetText() string {
  return t.text
}

func (t Text) GetAuthor() string {
  return t.author
}

