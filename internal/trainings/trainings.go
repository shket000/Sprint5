package trainings

import (
	"errors"
	"fmt"
	"log"
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
	s1 := strings.Split(datastring, ",")
	if len(s1) != 3 {
		return fmt.Errorf("data string does not contain 3 elements")
	}
	t.Steps, err = strconv.Atoi(s1[0])
	if err != nil {
		log.Println(err)
		return err
	}
	if t.Steps < 1 {
		log.Println("Invalid number of steps, should be at least 1")
		return errors.New("Invalid number of steps, should be at least 1")
	}
	t.TrainingType = s1[1]
	t.Duration, err = time.ParseDuration(s1[2])
	if err != nil {
		log.Println(err)
		return err
	}
	if t.Duration <= 0 {
		log.Println("Invalid Duration, should be at least 0")
		return errors.New("Invalid Duration, should be at least 0")
	}
	return nil
}

func (t Training) ActionInfo() (string, error) {
	dur := float64(t.Duration.Hours())
	//durs:=fmt.Sprintf("%.2f", dur)
	dist := spentenergy.Distance(t.Steps, t.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	callRun, err := spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	if err != nil {
		log.Println(err)
		return "", err
	}
	callWalk, err := spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	if err != nil {
		log.Println(err)
		return "", err
	}

	switch t.TrainingType {
	case "Бег":
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, dur, dist, speed, callRun), nil
	case "Ходьба":
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, dur, dist, speed, callWalk), nil
	default:
		return "", errors.New("unknown training type")
	}
}
