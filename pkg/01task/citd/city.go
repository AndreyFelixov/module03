package citd

import (
	"01task/wordz"
)

func City() string {
	wordz.Words = []string{"Пенза", "Ростов-На-Дону", "Нижний Новгород", "Кузнецк", "Санкт-Петербург", "Нижний Тагил"}
	wordz.Prefix = ""
	return wordz.Random()
}
