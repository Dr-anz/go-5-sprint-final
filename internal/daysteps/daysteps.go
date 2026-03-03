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
	Steps    int
	Duration time.Duration
	Personal personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	// разделяем строку на части
	parts := strings.Split(datastring, ",")
	// проверяем что частей ровно 2
	if len(parts) != 2 {
		return err
	}
	// парсим количество шагов
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return err
	}
	ds.Steps = steps
	// парсинг длительности прогулки
	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return err
	}
	ds.Duration = duration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	// вычисляем дистанцию
	distance := spentenergy.Distance(ds.Steps, ds.Personal.Height)
	// расчитываем количество сожженых калорий
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	// формирует итоговою строку
	result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила: %.2f км.\nВы сожгли: %.2f ккал.\n", ds.Steps, distance, calories)

	return result, nil
}
