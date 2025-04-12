package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	// TODO: добавить поля
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	parsim := strings.Split(datastring, ",")
	if len(parsim) != 3 {
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
	t.Steps = steps

	t.TrainingType = parsim[1]

	durationStr := parsim[2]
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return spentenergy.ErrData
	}
	if duration <= 0 {
		return spentenergy.ErrData
	}
	t.Duration = duration
	return
}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	var distanceKm, speed float64

	distanceKm = spentenergy.Distance(t.Steps, t.Height)
	speed = spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	switch t.TrainingType {
	case "Ходьба":
		calories, err := spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		durationHours := t.Duration.Hours()
		if err != nil {
			return "", spentenergy.ErrData
		}
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, durationHours, distanceKm, speed, calories), nil

	case "Бег":
		calories, err := spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		durationHours := t.Duration.Hours()
		if err != nil {
			return "", spentenergy.ErrData
		}
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, durationHours, distanceKm, speed, calories), nil

	default:
		return "", spentenergy.ErrData
	}
}
