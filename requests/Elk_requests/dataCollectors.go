package Elk_requests

import (
	"encoding/json"
	"github.com/algorithm_for_article/requests"
	"log"
	"strings"
)

func Mining1860(array []ServiceInfo, startDate string, finishDate string) []ServiceInfo {
	replacer := strings.NewReplacer("####", startDate)
	body := replacer.Replace(MapReturn("mining1860"))
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

func Mining174(array []DischargeInfo, startDate string, finishDate string) []DischargeInfo {
	replacer := strings.NewReplacer("####", startDate)
	body := replacer.Replace(MapReturn("mining174"))
	replacer = strings.NewReplacer("****", finishDate)
	body = replacer.Replace(body)
	res := requests.GetData("/174/_search", body)
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

func DataFromQqc186By1860(body string, object *Message186) {
	res := requests.GetData("/186/_search", body)
	data := Message186{}
	err := json.Unmarshal([]byte(res), &data)
	if err != nil {
		log.Fatal(err)
	}
	*object = data
	// Перебираем ответ эластика
}

func DataFromQqc153By1860(body string, object *Message153) {
	res := requests.GetData("/153/_search", body)
	data := Message153{}
	err := json.Unmarshal([]byte(res), &data)
	if err != nil {
		log.Fatal(err)
	}
	*object = data
}

func DataFromQqc83By1860(body string, object *Message83) {
	res := requests.GetData("/83/_search", body)
	data := Message83{}
	err := json.Unmarshal([]byte(res), &data)
	if err != nil {
		log.Fatal(err)
	}
	*object = data
}