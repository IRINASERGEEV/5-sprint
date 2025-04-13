package actioninfo

import (
	"fmt"
)

type DataParser interface {
	// TODO: добавить методы
	Parse(datastring string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	// TODO: реализовать функцию
	for _, v := range dataset {
		err := dp.Parse(v)
		if err != nil {
			fmt.Println("error in Parsing:", err)
			continue
		}

		ai, err := dp.ActionInfo()
		if err != nil {
			fmt.Println("error in forming the activity line:", err)
			continue
		}
		fmt.Println(ai)
	}
}
