package actioninfo

import (
	"fmt"

	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
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
			fmt.Println(spentenergy.ErrData)
			continue
		}

		ai, err := dp.ActionInfo()
		if err != nil {
			fmt.Println(spentenergy.ErrData)
			continue
		}
		fmt.Println(ai)
	}
}
