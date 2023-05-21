package main

import (
	"dojo/testpackage"
	"fmt"

	trackerRpc "github.com/tnh9570/RPCpackage/tracker"
)

func main() {

	payload := trackerRpc.ResvPayload{
		Name: "John Doe",
		Oid:  "123456",
		StateCard: trackerRpc.StateCard{
			Card: trackerRpc.PatientCard{
				RsvGroupId:   1,
				ScheduleidId: 2,
				Pname:        "John Doe",
				// ... (set other fields of PatientCard)
			},
			TodList: []trackerRpc.Todo{
				{
					ReservationId:   1001,
					MedicalRecordId: 2001,
					// ... (set other fields of Todo)
				},
				{
					ReservationId:   1002,
					MedicalRecordId: 2002,
					// ... (set other fields of Todo)
				},
			},
		},
	}

	result, err := testpackage.ToStateCard(payload.StateCard)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)

}
