package models

import (
	"time"
	"github.com/revel/revel"
)

type Material struct {
	ID           int64     `json:"id"`
	Name         string    `json:"name"`
	LectureId    int       `json:"lecture_id"`
	Url          string    `json:"url"`
	Type         string    `json:"type"`
	Brief        string    `json:"brief"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	FileName     string    `json:"file_name"`

	LessonName string     `json:"lesson_name"`
	TeaName    string     `json:"tea_name"`

}

const MaterialTable = "material"

func (Material) TableName() string {
	return MaterialTable
}

func GetMaterialList(page,size int,search string) (list []*Material) {
	search="%"+search+"%"
	err:=DB.Debug().Table(MaterialTable+" as n").Select("n.*,les.name as lesson_name,tea.name as tea_name").
		Joins("left join lecture as lec on lec.id=n.lecture_id").
		Joins("left join lesson as les on les.number=lec.lesson_no").
		Joins("left join teacher as tea on tea.number=lec.teacher_no").
		Where("n.name like ? or les.name like ? or brief like ?",search,search,search).
		Limit(size).Order("updated_at desc").Offset((page-1)*size).Scan(&list)
	if err!=nil{
		revel.WARN.Println("GetMaterialList Error:%v",err)
	}
	return
}
