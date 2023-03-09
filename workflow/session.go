package workflow

type Session struct {
	Jobs []Job `yaml:"job"`
}

func (session *Session) Execute() error {
	for _, job := range session.Jobs {
		err := job.Execute()
		if err != nil {
			return err
		}
	}
	return nil
}
