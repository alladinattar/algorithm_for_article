package Elk_requests


/*import (
	"encoding/json"
	"fmt"
	"strings"
)

func DischargeMining() {
	var dischargeUnits []DischargeUnit   //массив для требуемых объектов
	var dischargePatient []DischargeInfo //массив для 174 объектов
	dischargePatient = Mining174(dischargePatient)
	//итерируемся по массиву данных, полученных из эластика
	for index, item := range dischargePatient {
		fmt.Println("Object number:", index)

		qqc174 := []rune(item.DischargeId)
		qqc153 := qqc174[:7] //обрезаем qqc174 до нужной длины

		var replacer = strings.NewReplacer("\\", "\\\\")
		qqc153ForRequest := replacer.Replace(string(qqc153)) // заменяем \ на \\ для запроса к эластику
		message153 := DataFromQqc153By1860(qqc153ForRequest)

		var dischargeUnit DischargeUnit
		dischargeUnit.DischargeDate = item.DischargeDate
		dischargeUnit.SpecialistPosition = item.SpecialistPosition
		dischargeUnit.DischargeDate = item.DischargeDate
		dischargeUnit.AdmissionType = item.AdmissionType
		dischargeUnit.Department = item.Department
		dischargeUnit.DischargeDiagnosis = item.DischargeDiagnosis
		dischargeUnit.Specialist = item.Specialist
		dischargeUnit.Gender = message153.Hits.Hits[0].Source.Gender
		dischargeUnit.PatientId = message153.Hits.Hits[0].Source.PatientId
		dischargeUnits = append(dischargeUnits, dischargeUnit)
		resp, _ := json.MarshalIndent(dischargeUnit, "", "  ")
		fmt.Println(string(resp))

	}
}
*/