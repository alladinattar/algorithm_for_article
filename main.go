package main

import (
	"encoding/json"
	"fmt"
	"github.com/algorithm_for_article/requests/Elk_requests"
	"log"
	"strings"
	"sync"
	"time"
)

func arrayConv(array []Elk_requests.ServiceInfo) [][]Elk_requests.ServiceInfo {
	length := len(array)
	finalArray := [][]Elk_requests.ServiceInfo{}
	if length > 1000 {
		res1 := arrayConv(array[:length/2])
		res2 := arrayConv(array[length/2:])
		for i := 0; i < len(res1); i++ {
			finalArray = append(finalArray, res1[i])
		}
		for i := 0; i < len(res2); i++ {
			finalArray = append(finalArray, res2[i])
		}
	} else {

		finalArray = append(finalArray, array)
	}
	return finalArray
}

var replacer = strings.NewReplacer("\\", "\\\\")

var finishArray = struct {
	sync.Mutex
	result []Elk_requests.ServiceUnit
}{result: make([]Elk_requests.ServiceUnit, 0)}

func main() {
	/*pprof.StartCPUProfile(os.Stdout)
	defer pprof.StopCPUProfile()*/
	var wg sync.WaitGroup
	var Services []Elk_requests.ServiceInfo //массив для 1860 объектов
	Services = Elk_requests.Mining1860(Services, "20190302", "20200308")
	separateServices := arrayConv(Services)
	layout := "20060102" //требуется для создания объекта time
	//итерируемся по массиву данных, полученных из эластика
	for i := 0; i < len(separateServices); i++ {
		wg.Add(1)
		go func(array []Elk_requests.ServiceInfo) {
			requestMap := make(map[string]*Elk_requests.ServiceUnit)
			defer wg.Done()
			var message153Map = make(map[string]*Elk_requests.PatientInfo)
			var message83Map = make(map[string]*Elk_requests.CategorySubcategory)
			var message186Map = make(map[string]*Elk_requests.SpecialistInfo)
			var qqc153ArrTmp []string
			var qqc83ArrTmp []string
			var qqc186ArrTmp []string
			var qqc153ForRequest string
			var qqc83ForRequest string
			var qqc186ForRequest string
			for j := 0; j < len(array); j++ {
				qqc1860 := []rune(array[j].AppointmentId)
				qqc186 := qqc1860[:19] //обрезаем 1860 до нужной длины
				qqc153 := qqc1860[:7]
				requestMap[string(qqc1860)] = &Elk_requests.ServiceUnit{PatientId: string(qqc153), ServiceId: string(qqc1860), Specialist: array[j].Specialist, Service: array[j].Service,
					AppointmentType: array[j].AppointmentType, OKMUCode: array[j].OKMUCode, Department: array[j].Department, AppointmentId: array[j].AppointmentId, SpecialistId: string(qqc186)}

				var replacer = strings.NewReplacer("\\", "\\\\") // заменяем \ на \\ для запроса к эластику
				qqc83ForRequest = replacer.Replace(array[j].OKMUCode)
				qqc153ForRequest = replacer.Replace(string(qqc153))
				qqc186ForRequest = replacer.Replace(string(qqc186))
				qqc83ArrTmp = append(qqc83ArrTmp, `{"term": {"Du.keyword": "`+qqc83ForRequest+`"}}`)
				qqc153ArrTmp = append(qqc153ArrTmp, `{"term": {"qqc.keyword": "`+qqc153ForRequest+`"}}`)
				qqc186ArrTmp = append(qqc186ArrTmp, `{"term": {"qqc.keyword": "`+qqc186ForRequest+`"}}`)
				if j == len(array)-1 {
					concat153 := strings.Join(qqc153ArrTmp, ",")
					replacer = strings.NewReplacer("####", concat153)
					qqc153ForRequest = replacer.Replace(Elk_requests.MapReturn("get153"))
					concat186 := strings.Join(qqc186ArrTmp, ",")
					replacer = strings.NewReplacer("####", concat186)
					qqc186ForRequest = replacer.Replace(Elk_requests.MapReturn("get186"))
					concat83 := strings.Join(qqc83ArrTmp, ",")
					replacer = strings.NewReplacer("####", concat83)
					qqc83ForRequest = replacer.Replace(Elk_requests.MapReturn("get83"))
				}

			}
			var message153 Elk_requests.Message153
			var message186 Elk_requests.Message186
			var message83 Elk_requests.Message83


			Elk_requests.DataFromQqc153By1860(qqc153ForRequest, &message153)

			Elk_requests.DataFromQqc186By1860(qqc186ForRequest, &message186)

			Elk_requests.DataFromQqc83By1860(qqc83ForRequest, &message83)


			message153Arr := message153.Hits.Hits
			message186Arr := message186.Hits.Hits
			message83Arr := message83.Hits.Hits
			for _, object := range message186Arr {
				var specialistInfo Elk_requests.SpecialistInfo
				specialistInfo.SpecialistId = object.Source.SpecialistId
				specialistInfo.DestinationStatus = object.Source.DestinationStatus
				specialistInfo.DateOfCompletion = object.Source.DateOfCompletion
				specialistInfo.PositionResource = object.Source.PositionResource
				message186Map[object.Source.SpecialistId] = &specialistInfo
			}
			for _, object := range message153Arr {
				var patientInfo Elk_requests.PatientInfo
				patientInfo.PatientId = object.Source.PatientId
				patientInfo.RegId = object.Source.RegId
				patientInfo.Gender = object.Source.Gender
				patientInfo.DateOfBirth = object.Source.DateOfBirth
				message153Map[object.Source.PatientId] = &patientInfo
			}
			for _, object := range message83Arr {
				var categorySubcategory Elk_requests.CategorySubcategory
				categorySubcategory.Category = object.Source.Category
				categorySubcategory.Subcategory = object.Source.Subcategory
				categorySubcategory.OKMUCode = object.Source.OKMUCode
				message83Map[object.Source.OKMUCode] = &categorySubcategory
			}

			for _, service := range requestMap {
				input := message153Map[service.PatientId].DateOfBirth
				t, _ := time.Parse(layout, input)
				service.Gender = message153Map[service.PatientId].Gender
				service.RegId = message153Map[service.PatientId].RegId
				service.Age = Elk_requests.AgeCalculate(t, time.Now())
				service.PositionResource = message186Map[service.SpecialistId].PositionResource
				service.Date = message186Map[service.SpecialistId].DateOfCompletion
				service.DestinationStatus = message186Map[service.SpecialistId].DestinationStatus
				service.Category = message83Map[service.OKMUCode].Category
				service.Subcategory = message83Map[service.OKMUCode].Subcategory

				/*resp, _ := json.MarshalIndent(service, "", " ")
				fmt.Println(string(resp))*/
			}
			var resultPart []Elk_requests.ServiceUnit
			for _, value := range requestMap {
				resultPart = append(resultPart, *value)
			}

			finishArray.Lock()
			finishArray.result = append(finishArray.result, resultPart...)
			finishArray.Unlock()
		}(separateServices[i])

	}
	wg.Wait()

	for _,item:=range finishArray.result{
		resp,_:=json.MarshalIndent(item,"","  ")
		fmt.Println(string(resp))
	}

	log.Println("Number of result units", len(finishArray.result))
	finishArray.result = nil
}
