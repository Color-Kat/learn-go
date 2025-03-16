package link

import (
	"gorm.io/gorm"
	"math/rand"
)

type Link struct {
	Url  string `json:"url"`
	Hash string `json:"hash" gorm:"uniqueIndex"`
	gorm.Model
}

func NewLink(url string) *Link {
	link := &Link{
		Url: url,
	}
	link.GenerateHash()
	return link
}

func (link *Link) GenerateHash() {
	link.Hash = RandStrRunes(6)
}

var letterRunes = []rune("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM")

func RandStrRunes(n int) string {
	runes := make([]rune, n)
	for i := range runes {
		runes[i] = letterRunes[rand.Intn(n)]
	}

	return string(runes)
}
