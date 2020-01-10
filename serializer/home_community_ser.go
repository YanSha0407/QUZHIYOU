package serializer

import (
	"QUZHIYOU/models"
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"sort"
)

//序列化1
type Community struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	KeyWord string `json:"key_word"`
	Letter  string `json:"letter"`
}

//返回结果序列化2
type ResComs struct {
	Letter string      `json:"letter"`
	Data   []Community `json:"data"`
}

func BuildCommunity(item models.Communitys) Community {
	return Community{
		Id:      item.ID,
		Name:    item.Name,
		KeyWord: item.KeyWord,
		Letter:  item.Letter,
	}
}

type Animals []models.Communitys

//noinspection ALL
func (a Animals) Len() int { return len(a) }
//noinspection ALL
func (s Animals ) Less(i, j int) bool {
	a, _ := UTF82GBK(s[i].Letter)
	b, _ := UTF82GBK(s[j].Letter)
	bLen := len(b)
	for idx, chr := range a {
		if idx > bLen-1 {
			return false
		}
		if chr != b[idx] {
			return chr < b[idx]
		}
	}
	return true
}
//noinspection ALL
func (a Animals ) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

//UTF82GBK : transform UTF8 rune into GBK byte array
func UTF82GBK(src string) ([]byte, error) {
	GB18030 := simplifiedchinese.All[0]
	return ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(src)), GB18030.NewEncoder()))
}

//GBK2UTF8 : transform  GBK byte array into UTF8 string
func GBK2UTF8(src []byte) (string, error) {
	GB18030 := simplifiedchinese.All[0]
	bytes, err := ioutil.ReadAll(transform.NewReader(bytes.NewReader(src), GB18030.NewDecoder()))
	return string(bytes), err
}

func BuildCommunitys(item []models.Communitys) []interface{} {



	//强转类型
	sort.Sort(Animals(item))



	var cityList []Community

	for _, v := range item {

		cityList = append(cityList, BuildCommunity(v))
	}



	var totalList []interface{}

	for i := 0; i < len(cityList); i++ {
		//取出第一个数组
		item := cityList[i]

		var tempList []Community
		tempList = append(tempList, item)

		for j := i + 1; j < len(cityList); j++ {
			temp := cityList[j]

			if temp.Letter == item.Letter {
				tempList = append(tempList, temp)
				cityList = append(cityList[:j], cityList[j+1:]...)
				j = j - 1
			}

		}

		dic := ResComs{}

		dic.Letter = item.Letter
		dic.Data = tempList

		//添加相同的元素ABC
		totalList = append(totalList, dic)

	}

	return totalList

}
