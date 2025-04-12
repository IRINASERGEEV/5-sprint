package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	// TODO: добавить поля
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	parsim := strings.Split(datastring, ",")
	if len(parsim) != 2 {
		return spentenergy.ErrData
	}

	stepsStr := parsim[0]
	steps, err := strconv.Atoi(stepsStr)
	if err != nil {
		return spentenergy.ErrData
	}
	if steps <= 0 {
		return spentenergy.ErrData
	}

	ds.Steps = steps

	durationStr := parsim[1]
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return spentenergy.ErrData
	}
	if duration <= 0 {
		return spentenergy.ErrData
	}
	ds.Duration = time.Duration(duration)
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	distanceKm := spentenergy.Distance(ds.Steps, ds.Height)
	if ds.Steps <= 0 || ds.Duration <= 0 || ds.Height <= 0 || ds.Weight <= 0 {
		return "", spentenergy.ErrData
	}

	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", spentenergy.ErrData
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distanceKm, calories), nil
}
