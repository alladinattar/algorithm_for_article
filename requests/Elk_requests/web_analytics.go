package Elk_requests

import (
	"encoding/json"
	"github.com/Go-ELK-API/requests"
	"log"
)

func GetConsistedPatients(array []ConsistedPatient) []ConsistedPatient {
	res := requests.GetData("/174/_search", mapReturn("consistedPatients"))
	data := Message174{}
	err := json.Unmarshal([]byte(res), &data)
	if err != nil {
		log.Fatal(err)
	}
	// Перебираем ответ эластика
	for i := 0; i < len(data.Hits.Hits); i++ {
		var patient ConsistedPatient
		patient.DischargeDiagnosis = data.Hits.Hits[i].Source.DischargeDiagnosis
		patient.Department = data.Hits.Hits[i].Source.Department
		patient.ReceiptDate = data.Hits.Hits[i].Source.ReceiptDate
		array = append(array, patient)
	}
	return array
}

func GetServicesByName(array []ServiceInfo, serviceName string) []ServiceInfo {
	res := requests.GetData("/1860/_search", mapReturn(serviceName))
	data := Message1860{}
	err := json.Unmarshal([]byte(res), &data)
	if err != nil {
		log.Fatal(err)
	}
	// Перебираем ответ эластика
	array = AssembleServiceFrom1860(data)
	return array
}

func AssembleServiceFrom1860(data Message1860) []ServiceInfo {
	var array []ServiceInfo
	for i := 0; i < len(data.Hits.Hits); i++ {
		var service ServiceInfo
		service.Date = data.Hits.Hits[i].Source.Date
		service.Department = data.Hits.Hits[i].Source.Department
		service.Specialist = data.Hits.Hits[i].Source.Specialist
		service.OKMUCode = data.Hits.Hits[i].Source.OKMUCode
		service.AppointmentId = data.Hits.Hits[i].Source.AppointmentId
		service.Service = data.Hits.Hits[i].Source.Service
		service.AppointmentType = data.Hits.Hits[i].Source.AppointmentType
		array = append(array, service)
	}
	return array
}
