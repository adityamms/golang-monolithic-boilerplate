package Ticket

import (
	"github.com/mahdidl/golang_boilerplate/Common/Validator"

	Controller "github.com/mahdidl/golang_boilerplate/Components/User"
	"github.com/mahdidl/golang_boilerplate/dto"
)

type TicketService struct {
	ticketRepository *TicketRepository
	userService      *Controller.UserService
}

func NewTicketService(userService *Controller.UserService, ticketRepository *TicketRepository) *TicketService {
	return &TicketService{userService: userService, ticketRepository: ticketRepository}
}

func (ticketService TicketService) CreateTicket(createTicketRequest dto.CreateTicketRequest, userId string) (dto.CreateTicketResponse, error) {
	// validate username len and not empty
	validationError := Validator.ValidationCheck(createTicketRequest)

	if validationError != nil {
		return dto.CreateTicketResponse{}, validationError
	}

	ticket, ticketError := ticketService.ticketRepository.Create(createTicketRequest, userId)
	if ticketError != nil {
		return dto.CreateTicketResponse{}, ticketError
	}
	// we need a transformer
	return dto.CreateTicketResponse{Subject: ticket.Subject, Message: ticket.Message, Image: ticket.Image}, nil
}
