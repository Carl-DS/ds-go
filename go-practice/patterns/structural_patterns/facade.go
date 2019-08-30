package main

import "fmt"

type Facade struct {
	M Music
	V Video
	C Count
}

func (this *Facade) GetRecommandVideos() error {
	this.V.GetVideo()
	this.C.GetCountByID(111)
	return nil
}

type Music struct{}
type Video struct{}
type Count struct {
	PraiseCnt  int64 // 点赞数
	CommentCnt int64 // 评论数
	CollectCnt int64 // 收藏数
}

func (this *Music) GetMusic() error {
	fmt.Println("get music material")
	// logic code here
	return nil
}

func (this *Video) GetVideo() error {
	fmt.Println("get videos")
	return nil
}

func (this *Count) GetCountByID(id int64) (*Count, error) {
	fmt.Println("get video counts")
	return this, nil
}

/**
外观模式
它为一套复杂的调度子系统提供一个统一的接入接口。外部所有对子系统的调用都通过这个外观角色进行统一调用，降低子系统与调用者之间的耦合度。
*/
func main() {
	f := &Facade{}
	f.GetRecommandVideos()
}
