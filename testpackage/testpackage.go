package testpackage

import (
	"errors"
)

type ReservationCard struct {
	Card    PatientCard `json:"reservationGroup"`
	TodList []Todo      `json:"reservation"`
}

type ScheduleCard struct {
	Card    PatientCard `json:"Schedule"`
	TodList []Todo      `json:"medicalRecord"`
}

type StateCard struct {
	Card    PatientCard `json:"patientCard"`
	TodList []Todo      `json:"todoList"`
}

type PatientCard struct {
	RsvGroupId       int64  `json:"groupId"`
	ScheduleidId     int64  `json:"scheduleId"`
	Pname            string `json:"name"`
	ChartNo          string `json:"chartNo"`
	Ssn              string `json:"ssn"`
	CellPhone        string `json:"cellPhone"`
	BirthDate        string `json:"birthDate"`
	Gender           string `json:"gender"`
	VisitNote        string `json:"visitNote"`
	CheckinAt        string `json:"checkinAt"`
	ReservationNote  string `json:"reservationNote"`
	PatientId        int64  `json:"patientId"`
	MedicalAreaCode  int    `json:"medicalAreaCode"`
	DetailedAreaCode int    `json:"detailedAreaCode"`
	DoctorId         int    `json:"doctorId"`
	EmployeeId       int    `json:"employeeId"`
	ProcedureCode    int8   `json:"procedureCode"`
	VisitType        int8   `json:"visitType"`
	InsType          int8   `json:"insType"`
	ScheduleStatus   int8   `json:"scheduleStatus"`
	Bookmark         int8   `json:"bookmark"`
	Pregnant         int8   `json:"pregnant"`
}

type Todo struct {
	ReservationId   int64  `json:"reservationId"`
	MedicalRecordId int64  `json:"medicalRecordId"`
	ReservationType string `json:"reservationType"`
	RsvState        int8   `json:"rsvState"`
	AppointmentAt   string `json:"appointedAt"`
	StartedAt       string `json:"startedAt"`
	CompletedAt     string `json:"completedAt"`
	ScheduleidId    int64  `json:"scheduleId"`
	DoctorId        int    `json:"doctorId"`
	EmployeeId      int    `json:"employeeId"`
	ProcType        int8   `json:"procType"`
	ProcStatus      int8   `json:"procStatus"`
	Seqno           int8   `json:"seqno"`
	ProcedureCode   int16  `json:"procedure"`
	ChangeLog       string `json:"changeLog"`
	ReservationNote string `json:"reservationNote"`
	History         string `json:"history"`
}

func ToStateCard(t interface{}) (StateCard, error) {
	if v, ok := t.(ReservationCard); ok {
		return StateCard{
			Card:    v.Card,
			TodList: v.TodList,
		}, nil
	} else if v, ok := t.(ScheduleCard); ok {
		return StateCard{
			Card:    v.Card,
			TodList: v.TodList,
		}, nil
	} else {
		return StateCard{}, errors.New("invalid type")
	}
}
