package model

// 定义一个结构体
// type Student struct {
// 	Name string
// 	Score float64
// }

type student struct {
	Name string
	Score float64
}
// 2. 如果student结构体首字母是小写,因此是只能在model中使用
// 3. 我们通过工厂模式来解决
func NewStudent(n string, s float64) *student {
	return &student{
		Name: n,
		Score: s,
	}
}
