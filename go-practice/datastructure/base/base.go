package base

// just list Object in java
type Object interface{}

// an function type for match data
type MatchFun func(data1, data2 Object) int

type MyStr struct {
	name string
}

func match(data1 Object, data2 Object) int {
	myStr1 := data1.(*MyStr)
	myStr2 := data2.(*MyStr)

	if (*myStr1).name == (*myStr2).name {
		return 0
	} else {
		return 1
	}
}

type Iterator interface {
	HasNext() bool
	Next() Object
}
