package models

import (
	"github.com/revel/revel"
)

type Grade struct {
	Id         int64   `json:"id"`
	LectureId  int64   `json:"lecture_id"`
	StuNo      string  `json:"stu_no"`
	Score      float64 `json:"score"`

	LessonName string  `json:"lesson_name"`
	TeaName    string  `json:"tea_name"`
	Credit     int     `json:"credit"`
}

const ClassTable = "class"

func (Grade) TableName() string {
	return ClassTable
}

func GetGradeList(term,stuNo string) (list []*Grade) {
	err:=DB.Debug().Table(ClassTable+" as n").Select("n.*,les.name as lesson_name,les.credit as credit,tea.name as tea_name").
		Joins("left join lecture as lec on lec.id=n.lecture_id").
		Joins("left join lesson as les on les.number=lec.lesson_no").
		Joins("left join teacher as tea on tea.number=lec.teacher_no").
		Where("n.stu_no = ? and lec.term = ?",stuNo,term).Scan(&list)
	if err!=nil{
		revel.WARN.Println("GetGradeList Error:%v",err)
	}
	return
}
