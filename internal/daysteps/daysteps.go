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
	// разделяем строку на части
	parts := strings.Split(datastring, ",")
	// проверяем что частей ровно 2
	if len(parts) != 2 {
		return errors.New("the parameters are incorrect")
	}
	// парсим количество шагов
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return err
	}
	if steps <= 0 {
		return errors.New("steps must be greater than zero")
	}
	ds.Steps = steps

	// парсинг длительности прогулки
	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return err
	}
	if duration <= 0 {
		return errors.New("duration must be greater than zero")
	}
	ds.Duration = duration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	// вычисляем дистанцию
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	// расчитываем количество сожженых калорий
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	// формирует итоговою строку
	result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distance, calories)

	return result, nil
}
