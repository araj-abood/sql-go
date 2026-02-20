package config

type Command struct {
	Name      string
	Arguments []string
}

type isArgsEmptyInterface interface {
	IsArgsEmpty(moreThan int) bool
}

func (c Command) IsArgsEmpty(moreThan int) bool {
	if len(c.Arguments) < moreThan {
		return true
	}
	return false
}
