package gokayako

import (
	"encoding/xml"
	"errors"
	"fmt"
	"strconv"
)

type TicketStatus struct {
	XMLName               xml.Name `xml:"ticketstatus"`
	ID                    int      `xml:"id"`
	Title                 string   `xml:"title"`
	DisplayOrder          string   `xml:"displayorder"`
	Type                  string   `xml:"type"`
	DisplayInMainList     int      `xml:"displayinmainlist"`
	MarkAsResolved        int      `xml:"MarkAsResolved"`
	DisplayCount          int      `xml:"DisplayCount"`
	StatusColor           string   `xml:"StatusColor"`
	StatusBgColor         string   `xml:"statusbgcolor"`
	ResetDueTime          int      `xml:"resetduetime"`
	StaffVisibilityCustom int      `xml:"staffvisibilitycustom"`
}

type TicketStatuses struct {
	XMLName        xml.Name       `xml:"ticketstatuses"`
	TicketStatuses []TicketStatus `xml:"ticketstatus"`
}

type Tickets struct {
	XMLName xml.Name `xml:"tickets"`
	Tickets []Ticket `xml:"ticket"`
}

type Ticket struct {
	XMLName            xml.Name `xml:"ticket"`
	ID                 int      `xml:"id,attr"`
	FlagType           int      `xml:"flagtype,attr"`
	DisplayID          string   `xml:"displayid"`
	DepartmentID       int      `xml:"departmentid"`
	StatusID           int      `xml:"statusid"`
	PriorityID         int      `xml:"priorityid"`
	TypeID             int      `xml:"typeid"`
	UserID             int      `xml:"userid"`
	UserOrganization   string   `xml:"userorganization"`
	UserOrganizationID string   `xml:"userorganizationid"`
	OwnerStaffID       int      `xml:"ownerstaffid"`
	FullName           string   `xml:"fullname"`
	Email              string   `xml:"email"`
	LastReplier        string   `xml:"lastreplier"`
	Subject            string   `xml:"subject"`
}

type TicketPriorities struct {
	XMLName          xml.Name         `xml:"ticketpriorities"`
	TicketPriorities []TicketPriority `xml:"ticketpriority"`
}

type TicketPriority struct {
	XMLName   xml.Name `xml:"ticketpriority"`
	ID        int      `xml:"id"`
	Title     string   `xml:"title"`
	ColorCode string   `xml:"frcolorcode"`
}

func (kayako *Kayako) GetTicketPriorityByID(id int) (*TicketPriority, error) {
	priorities, err := kayako.GetTicketPriorities()
	if err != nil {
		return nil, nil
	}

	for _, p := range priorities.TicketPriorities {
		if p.ID == id {
			return &p, nil
		}
	}

	return nil, errors.New("Priority ID not found")

}

func (kayako *Kayako) GetTicketPriorities() (*TicketPriorities, error) {
	var p TicketPriorities

	body, err := kayako.buildAndGetBody("/Tickets/TicketPriority", nil)
	if err != nil {
		return nil, err
	}

	err = xml.Unmarshal(body, &p)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (kayako *Kayako) GetTicketStatusID(name string) (int, error) {
	ticketstatuses, err := kayako.GetTicketStatuses()
	if err != nil {
		return -1, err
	}

	for _, ts := range ticketstatuses.TicketStatuses {
		if ts.Title == name {
			return ts.ID, nil
		}
	}

	return -1, errors.New("Ticket Status not found")
}

func (kayako *Kayako) GetTicketStatuses() (*TicketStatuses, error) {
	var s TicketStatuses

	body, err := kayako.buildAndGetBody("/Tickets/TicketStatus", nil)
	if err != nil {
		return nil, err
	}

	err = xml.Unmarshal(body, &s)
	if err != nil {
		return nil, err
	}

	return &s, nil
}

func (kayako *Kayako) GetTickets(deptID int, statusID int, ownerID int, userID int) (*Tickets, error) {
	var t Tickets

	dep := strconv.Itoa(deptID)
	status := strconv.Itoa(statusID)
	owner := strconv.Itoa(ownerID)
	user := strconv.Itoa(userID)

	url := fmt.Sprintf("/Tickets/Ticket/ListAll/%v/%v/%v/%v/-1/-1/ticketid/ASC", dep, status, owner, user)

	body, err := kayako.buildAndGetBody(url, nil)
	if err != nil {
		return nil, err
	}

	err = xml.Unmarshal(body, &t)
	if err != nil {
		return nil, err
	}

	return &t, nil
}
