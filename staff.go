package gokayako

import (
	"encoding/xml"
	"errors"
)

type Staff struct {
	XMLName      xml.Name `xml:"staff"`
	ID           int      `xml:"id"`
	StaffGroupID int      `xml:"staffgroupid"`
	FirstName    string   `xml:"firstname"`
	LastName     string   `xml:"lastname"`
	FullName     string   `xml:"fullname"`
	UserName     string   `xml:"username"`
	Email        string   `xml:"email"`
	Designation  string   `xml:"designation"`
	Greeting     string   `xml:"greeting"`
	MobileNumber string   `xml:"mobilenumber"`
	Enabled      string   `xml:"isenabled"`
	Timezone     string   `xml:"timezone"`
	EnabledSt    string   `xml:"enabledst"`
	Signature    string   `xml:"signature"`
}

type StaffUsers struct {
	XMLName xml.Name `xml:"staffusers"`
	Staff   []Staff  `xml:"staff"`
}

func (kayako *Kayako) GetStaffByID(id int) (*Staff, error) {
	staff, err := kayako.GetStaff()
	if err != nil {
		return nil, err
	}

	for _, s := range staff.Staff {
		if s.ID == id {
			return &s, nil
		}
	}

	return nil, errors.New("Staff ID not found.")
}

func (kayako *Kayako) GetStaff() (*StaffUsers, error) {
	var s StaffUsers

	body, err := kayako.buildAndGetBody("/Base/Staff", nil)
	if err != nil {
		return nil, err
	}

	err = xml.Unmarshal(body, &s)
	if err != nil {
		return nil, err
	}

	return &s, nil
}
