package people

type People struct{
	Name string
	Age int

}

type Common interface {
	GetString() string
}

func Action(action string) string{
	return action
}