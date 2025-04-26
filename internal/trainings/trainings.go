package trainings

type Training struct {
	// TODO: добавить поля
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
	t.TrainingType = s1[1]
	t.Duration, err = time.ParseDuration(s1[2])
	if err != nil {
		log.Println(err)
		return err
	}

}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
}
