package check

type Check struct {
	name string
}

func New(name string) Check {
	return Check{name: name}
}

func (check *Check) Name() string {
	return check.name
}
