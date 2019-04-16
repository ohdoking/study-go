package common

type CommonInterface interface{
	Callback(a,b int) int
}

func Init(c int, callback CommonInterface) int {
	return c + callback.Callback(1,2)
}