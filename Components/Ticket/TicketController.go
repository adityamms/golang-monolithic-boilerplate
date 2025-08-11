package Ticket

import (
	"github.com/gin-gonic/gin"
	"github.com/mahdidl/golang_boilerplate/Common/Helper"
	General "github.com/mahdidl/golang_boilerplate/Common/Response"
	"github.com/mahdidl/golang_boilerplate/Common/Validator"
	"github.com/mahdidl/golang_boilerplate/dto"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"log"

	"net/http"
)

type TicketController struct {
	ticketService *TicketService
}

func NewTicketController(ticketService *TicketService) *TicketController {
	return &TicketController{ticketService: ticketService}
}

func (ticketControler *TicketController) CreateTicket(context *gin.Context) {
	var ticketRequest dto.CreateTicketRequest
	Helper.Decode(context.Request, &ticketRequest)

	validationError := Validator.ValidationCheck(ticketRequest)
	log.Println(validationError)
	if validationError != nil {
		response := General.GeneralResponse{Error: true, Message: validationError.Error()}
		context.JSON(http.StatusBadRequest, gin.H{"response": response})
		return
	}

	userId := context.Param("userId")

	validationErr := primitive.IsValidObjectID(userId)

	if !validationErr {
		response := General.GeneralResponse{Error: true, Message: "id is not valid"}
		context.JSON(http.StatusBadRequest, gin.H{"response": response})
		return
	}

	ticketResponse, responseError := ticketControler.ticketService.CreateTicket(ticketRequest, userId)

	if responseError != nil {
		context.JSON(http.StatusBadRequest, gin.H{"response": General.ErrorResponse{Error: responseError.Error()}})
		return
	}

	// all ok
	// create general response
	response := General.GeneralResponse{Error: false, Message: "ticket have been created", Data: dto.CreateTicketResponse{Message: ticketResponse.Message, Subject: ticketResponse.Subject, Image: ticketResponse.Image}}
	context.JSON(http.StatusOK, gin.H{"response": response})
}
