
package uslugi

import (
	".."
	"encoding/json"
	"strings"
)

//var dataSet = DataSet{}

var replacer = strings.NewReplacer( "\\", "\\\\")


//Структуры данных, считываемые из ответов elasticsearch
type Message1860 struct {
	Hits struct {
		Hits  [] struct {
			Source struct  {
				Date  string `json:"dE"` //дата
				OKMUCode string `json:"Du"` //код ОКМУ
				ServiceId string `json:"qqc"` //qqc
				Service   string `json:"u"` //услуга
				Specialist string `json:"pAz"` //Специалист
				Department string `json:"pID"` //Отделение
				AppointmentType  string `json:"tn"` //Тип назначения
			} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

type Message153 struct {
	Hits struct {
		Hits  [] struct {
			Source struct  {
				PatientId string `json:"qqc"` //qqc
				RegId   string `json:"pB"` //услуга
			} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

//Финальная структура данных для хранения всей информации "вокруг" услуг
type Service struct {
	Date  		 	 string //дата
	OKMUCode 		 string //код ОКМУ
	ServiceId 		 string //qqc
	Service   		 string //услуга
	Specialist 	 	 string //Специалист
	Department 		 string //Отделение
	AppointmentType  string //Тип назначения
	RegId   		 string //Регистрационный номер пациента
}

//Массив содержащий финальные данные по услугам
var Services []Service

//"size": 10000,
func GetRegIdFrom153(id string) (string) {
	if strings.Contains(id, "\\"){
		id = SlashShielding(id) //экранируем слэш "\", чтобы эластик хорошо проглатывал
	}

	res := requests.GetData("/153/_search", `
{
 "_source": [
  "qqc",
  "pB",
  "pI"
   ],
  "query": {
    "term": {
      "qqc.keyword": "`+id+`"
    }
  }
}
`)
	data := Message153{}
	_ = json.Unmarshal([]byte(res), &data)
	// Перебираем ответ эластика
	//fmt.Println("Кол-во объектов 153: ", len(data.Hits.Hits))
	var regN string
	if len(data.Hits.Hits)>0 {
		regN = data.Hits.Hits[0].Source.RegId
	}
	//println("qqc,regN : ", qqc,regN)
	return regN
}

//символ слэша воспринимается эластиком как зарезерированный. Поэтому экранируем его дополнительным слэшом.
//работает с "\", осталось разобраться с "/"
func SlashShielding(str string) (string) {
	//slashCount := strings.Count(str,"\\")
	i := 0
	for byteIndex, runeValue := range str {
		//fmt.Printf("character %d starts at byte position %d\n", runeValue, byteIndex)
		if runeValue == 92 { // 92 - значение руны, соответствующее слэшу "/"
			i++
			str = str[:byteIndex+i] + "\\" + str[byteIndex+i:]
		}
	}
	return str
}

func GetServicesFrom1860() {

	res := requests.GetData("/1860/_search", `
			{
			"size": 10000, 
		   "_source": [
			"dE",
			"Du",
			"qqc",
		   	"u",
			"pAz",
		   	"pID",
		   	"tn"
		   ],  
		  "query": {
			   "bool": {
				 "must": [
				   {"match": {
					 "tn.keyword": "Консультация"
				   }}
				 ], 
				 "filter":[
				   {"range":{"dE": {"gte":"20200420"}}},
				   {"range":{"dE": {"lte":"20200521"}}}
				   ]
			   }
		  }
		
		}
					`)

	data := Message1860{}
	_ = json.Unmarshal([]byte(res), &data)
	// Перебираем ответ эластика
	//fmt.Println("Кол-во объектов 1860 ",len(data.Hits.Hits))

	for i := 0; i < len(data.Hits.Hits); i++ {
		var service Service
		service.AppointmentType = data.Hits.Hits[i].Source.AppointmentType
		service.OKMUCode = data.Hits.Hits[i].Source.OKMUCode
		service.Date = data.Hits.Hits[i].Source.Date
		service.Department = data.Hits.Hits[i].Source.Department
		service.ServiceId = data.Hits.Hits[i].Source.ServiceId
		service.Service = data.Hits.Hits[i].Source.Service
		service.Specialist = data.Hits.Hits[i].Source.Specialist
		service.RegId =GetRegIdFrom153(string([] rune(service.ServiceId)[:7]))

		//println("service : ",service.Specialist, service.Service, service.Department, service.Date,
		//	service.AppointmentType, service.OKMUCode, service.ServiceId, service.RegId)
		Services = append(Services, service)
	}

}



