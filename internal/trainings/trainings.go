package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	slice := strings.Split(datastring, ",")
	if len(slice) != 3 {
		return fmt.Errorf("wait for 3 params, got %d", len(slice))
	}
	steps, err := strconv.Atoi(slice[0])
	if err != nil {
		return fmt.Errorf("error convert to int: %v", err)
	}
	if steps <= 0 {
		return errors.New("amount of steps mustn't be 0")
	}
	t.Steps = steps

	t.TrainingType = slice[1]

	duration, err := time.ParseDuration(slice[2])
	if err != nil {
		return fmt.Errorf("error to convert time: %v", err)
	}
	if duration <= 0 {
		return errors.New("duration must be positive")
	}
	t.Duration = duration

	return nil
}

func (t Training) ActionInfo() (string, error) {

	finalDistance := spentenergy.Distance(t.Steps, t.Height)
	finalSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	var calories float64
	var err error

	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
	default:
		return "", errors.New("неизвестный тип тренировки")
	}
	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), finalDistance, finalSpeed, calories), nil

}
