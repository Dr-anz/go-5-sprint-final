package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(datastring string) (err error)
	PrActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		// перебираем все значения слайса dataset в цикле
		err := dp.Parse(data) // распарсить значение с помощью метода Parse()
		if err != nil {       // обработать ошибку и залогировать её
			log.Println(err)
			continue
		}
		// сформировать строку с информацией об активности
		info, err := dp.ActionInfo()
		if err != nil {
			log.Println(err)
			continue
		}
		// вывести информацию об активности
		fmt.Println(info)
	}
}
