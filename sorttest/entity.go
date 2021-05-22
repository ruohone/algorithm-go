package sorttest

type user struct {
	Id int
	Name string
	Age int
}

var testData = []user{
	{Id:1,Name:"Tony",Age:1},{Id:567,Name:"Jack",Age:45},{Id:456,Name:"Pony",Age:76},
	{Id:87,Name:"Jason",Age:65},{Id:34,Name:"ZhangSan",Age:4},
	{Id:43,Name:"Lily",Age:4},{Id:43,Name:"WangEr",Age:65},{Id:546,Name:"LiSi",Age:5},
}

func swap (data []user,i,j int)  {
	data[i],data[j] = data[j], data[i]
}