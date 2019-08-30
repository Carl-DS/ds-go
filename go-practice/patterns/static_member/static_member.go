package static_member

/**
这个模式只用在很稀少的场景，你需要某个struct类型的所有的实例需要共享同一个值
*/
type myImpl struct{}

/**
注意Kind()需要一个指针类型的receiver，但是它并没有为这个receiver命名，所以很清晰的表明我们并不使用类型实例。
*/
func (*myImpl) Kind() string {
	return "Best implementation"
}
