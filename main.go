package main

import (
	"github.com/Go-ELK-API/requests/Elk_requests"
	"log"
	"os"
	"runtime/pprof"
	"strings"
	"time"
)

func main() {
	pprof.StartCPUProfile(os.Stdout)
	defer pprof.StopCPUProfile()

	var ServicesUnit []Elk_requests.ServiceUnit //массив для требуемых объектов
	var Services []Elk_requests.ServiceInfo                  //массив для 1860 объектов
	Services = Elk_requests.Mining1860_2(Services, "20190302", "20200308")
	layout := "20060102" //требуется для создания объекта time
	//итерируемся по массиву данных, полученных из эластика
	for _, service := range Services {

		qqc1860 := []rune(service.AppointmentId)
		qqc186 := qqc1860[:19] //обрезаем 1860 до нужной длины
		qqc153 := qqc1860[:7]
		var replacer = strings.NewReplacer("\\", "\\\\") // заменяем \ на \\ для запроса к эластику
		qqc153ForRequest := replacer.Replace(string(qqc153))
		qqc186ForRequest := replacer.Replace(string(qqc186))
		message153 := Elk_requests.DataFromQqc153By1860(qqc153ForRequest)
		message186 := Elk_requests.DataFromQqc186ByQqc1860(qqc186ForRequest)
		message83 := Elk_requests.DataFromQqc83By1860(service.OKMUCode)

		var serviceUnit Elk_requests.ServiceUnit
		input := message153.Hits.Hits[0].Source.DateOfBirth
		t, _ := time.Parse(layout, input)

		for _, item := range message83.Hits.Hits {
			if item.Source.Subcategory == "" || item.Source.Category == "" {
				continue
			} else {
				serviceUnit.Category = item.Source.Category
				serviceUnit.Subcategory = item.Source.Subcategory
			}
		}
		serviceUnit.Age = Elk_requests.AgeCalculate(t, time.Now())
		serviceUnit.AppointmentId = service.AppointmentId
		serviceUnit.Specialist = service.Specialist
		serviceUnit.Service = service.Service
		serviceUnit.AppointmentType = service.AppointmentType
		serviceUnit.OKMUCode = service.OKMUCode
		serviceUnit.Department = service.Department
		serviceUnit.Date = message186.Hits.Hits[0].Source.DateOfCompletion
		serviceUnit.DestinationStatus = message186.Hits.Hits[0].Source.DestinationStatus
		serviceUnit.PositionResource = message186.Hits.Hits[0].Source.PositionResource
		serviceUnit.SpecialistId = string(qqc1860[:19])
		serviceUnit.PatientId = message153.Hits.Hits[0].Source.PatientId
		serviceUnit.Gender = message153.Hits.Hits[0].Source.Gender
		serviceUnit.RegId = message153.Hits.Hits[0].Source.RegId
		serviceUnit.ServiceId = string(qqc1860)

		ServicesUnit = append(ServicesUnit, serviceUnit)

	}
	log.Println("Number of objects:", len(ServicesUnit))
}
