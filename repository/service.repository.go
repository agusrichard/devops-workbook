package repository

import (
	"fmt"
	"golang-restapi/model"
)

// CreateServiceRequest -- create service request
func CreateServiceRequest(service *model.Service) {
	sqlQuery := `
		INSERT INTO services (
			request_id, 
			status, 
			vessel_name, 
			service_type,
			data_agent,
			cargo,
			etd,
			eta
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8);
	`

	_, err := DB.Exec(sqlQuery,
		service.RequestID,
		service.Status,
		service.VesselName,
		service.ServiceType,
		service.DataAgent,
		service.Cargo,
		service.ETD,
		service.ETA,
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Success to create service request %v\n", *service)
}
