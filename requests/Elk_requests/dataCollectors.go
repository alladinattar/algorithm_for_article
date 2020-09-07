package Elk_requests

import (
	"encoding/json"
	"github.com/Go-ELK-API/requests"
	"log"
	"strings"
)


func Mining1860_2(array []ServiceInfo, startDate string, finishDate string) []ServiceInfo {
	replacer := strings.NewReplacer("####", startDate)
	body := replacer.Replace(mapReturn("mining1860"))
	replacer = strings.NewReplacer("****", finishDate)
	body = replacer.Replace(body)

	res := requests.GetData("/1860/_search?scroll=1m", body)
	data := Message1860{}
	err := json.Unmarshal([]byte(res), &data)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(data.Hits.Hits); i++ {
		var service ServiceInfo
		service.AppointmentType = data.Hits.Hits[i].Source.AppointmentType
		service.OKMUCode = data.Hits.Hits[i].Source.OKMUCode
		service.Date = data.Hits.Hits[i].Source.Date
		service.Department = data.Hits.Hits[i].Source.Department
		service.AppointmentId = data.Hits.Hits[i].Source.AppointmentId
		service.Service = data.Hits.Hits[i].Source.Service
		service.Specialist = data.Hits.Hits[i].Source.Specialist
		array = append(array, service)
	}
	scrollId := data.ScrollId

	for len(data.Hits.Hits) != 0 {
		scrollBody := `{
			"scroll":"1m",
				"scroll_id":"` + scrollId + `"
		}`
		res := requests.GetData("/_search/scroll", scrollBody)
		err := json.Unmarshal([]byte(res), &data)
		if err != nil {
			log.Fatal(err)
		}
		scrollId = data.ScrollId
		for i := 0; i < len(data.Hits.Hits); i++ {
			var service ServiceInfo
			service.AppointmentType = data.Hits.Hits[i].Source.AppointmentType
			service.OKMUCode = data.Hits.Hits[i].Source.OKMUCode
			service.Date = data.Hits.Hits[i].Source.Date
			service.Department = data.Hits.Hits[i].Source.Department
			service.AppointmentId = data.Hits.Hits[i].Source.AppointmentId
			service.Service = data.Hits.Hits[i].Source.Service
			service.Specialist = data.Hits.Hits[i].Source.Specialist
			array = append(array, service)
		}
	}
	return array
}
func Mining1860(array []ServiceInfo) []ServiceInfo {
	res := requests.GetData("/1860/_search", mapReturn("mining1860"))
	data := Message1860{}
	err := json.Unmarshal([]byte(res), &data)
	if err != nil {
		log.Fatal(err)
	}
	// Перебираем ответ эластика
	for i := 0; i < len(data.Hits.Hits); i++ {
		var service ServiceInfo
		service.AppointmentType = data.Hits.Hits[i].Source.AppointmentType
		service.OKMUCode = data.Hits.Hits[i].Source.OKMUCode
		service.Date = data.Hits.Hits[i].Source.Date
		service.Department = data.Hits.Hits[i].Source.Department
		service.AppointmentId = data.Hits.Hits[i].Source.AppointmentId
		service.Service = data.Hits.Hits[i].Source.Service
		service.Specialist = data.Hits.Hits[i].Source.Specialist
		array = append(array, service)
	}
	return array
}

func Mining174(array []DischargeInfo) []DischargeInfo {
	res := requests.GetData("/174/_search", mapReturn("mining174"))
	data := Message174{}
	err := json.Unmarshal([]byte(res), &data)
	if err != nil {
		log.Fatal(err)
	}
	// Перебираем ответ эластика
	for i := 0; i < len(data.Hits.Hits); i++ {
		var dischargeInfo DischargeInfo
		dischargeInfo.DischargeDiagnosis = data.Hits.Hits[i].Source.DischargeDiagnosis
		dischargeInfo.Department = data.Hits.Hits[i].Source.Department
		dischargeInfo.Specialist = data.Hits.Hits[i].Source.Specialist
		dischargeInfo.AdmissionType = data.Hits.Hits[i].Source.AdmissionType
		dischargeInfo.DischargeDate = data.Hits.Hits[i].Source.DischargeDate
		dischargeInfo.SpecialistPosition = data.Hits.Hits[i].Source.DischargeDate
		dischargeInfo.DischargeId = data.Hits.Hits[i].Source.DischargeId
		array = append(array, dischargeInfo)
	}
	return array
}

func DataFromQqc186ByQqc1860(AppointmentId string) Message186 {
	replacer := strings.NewReplacer("####", AppointmentId)
	body := replacer.Replace(mapReturn("get186"))
	res := requests.GetData("/186/_search", body)
	data := Message186{}
	err := json.Unmarshal([]byte(res), &data)
	if err != nil {
		log.Fatal(err)
	}
	// Перебираем ответ эластика
	return data
}

func DataFromQqc153By1860(AppointmentId string) Message153 {
	replacer := strings.NewReplacer("####", AppointmentId)
	body := replacer.Replace(mapReturn("get153"))
	res := requests.GetData("/153/_search", body)
	data := Message153{}
	err := json.Unmarshal([]byte(res), &data)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func DataFromQqc83By1860(OKMUCode string) Message83{
	replacer := strings.NewReplacer("####", OKMUCode)
	body := replacer.Replace(mapReturn("get83"))
	res := requests.GetData("/83/_search", body)
	data := Message83{}
	err := json.Unmarshal([]byte(res), &data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}
