package models

import (
	"time"
	"github.com/revel/revel"
)

type Video struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Url          string    `json:"url"`
	Uploader     string    `json:"uploader"`
	Type         string    `json:"type"`
	Brief        string    `json:"brief"`
	CreatedAt    time.Time `json:"created_at"`
	Poster       string    `json:"poster"`
}

const VideoTable = "video"

func (Video) TableName() string {
	return VideoTable
}
func GetVideoList(page,size int,name string) (list []*Video) {
	err := DB.Debug().Where("name like ?","%"+name+"%").
		Limit(size).Order("created_at desc").Offset((page-1)*size).Find(&list).Error
	if err!=nil{
		revel.WARN.Println("GetVideoList Error:%v",err)
	}

	return
}

func GetVideoById(id int) Video {
	var video Video
	DB.First(&video,id)
	return video
}