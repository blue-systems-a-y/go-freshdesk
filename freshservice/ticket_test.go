package freshservice

// func TestClient_CreateTicket(t *testing.T) {
// 	c := NewClient(domain, apiKey)
// 	tkt := &Ticket{
// 		Name:        "danny",
// 		Email:       "danny@gabriel.network",
// 		Subject:     "Something does not !",
// 		Description: "Something is not working as i think it should",
// 		Status:      2,
// 		Priority:    2,
// 		Source:      2,
// 		Tags:        []string{"one"},
// 	}

// 	ticket, err := c.CreateTicket(tkt)
// 	assert.NoError(t, err)
// 	assert.NotEmpty(t, ticket.ID)
// }

// func TestClient_ListAllTickets(t *testing.T) {
// 	c := NewClient(domain, apiKey)
// 	tickets, err := c.ListAllTickets()
// 	assert.NoError(t, err)
// 	for _, ticket := range tickets {
// 		log.Printf("%+v", ticket)
// 	}
// }
