package Elk_requests

/*import (
	"log"
	"strings"
	"time"
)

func ServiceUnitMining(startDate string, finishDate string) {
	var ServicesUnit []ServiceUnit //массив для требуемых объектов
	var Services []ServiceInfo     //массив для 1860 объектов
	Services = Mining1860(Services, startDate, finishDate)
	layout := "20060102" //требуется для создания объекта time
	//итерируемся по массиву данных, полученных из эластика
	for _, service := range Services {

		qqc1860 := []rune(service.AppointmentId)
		qqc186 := qqc1860[:19] //обрезаем 1860 до нужной длины
		qqc153 := qqc1860[:7]
		var replacer = strings.NewReplacer("\\", "\\\\") // заменяем \ на \\ для запроса к эластику
		qqc153ForRequest := replacer.Replace(string(qqc153))
		qqc186ForRequest := replacer.Replace(string(qqc186))
		message153 := DataFromQqc153By1860(qqc153ForRequest)
		message186 := DataFromQqc186ByQqc1860(qqc186ForRequest)
		message83 := DataFromQqc83By1860(service.OKMUCode)

		var serviceUnit ServiceUnit
		input := message153.Hits.Hits[0].Source.DateOfBirth
		t, _ := time.Parse(layout, input)

		for _,item:=range message83.Hits.Hits{
			if item.Source.Subcategory == "" || item.Source.Category == ""{
				continue
			}else{
				serviceUnit.Category = item.Source.Category
				serviceUnit.Subcategory = item.Source.Subcategory
			}
		}
		serviceUnit.Age = AgeCalculate(t, time.Now())
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
	log.Println("Number of objects:",len(ServicesUnit))
}
*/