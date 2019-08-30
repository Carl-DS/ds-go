package main

type Talk interface {
	Hello(username string) string
	Talk(heard string) (saying string, end bool, err error)
}

type myTalk string

func (talk *myTalk) Hello(username string) string {
	return ""
}

func (talk *myTalk) Talk(heard string) (saying string, end bool, err error) {
	return "", false, nil
}

// 嵌入其他接口类型,相当于将其声明的方法集导入.这就要求不能有同名方法,因为不支持重载.
// 还有,不能嵌入自身或循环嵌入,那会导致递归错误
type Chatbot interface {
	Name() string
	Begin() (string, error)
	Talk
	ReportError(err error) string
	End() error
}

func main() {
	var talk Talk = new(myTalk)
	_, ok := talk.(*myTalk)
}
