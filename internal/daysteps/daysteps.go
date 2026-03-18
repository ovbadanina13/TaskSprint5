package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	slice := strings.Split(datastring, ",")
	if len(slice) != 2 {
		return fmt.Errorf("wait for 2 params, got %d", len(slice))
	}
	steps, err := strconv.Atoi(slice[0])
	if err != nil {
		return fmt.Errorf("error convert to int: %v", err)
	}
	if steps <= 0 {
		return errors.New("amount of steps mustn't be 0")
	}
	ds.Steps = steps

	duration, err := time.ParseDuration(slice[1])
	if err != nil {
		return fmt.Errorf("error to convert time: %v", err)
	}
	if duration <= 0 {
		return errors.New("duration must be positive")
	}
	ds.Duration = duration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	finalDistance := spentenergy.Distance(ds.Steps, ds.Height)
	var calories float64
	var err error
	calories, err = spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, finalDistance, calories), nil

}
