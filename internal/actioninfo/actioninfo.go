package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, v := range dataset {
		err := dp.Parse(v)
		if err != nil {
			log.Printf("Ошибка при парсинге данных '%s': %v", v, err)
			continue
		}
		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("ошибка при формировании данных: %v", err)
			continue
		}
		fmt.Println(info)
	}
}
