package data_mining

import (
	".."
	"../w-parser"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Message293 struct {
	Aggs struct {
		Aggs_name struct {
			Buckets  [] struct {
				Key struct  {
					Pans    string `json:"pans"`
				} `json:"key"`
				Count    int `json:"doc_count"`
			} `json:"buckets"`
		} `json:"agg_name"`
	} `json:"aggregations"`
}



type Mapping map[string]int

type DataStruct struct {
	Status string
	Count int
	Mapping294 map[string]int
	Mapping293 map[string]int
	W294 map[string]int
	W293 map[string]int
}

type DataStructArray [] DataStruct


var dataStructArray = DataStructArray{}
var resultStructArray = DataStructArray{}
var replacer = strings.NewReplacer( "\\", "\\\\")

//
// Основная функция
//
func DataMining() {

	getAllStatus293(getMapping294(), getMapping293()) // получаем мапинги объектов и собираем существующие статусы
	for i:=0; i<len(dataStructArray); i++ {
		getQQCFrom293ByStatus(dataStructArray[i])
	}
	file, _ := os.Create("./struct.txt")
	strObj, _ := json.Marshal(resultStructArray)
	file.WriteString(string(strObj))
	defer file.Close()

}

//
// Cначала получаем маппинги 294 и 293 объектов
//
func getMapping294() Mapping {
	type mapping294 struct {
		Obj struct {
			Mappings struct {
				Properties  map[string]interface{} `json:"properties"`
			} `json:"mappings"`
		} `json:"294"`
	}
	res := requests.GetData("/294/_mapping", ``)
	data := mapping294{}
	_ = json.Unmarshal([]byte(res), &data)
	rdata := Mapping{}
	for k, _ := range data.Obj.Mappings.Properties{
		rdata[k] = 0
	}
	return rdata
}
//
// Получаем маппинг 293 объекта
//
func getMapping293() Mapping {
	type mapping294 struct {
		Obj struct {
			Mappings struct {
				Properties  map[string]interface{} `json:"properties"`
			} `json:"mappings"`
		} `json:"293"`
	}
	res := requests.GetData("/293/_mapping", ``)
	data := mapping294{}
	_ = json.Unmarshal([]byte(res), &data)
	rdata := Mapping{}
	for k, _ := range data.Obj.Mappings.Properties{
		rdata[k] = 0
	}
	return rdata
}

//
// Собираем существующие статусы
//
func getAllStatus293(m294 Mapping, m293 Mapping) {
	res := requests.GetData("/293/_search", `
{
  "size": 0,
  "aggs": {
    "agg_name": {
      "composite": {
        "size": 1000, 
        "sources": [
          {
            "pans": {
              "terms": {
                "field": "pAN.keyword"
              }
            }
          }
        ]
      }
    }
  }
}
					`)

	data := Message293{}
	_ = json.Unmarshal([]byte(res), &data)

	//log.Print("Кол-во объектов 293 " ,res)
	for  i := 0; i < len(data.Aggs.Aggs_name.Buckets); i++ {
		fmt.Println(data.Aggs.Aggs_name.Buckets[i].Key.Pans, data.Aggs.Aggs_name.Buckets[i].Count)
		_M293 := map[string]int{}
		for key, _ := range m293{
			_M293[key] = 0
		}
		_M294 := map[string]int{}
		for key, _ := range m294{
			_M294[key] = 0
		}
		_W293 := map[string]int{}
		_W294 := map[string]int{}
		ds := DataStruct{
			Status:     data.Aggs.Aggs_name.Buckets[i].Key.Pans,
			Count:      data.Aggs.Aggs_name.Buckets[i].Count,
			Mapping293: _M293,
			Mapping294: _M294,
			W293: _W293,
			W294: _W294,
		}
		dataStructArray = append(dataStructArray, ds )
	}
}
//
// Собираем маппинг 293 объекта + qqc для выгрузки 294
// делаем это в два этапа
// первый - обычная выгрузка с сортировкой по ключу
// второй - выгрузка со смещением
func getQQCFrom293ByStatus(obj DataStruct )  {
	type _qqc293 struct {
		Hits struct {
			Hits  [] struct {
				Source map[string]string `json:"_source"`
				Id string `json:"_id"`
			} `json:"hits"`
		} `json:"hits"`
	}

	qqcArray := [] string{}
	res := requests.GetData("/293/_search", `
	{
  "size":10000,
  "query": {
		"bool": {
			"filter":
			{
				"bool": {
				"must": [{"term": {"pAN.keyword": "`+obj.Status+`"}}]
      	}
	}
	}

	},
		"sort": [
        {"_id": "asc"}]
}`)

	//выгружаем все 293, которые соответствуют указанному статусу
	data := _qqc293{}
	_ = json.Unmarshal([]byte(res), &data)
	lastID := ""
	for i := 0; i < len(data.Hits.Hits); i++ {
		//Получаем маппинг 293 объекта
		for k, _ := range data.Hits.Hits[i].Source{
			obj.Mapping293[k]++
		}
		if len(data.Hits.Hits[i].Source["W"])>0 {
			for _, s := range w_parser.ParseStrWithoutText(data.Hits.Hits[i].Source["W"]){
				obj.W293[s]++
			}
		}
		//Получаем id пациента и текущий id медзаписи этого пациента
		qqcArray = append(qqcArray, string([] rune(data.Hits.Hits[i].Source["idStatusLong"])))
		lastID = string([] rune(data.Hits.Hits[i].Id))
	}

//
// конец 1 этапа, начало второго
//
	for  i := 10000; i < obj.Count;i = i + 10000 {
		res = requests.GetData("/293/_search", `
{
  "size":10000,
  "query": {
		"bool": {
			"filter":
			{
				"bool": {
				"must": [{"term": {"pAN.keyword": "`+obj.Status+`"}}]
      	}
	}
	}
	},
		"search_after": ["`+lastID+`"],
		"sort": [
        {"_id": "asc"}]
}`)
		_ = json.Unmarshal([]byte(res), &data)
		for i := 0; i < len(data.Hits.Hits); i++ {
			//Получаем маппинг 293 объекта
			for k, _ := range data.Hits.Hits[i].Source{
				obj.Mapping293[k]++
			}
			if len(data.Hits.Hits[i].Source["W"])>0 {
				for _, s := range w_parser.ParseStrWithoutText(data.Hits.Hits[i].Source["W"]){
					obj.W293[s]++
				}
			}
			//Получаем id пациента и текущий id медзаписи этого пациента
			qqcArray = append(qqcArray, string([] rune(data.Hits.Hits[i].Source["idStatusLong"])))
			lastID = string([] rune(data.Hits.Hits[i].Id))
		}


	}
	//запускаем выгрузку 294 объектов
	countFieldsof294ByStatus(obj, qqcArray)

}

//
// выгружаем 294 объект и считаем его поля
//
func countFieldsof294ByStatus(obj DataStruct, qqcArray []string){
	type _294 struct {
		Hits struct {
			Hits  [] struct {
				Source map[string]string `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}


	array := Delimeter(qqcArray)
	for i:=0;i<len(array);i++ {
		res := requests.GetData("/294/_search", `
		{
  		"size":10000,
  		"query": {
		"bool": {
			"filter":
			{
				"bool": {
					"should": [`+array[i]+`]
      				}
				}
			}
		}
		}`)

		//Проверка отправляемого запроса
		//fmt.Print(`
		//{
  		//"size":10000,
  		//"query": {
		//"bool": {
		//	"filter":
		//	{
		//		"bool": {
		//			"should": [`+array[i]+`]
      	//			}
		//		}
		//	}
		//}
		//}`)
		data := _294{}
		_ = json.Unmarshal([]byte(res), &data)
		for j:=0; j<len(data.Hits.Hits); j++{
			for k, _ := range data.Hits.Hits[j].Source{
				obj.Mapping294[k]++
			}
			if len(data.Hits.Hits[j].Source["W"])>0 {
				for _, s := range w_parser.ParseStrWithoutText(data.Hits.Hits[j].Source["W"]){
					obj.W294[s]++
				}
			}

		}



	}

	//печатаем количество элементов в статусе, проверить!!!
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("Объект", obj.Status)
	fmt.Println("Кол-во записей в 293:",obj.Count)
	fmt.Println("~~~~~~~~~~~~~~~~~~~")
	fmt.Println("~~~~~~~~~~~~~~~~~~~")
	fmt.Println("Структура в 294:")
	for k,v := range obj.Mapping294{
		if v>0 {
			fmt.Println(k,v)
		} else {
			delete(obj.Mapping294, k)
		}

	}
	fmt.Println("~~~~~~~~~~~~~~~~~~~")
	fmt.Println("~~~~~~~~~~~~~~~~~~~")
	fmt.Println("Структура в W294:")
	for k,v := range obj.W294{
			fmt.Println(k,v)
	}
	fmt.Println("~~~~~~~~~~~~~~~~~~~")
	fmt.Println("~~~~~~~~~~~~~~~~~~~")
	fmt.Println("Структура в 293:")
	for k,v := range obj.Mapping293{
		if v>0 {
			fmt.Println(k,v)
		} else {
			delete(obj.Mapping293, k)
		}
	}

	fmt.Println("~~~~~~~~~~~~~~~~~~~")
	fmt.Println("~~~~~~~~~~~~~~~~~~~")
	fmt.Println("Структура в W293:")
	for k,v := range obj.W293{
		fmt.Println(k,v)
	}
	resultStructArray = append(resultStructArray, obj)

}

func Delimeter(array [] string) ([] string) {
	x := len(array)
	destArray := [] string{}
	if (x>1000){
		res1:=Delimeter(array[:(x/2)])
		res2:=Delimeter(array[(x/2):])
		for i:=0; i<len(res1);i++ {
			destArray = append(destArray, res1[i])
		}
		for i:=0; i<len(res2);i++ {
			destArray = append(destArray, res2[i])
		}

	} else {
		tmp := []string{}
		for i:=0; i<len(array); i++{
			tmp = append(tmp, `{"term": {"idStatusLong.keyword": "`+array[i]+`"}}`)
		}

		concat := strings.Join(tmp, ",")
		concat = replacer.Replace(concat)
		destArray = append(destArray,concat)
	}
	return destArray
}