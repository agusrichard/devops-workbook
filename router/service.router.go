package router

import (
	"golang-restapi/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateServiceRequest -- create service request
func CreateServiceRequest(c *gin.Context) {
	var serviceRequest model.Service
	c.Bind(&serviceRequest)
	c.JSON(http.StatusOK, gin.H{
		"message": "Success to create service request",
		"data": gin.H{
			"requestID":   serviceRequest.RequestID,
			"status":      serviceRequest.Status,
			"vesselName":  serviceRequest.VesselName,
			"serviceType": serviceRequest.ServiceType,
			"dataAgent":   serviceRequest.DataAgent,
			"cargo":       serviceRequest.Cargo,
			"etd":         serviceRequest.ETD,
			"eta":         serviceRequest.ETA,
		},
	})
}
