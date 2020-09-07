package Elk_requests

type Message1860 struct {
	ScrollId string `json:"_scroll_id"`
	Hits     struct {
		Hits []struct {
			Source struct {
				Date            string `json:"dE"`  //дата
				OKMUCode        string `json:"Du"`  //код ОКМУ
				AppointmentId   string `json:"qqc"` //qqc
				Service         string `json:"u"`   //услуга
				Specialist      string `json:"pAz"` //Специалист
				Department      string `json:"pID"` //Отделение
				AppointmentType string `json:"tn"`  //Тип назначения
			} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

type Message153 struct {
	Hits struct {
		Hits []struct {
			Source struct {
				PatientId   string `json:"qqc"`
				RegId       string `json:"pB"`
				Gender      string `json:"pJ"`
				DateOfBirth string `json:"pI"`
			} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

type Message186 struct {
	Hits struct {
		Hits []struct {
			Source struct {
				SpecialistId      string `json:"qqc"`    //qqc специалиста
				DateOfCompletion  string `json:"pAE"`    //дата выполнения
				PositionResource  string `json:"puR"`    //Должность_Ресурс
				DestinationStatus string `json:"pANdop"` // Состояние_назначения
			} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

type Message83 struct {
	Hits struct {
		Hits []struct {
			Source struct {
				Category    string `json:"pu"`
				Subcategory string `json:"pu1"`
				OKMUCode    string `json:"Du"`
			} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

type Message174 struct {
	Hits struct {
		Hits []struct {
			Source struct {
				DischargeDiagnosis string `json:"pKDiag"`
				DischargeDate      string `json:"pAG"`
				Specialist         string `json:"pAz"`
				SpecialistPosition string `json:"puR"`
				AdmissionType      string `json:"pvs"`
				Department         string `json:"pID"`
				DischargeId        string `json:"qqc"`
				ReceiptDate        string `json:"datqq"`
			} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

type ConsistedPatient struct {
	DischargeDiagnosis string `json:"pKDiag"`
	Department         string `json:"pID"`
	ReceiptDate        string `json:"datqq"`
}
type DischargeInfo struct {
	DischargeDiagnosis string `json:"pKDiag"`
	DischargeDate      string `json:"pAG"`
	Specialist         string `json:"pAz"`
	SpecialistPosition string `json:"puR"`
	AdmissionType      string `json:"pvs"`
	Department         string `json:"pID"`
	DischargeId        string `json:"qqc"`
}
type DischargeUnit struct { //итоговая структура
	Gender             string `json:"pJ"`
	PatientId          string `json:"qqc"`
	Department         string `json:"pID"`
	DischargeDiagnosis string `json:"pKDiag"`
	DischargeDate      string `json:"pAG"`
	Specialist         string `json:"pAz"`
	SpecialistPosition string `json:"puR"`
	AdmissionType      string `json:"pvs"`
}

type PatientInfo struct {
	PatientId   string
	RegId       string
	Gender      string
	DateOfBirth string
}

type CategorySubcategory struct {
	Category    string
	Subcategory string
	OKMUCode    string
}

type SpecialistInfo struct {
	SpecialistId      string
	DateOfCompletion  string
	PositionResource  string
	DestinationStatus string
}
type ServiceInfo struct {
	Date            string //дата
	OKMUCode        string //код ОКМУ
	AppointmentId   string //qqc
	Service         string //услуга
	Specialist      string //Специалист
	Department      string //Отделение
	AppointmentType string //Тип назначения
	RegId           string //Регистрационный номер пациента
}
type ServiceUnit struct { //итоговая структура
	Specialist        string //фио специалиста
	Service           string //услуга
	AppointmentType   string //Тип назначения
	AppointmentId     string //qqc
	SpecialistId      string //qqc
	PatientId         string //qqc
	ServiceId         string //qqc
	Date              string //дата
	OKMUCode          string //код ОКМУ
	Department        string //Отделение
	RegId             string //Регистрационный номер пациента
	Gender            string //Пол пациента
	Age               int
	DestinationStatus string //Состояние назначения
	PositionResource  string // Должность_Ресурс
	Category          string //Категория услуги
	Subcategory       string //Субкатегория услуги
}