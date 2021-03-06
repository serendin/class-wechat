package models

import (
	"time"
	"github.com/revel/revel"
)

type Question struct {
	ID           int64     `json:"id"`
	LectureId    int       `json:"lecture_id"`
	StuNo        string    `json:"stu_no"`
	Title        string    `json:"title"`
	Question     string    `json:"question"`
	CreatedAt    time.Time `json:"created_at"`
	Answer       string    `json:"answer"`
	Number       int       `json:"number"`
	UpdatedAt    time.Time `json:"updated_at"`
	Status       int       `gorm:"default:0"`

	LessonName string     `json:"lesson_name" gorm:"-"`
	StuName    string     `json:"stu_name" gorm:"-"`

}

const QuestionTable = "question"

func (Question) TableName() string {
	return QuestionTable
}

type Lesson struct {
	Value    int      `json:"value"`
	Title    string   `json:"title"`
}

func GetQuestionList(page,size int,search string,stu string) (list []*Question) {
	search="%"+search+"%"
	_db:=DB.Debug().Table(QuestionTable+" as n").Select("n.*,les.name as lesson_name,stu.name as stu_name").
		Joins("left join lecture as lec on lec.id=n.lecture_id").
		Joins("left join lesson as les on les.number=lec.lesson_no").
		Joins("left join student as stu on stu.stu_id=n.stu_no").
		Where("n.title like ? or les.name like ? or question like ?",search,search,search)
	if stu!=""{
		_db=_db.Where("stu.stu_id = ?",stu)
	}
	err:=_db.Limit(size).Order("updated_at desc").Offset((page-1)*size).Scan(&list)

	if err!=nil{
		revel.WARN.Println("GetQuestionList Error:%v",err)
	}
	return
}

func GetLessonList(stu string) (list []*Lesson) {
	err:=DB.Debug().Table("class as n").Select("n.lecture_id as value,les.name as title").
		Joins("left join lecture as lec on lec.id=n.lecture_id").
		Joins("left join lesson as les on les.number=lec.lesson_no").
		Where("n.stu_no = ?",stu).Scan(&list)

	if err!=nil{
		revel.WARN.Println("GetLessonList Error:%v",err)
	}
	return
}

func (q *Question) Save() error {
	err:= DB.Save(q).Error
	return err
}
