package daysteps

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

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {

	if datastring == "" {
		log.Println("Datastring is empty")
		return errors.New("datastring is empty")
	}

	s1 := strings.Split(datastring, ",")
	if len(s1) != 2 {
		log.Println("Datastring is invalid")
		return errors.New("datastring is invalid")
	}
	ds.Steps, err = strconv.Atoi(s1[0])
	if err != nil {
		log.Println("Steps is invalid")
		return errors.New("steps is invalid")
	}
	ds.Duration, err = time.ParseDuration(s1[1])
	if err != nil {
		log.Println("Duration is invalid")
		return errors.New("duration is invalid")
	}

	if ds.Steps <= 0 {
		return errors.New("steps must be greater than zero")
	}
	if ds.Duration <= 0 {
		return errors.New("duration must be greater than zero")
	}
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	dist := spentenergy.Distance(ds.Steps, ds.Height)
	if ds.Steps <= 0 {
		log.Println("Steps is invalid")
		return "", errors.New("steps is invalid")
	}
	if ds.Duration <= 0 {
		log.Println("Duration is invalid")
		return "", errors.New("duration is invalid")
	}
	callWalk, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, dist, callWalk), nil
}
