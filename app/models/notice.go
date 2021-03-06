package models

import (
	"time"
	"github.com/revel/revel"
)

type Notice struct {
	Id         int64      `json:"id"`
	Title      string     `json:"title"`
	LectureId  string     `json:"lecture_id"`
	Content    string     `json:"content"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	
	LessonName string     `json:"lesson_name"`
	TeaName    string     `json:"tea_name"`

}

const NoticeTable = "notice" 

func (Notice) TableName() string {
	return NoticeTable
}

func GetNoticeList(page,size int,search string) []Notice {
	var list []Notice
	search="%"+search+"%"
	err:=DB.Debug().Table(NoticeTable+" as n").Select("n.*,les.name as lesson_name,tea.name as tea_name").
		Joins("left join lecture as lec on lec.id=n.lecture_id").
			Joins("left join lesson as les on les.number=lec.lesson_no").
				Joins("left join teacher as tea on tea.number=lec.teacher_no").
			Where("title like ? or les.name like ?",search,search).
				Limit(size).Order("updated_at desc").Offset((page-1)*size).Scan(&list)
	if err!=nil{
		revel.WARN.Println("GetNoticeList Error:%v",err)
	}
	return list
}