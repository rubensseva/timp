package model

// Text represented a text
type Text struct {
	text   string
	author string
}

type TextJSON struct {
	Text   string
	Author string
}

func NewText(text string, author string) Text {
	return Text{text, author}
}

func NewTextCopy(t Text) Text {
	return Text{t.text, t.author}
}

func (t Text) ToJSONobj() TextJSON {
	return TextJSON{t.text, t.author}
}

func (t TextJSON) ToRegularObj() Text {
	return Text{t.Text, t.Author}
}

func (t Text) GetText() string {
	return t.text
}

func (t Text) GetAuthor() string {
	return t.author
}

func TextListToJSON(texts []Text) []TextJSON {
	var textsJSON []TextJSON
	for _, text := range texts {
		textsJSON = append(textsJSON, text.ToJSONobj())
	}
	return textsJSON
}

func TextJSONListToRegular(textsJSON []TextJSON) []Text {
	var texts []Text
	for _, textJSON := range textsJSON {
		texts = append(texts, textJSON.ToRegularObj())
	}
	return texts
}
