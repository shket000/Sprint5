package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		err := dp.Parse(data)
		if err != nil {
			// Логируем ошибку и переходим к следующей итерации
			log.Println("Ошибка парсинга:", err)
			continue
		}

		info, err2 := dp.ActionInfo()
		if err2 != nil {
			// Логируем ошибку и переходим к следующей итерации
			log.Println("Ошибка получения информации:", err)
			continue
		}
		fmt.Println(info)
	}

}
