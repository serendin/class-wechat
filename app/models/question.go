package models

import (
	"time"
	"github.com/revel/revel"
)

type Question struct {
	ID           int64     `json:"id"`
	Name         string    `json:"name"`
	LectureId    int       `json:"lecture_id"`
	StuId        string    `json:"stu_id"`
	Title        string    `json:"title"`
	Question     string    `json:"question"`
	CreatedAt    time.Time `json:"created_at"`
	Answer       string    `json:"answer"`
	Number       int       `json:"number"`
	UpdatedAt    time.Time `json:"updated_at"`

	LessonName string     `json:"lesson_name"`
	StuName    string     `json:"stu_name"`

}

const QuestionTable = "question"

func (Question) TableName() string {
	return QuestionTable
}

func GetQuestionList(page,size int,search string) (list []*Question) {
	search="%"+search+"%"
	err:=DB.Debug().Table(QuestionTable+" as n").Select("n.*,les.name as lesson_name,stu.name as stu_name").
		Joins("left join lecture as lec on lec.id=n.lecture_id").
		Joins("left join lesson as les on les.number=lec.lesson_no").
		Joins("left join student as stu on stu.stu_id=n.stu_id").
		Where("n.title like ? or les.name like ? or question like ?",search,search,search).
		Limit(size).Order("updated_at desc").Offset((page-1)*size).Scan(&list)
	if err!=nil{
		revel.WARN.Println("GetQuestionList Error:%v",err)
	}
	return
}
