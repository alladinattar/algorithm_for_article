package sberbank

import (
	".."
	"encoding/json"
	"log"
	"strings"
	"sync"
	"time"
)

var dataSet = DataSet{}

var replacer = strings.NewReplacer( "\\", "\\\\")


//Структуры данных, считываемые из ответов elasticsearch
type Message294 struct {
	Hits struct {
		Hits  [] struct {
			Source struct  {
				Qqc    string `json:"qqc"`
				W      string `json:"W"`
				PDdate string `json:"pDdate"`
				PDtype string `json:"pDtype"`
			} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`

}

type Message293 struct {
	Hits struct {
		Hits  [] struct {
			Source struct  {
				Qqc    string `json:"qqc"`
				Status    string `json:"pAN"`
			} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}


type Message153 struct {
	Hits struct {
		Hits  [] struct {
			Source struct  {
				Qqc    string `json:"qqc"`
				Sex    string `json:"pJ"`
				Birthday    string `json:"pI"`
			} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}
// Структуры данных, для ответа на запрос

type DataSet [] Data

type Data struct {
	idPatient string
	Sex string
	Birthday string
	Data [] Obj
}

type Obj struct {
	Status string
	idEpisod string
	Description string
	Type string
	Date string
}

// Временные структуры данных для оптимизации скорости подбора данных
var dataDict  = struct {
	sync.RWMutex
	m map[string] DataTmp
}{m: make(map[string] DataTmp)}

type DataTmp struct {
	Qqc string
	idPatient string
	Sex string
	Birthday string
	Data map[string] [] ObjTmp
}

type ObjTmp struct {
	Status string
	QQCtemp string
	idEpisod string
	Description string
	Type string
	Date string
}


func Report() []byte{
	start_time :=time.Now()
	Get294bypDcodeRg()
	limitedQueries()
	response, _ := json.Marshal(dataDict.m)
	log.Print(time.Since(start_time))
	return response
}


func Get294bypDcodeRg() {
	res := requests.GetData("/294/_search", `
		{
             "size": 1000,
			 "query": {
    			"match": {
      				"W": "C73 Злокачественное новообразование щитовидной железы"
    					}
  					}
				}
	`)

	data := Message294{}
	_ = json.Unmarshal([]byte(res), &data)
	// Перебираем ответ эластика
	log.Print("Кол-во объектов 294 " ,len(data.Hits.Hits))
	for i := 0; i < len(data.Hits.Hits); i++ {
		idPatient := string([] rune(data.Hits.Hits[i].Source.Qqc)[:7])
		if val, ok := dataDict.m[idPatient]; ok {
			// если объект уже существует, дополняем его поля
			obj294 := ObjTmp{}
			obj294.idEpisod = string([] rune(data.Hits.Hits[i].Source.Qqc)[7:10])
			obj294.Date = data.Hits.Hits[i].Source.PDdate
			obj294.Type = data.Hits.Hits[i].Source.PDtype
			obj294.QQCtemp = string([] rune(data.Hits.Hits[i].Source.Qqc)[:20])

			obj294.Description = data.Hits.Hits[i].Source.W
			//добавляем обратно обновленное значение
			if _, ok := val.Data[obj294.QQCtemp]; ok {
				val.Data[obj294.QQCtemp] = append(val.Data[obj294.QQCtemp], obj294)
				dataDict.m[idPatient] = val
			} else {
				tempDataArray := [] ObjTmp{}
				tempDataArray = append(tempDataArray, obj294)
				val.Data[obj294.QQCtemp] = tempDataArray
				dataDict.m[idPatient] = val
			}
		} else {
			// в случае, если объект новый
			tmp := DataTmp{}
			tmp.Qqc = data.Hits.Hits[i].Source.Qqc
			tmp.idPatient = idPatient

			// создаем описание и добавляем
			obj294 := ObjTmp{}
			obj294.idEpisod = string([] rune(data.Hits.Hits[i].Source.Qqc)[7:10])
			obj294.Date = data.Hits.Hits[i].Source.PDdate
			obj294.Type = data.Hits.Hits[i].Source.PDtype
			obj294.QQCtemp = string([] rune(data.Hits.Hits[i].Source.Qqc)[:20])

			obj294.Description = data.Hits.Hits[i].Source.W
			tmp.Data = make(map[string] []ObjTmp)
			tmp.Data[obj294.QQCtemp] = append(tmp.Data[obj294.QQCtemp],obj294)
			dataDict.m[idPatient] = tmp
		}
	}

	log.Print("patients ",len(dataDict.m))

	//печать информации о количестве объектов
	//for k := range dataDict {
	//	log.Printf("Пациент", dataDict[k].idPatient , len(dataDict[k].Data))
	//}

}

func GetStatusFrom293(concat string, wg *sync.WaitGroup) {

	res := requests.GetData("/293/_search", `
			{
             "size": 10000,
			 "query": {
			   "bool": {
				 "filter":
				 {
				   "bool" : {
           				"should" :[`+ concat + `]
					   }
					 }
				   }
				 }
				}
					`)

	data := Message293{}
	_ = json.Unmarshal([]byte(res), &data)
	// Перебираем ответ эластика
	log.Print("Кол-во объектов 293 " ,len(data.Hits.Hits))
	for i := 0; i < len(data.Hits.Hits); i++ {
		//Получаем id пациента и текущий id медзаписи этого пациента
		idPatient := string([] rune(data.Hits.Hits[i].Source.Qqc)[:7])
		currentQQC := data.Hits.Hits[i].Source.Qqc
		// Получаем пациента по его идентификатору
		dataDict.RLock()
		if val, ok := dataDict.m[idPatient]; ok {
			dataDict.RUnlock()
			// если объект уже существует, дополняем его поля
			value := val.Data[currentQQC]
			for j := range value {
				value[j].Status= data.Hits.Hits[i].Source.Status
			}
			//добавляем обратно обновленное значение
			dataDict.Lock()
			dataDict.m[idPatient].Data[currentQQC] = value
			dataDict.Unlock()
		} else {
			dataDict.RUnlock()
			// в случае, если объекта не существует
			continue
		}
	}
	//печать информации о количестве объектов
	//for k := range dataDict {
	//	log.Print("\n")
	//	log.Print("Пациент", dataDict[k].idPatient)
	//	log.Print("Пол", dataDict[k].Sex)
	//	log.Print("Дата Рождения", dataDict[k].Birthday)
	//	log.Print("Кол-во 293 объектов:", len(dataDict[k].Data))
	//	for i := range dataDict[k].Data {
	//		for j:= range dataDict[k].Data[i] {
	//			log.Print(dataDict[k].Data[i][j].Status)
	//		}
	//	}
	//}
	defer wg.Done()
}


func GetInfoFrom153(concat string, wg *sync.WaitGroup) {

	res := requests.GetData("/153/_search", `
			{
  			 "size": 10000,
			 "query": {
			   "bool": {
				 "filter":
				 {
				   "bool" : {
           				"should" :[`+ concat + `]
					   }
					 }
				   }
				 }
				}
					`)

	data := Message153{}
	_ = json.Unmarshal([]byte(res), &data)
	// Перебираем ответ эластика
	log.Print("Кол-во объектов 153 " ,len(data.Hits.Hits))
	for i := 0; i < len(data.Hits.Hits); i++ {
		//Получаем id пациента и текущий id медзаписи этого пациента
		idPatient := data.Hits.Hits[i].Source.Qqc

		// Получаем пациента по его идентификатору
		dataDict.RLock()
		if val, ok := dataDict.m[idPatient]; ok {
			dataDict.RUnlock()
			// если объект уже существует, дополняем его поля
			val.Sex = data.Hits.Hits[i].Source.Sex
			val.Birthday = data.Hits.Hits[i].Source.Birthday
			//добавляем обратно обновленное значение
			dataDict.Lock()
			dataDict.m[idPatient] = val
			dataDict.Unlock()
		} else {
			// в случае, если объекта не существует
			dataDict.RUnlock()
			continue
		}
	}
	defer wg.Done()
}

func limitedQueries() {
	var wg sync.WaitGroup
	tmp293 := [] string{}
	dataDict.RLock()
	for k := range dataDict.m {
		for i := range dataDict.m[k].Data {
			for j := range dataDict.m[k].Data[i] {
				tmp293 = append(tmp293, `{"term": {"qqc.keyword": "`+dataDict.m[k].Data[i][j].QQCtemp+`"}}`)
			}
		}
	}
	dataDict.RUnlock()
	wg.Add(1)
	go Delimeter(tmp293, 293, &wg)

	tmp153 := [] string{}
	dataDict.RLock()
	for k := range dataDict.m {
		tmp153 = append(tmp153, `{"term": {"qqc.keyword": "`+dataDict.m[k].idPatient+`"}}`)
	}
	dataDict.RUnlock()
	wg.Add(1)
	go Delimeter(tmp153, 153, &wg)
	wg.Wait()
	log.Print("Complete")

}

func Delimeter(array [] string, flag int, wg *sync.WaitGroup) {
	x := len(array)
	if (x>1000){
		wg.Add(2)
		Delimeter(array[:(x/2)], flag, wg)
		Delimeter(array[(x/2):], flag, wg)
	} else {
		concat := strings.Join(array, ",")
		concat = replacer.Replace(concat)
		wg.Add(1)
		if (flag==153) {
			go GetInfoFrom153(concat, wg)
		} else if (flag==293) {
			go GetStatusFrom293(concat, wg)
		}
	}
	wg.Done()
}