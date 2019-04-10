package people

type people struct {
	Name string
	Age int
}

var peopleInfo people

func SetPeople(name string, age int) {
	peopleInfo = people {Name: name, Age: age}
}

func GetPeople() people{
	return peopleInfo
}