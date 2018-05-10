package models

import (
	"time"
	"github.com/revel/revel"
)

type Video struct {
	ID           int       `json:"id"`
	AlbumID      int       `json:"album_id"`
	UserOpenID   string    `json:"user_open_id"`// 所属用户
	PetNumber    string    `json:"pet_number"`  // 所属宠物
	Video        string    `json:"video"`
	CreatedAt    time.Time `json:"created_at"`  // 添加时间
	Poster       string    `json:"poster"`      // 第一帧图片地址
}

const VideoTable = "album_videos"

func GetVideoList(page,size int,name string) (list []*Video) {
	_db:=DB.Where("name like '%?%'",name)
	_db=_db.Limit(size).Offset((page-1)*size)
	err := _db.Find(&list).Error
	if err!=nil{
		revel.WARN.Println("GetVideoList Error:%v",err)
	}
	return
}
