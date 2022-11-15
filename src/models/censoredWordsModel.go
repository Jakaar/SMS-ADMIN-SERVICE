package model

type CensoredWords struct {
	MN string `json:"mn"`
}

type CensoredWordsName interface {
	TableName() string
}

func (CensoredWords) TableName() string {
	return "censored_words"
}
