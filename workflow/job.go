package workflow

type Job struct {
	JobName string            `yaml:"name"`
	Action  ActionName        `yaml:"action"`
	With    map[string]string `yaml:"with"`
	Env     map[string]string `yaml:"env"`
	Output  string            `yaml:"output"`
}

func (job *Job) Execute() error {
	return nil
}
