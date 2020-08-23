package freshservice

import (
	"fmt"
	"time"
)

//Ticket ...
type Ticket struct {
	ID int `json:"id,omitempty"`
	//Name of the requester
	Name string `json:"name,omitempty"`
	//Email address of the requester. If no contact exists with this email address in Freshservice, it will be added as a new contact.
	Email string `json:"email,omitempty"`
	//Phone number of the requester. If no contact exists with this phone number in Freshservice, it will be added as a new contact.
	Phone string `json:"phone,omitempty"`
	//Subject of the ticket. The default value is null.
	Subject string `json:"subject,omitempty"`
	//Type Helps categorize the ticket according to the different kinds of issues your support team deals with. The default Value is incident. * As of now, API v2 supports only type ‘incident’
	Type string `json:"type,omitempty"`
	//Status of the ticket.
	Status int `json:"status,omitempty"`
	//Priority of the ticket.
	Priority int `json:"priority,omitempty"`
	//Description HTML content of the ticket.
	Description string `json:"description,omitempty"`
	//CCEmails Email address added in the 'cc' field of the incoming ticket email.
	CCEmails     []string               `json:"cc_emails,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	DueBy        *time.Time             `json:"due_by,omitempty"`
	Source       int                    `json:"source,omitempty"`
	Tags         []string               `json:"tags,omitempty"`
	Category     string                 `json:"category,omitempty"`
}

//TicketList holds a list of tickets
type TicketList struct {
	Tickets []Ticket `json:"tickets,omitempty"`
}

//TicketItem holds one ticket
type TicketItem struct {
	Ticket Ticket `json:"ticket,omitempty"`
}

//CreateTicket creates a ticket
func (c *Client) CreateTicket(ticket *Ticket) (*Ticket, error) {
	var out TicketItem
	err := c.WriteObject("/api/v2/tickets", "POST", ticket, &out)
	if err != nil {
		return nil, err
	}
	return &out.Ticket, nil
}

//ListTickets get locations by page
func (c *Client) ListTickets(page int) ([]Ticket, error) {
	var list TicketList
	err := c.ReadObject(fmt.Sprintf("/api/v2/tickets?page=%v", page), &list)
	return list.Tickets, err
}

//ListAllTickets read all locations
func (c *Client) ListAllTickets() ([]Ticket, error) {
	all := make([]Ticket, 0)
	i := 1
	for true {
		page, err := c.ListTickets(i)
		if err != nil {
			return nil, err
		}
		if len(page) == 0 {
			break
		}
		all = append(all, page...)
		i++
	}
	return all, nil
}
