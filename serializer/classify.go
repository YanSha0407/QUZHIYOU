package serializer

import "QUZHIYOU/models"

type Classify struct {
	ID int64 `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}


//单行序列化
func BuildClassify(item models.Classify) *Classify {

	return &Classify{
		ID:   item.ID,
		Name: item.Name,
		Icon: item.Icon,
	}
}


//多行序列化
func BuildClassifys(item []models.Classify) (classify []*Classify) {


	for _,v:=range item{
		classify = append(classify, BuildClassify(v))
	}

	return

}