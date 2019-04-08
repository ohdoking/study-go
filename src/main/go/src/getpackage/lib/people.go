package people

type people struct {
	name string
	age int
}

type peopleInfo people

func SetPeople(name string, age int) {
	peopleInfo = people {
		name : name,
		age : age
	}
}

func GetPeople() people{
	return peopleInfo
}