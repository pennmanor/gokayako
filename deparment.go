package gokayako

import (
	"encoding/xml"
	"errors"
)

type Department struct {
	XMLName              xml.Name `xml:"department"`
	ID                   int      `xml:"id"`
	Title                string   `xml:"title"`
	Type                 string   `xml:"type"`
	App                  string   `xml:"app"`
	DisplayOrder         string   `xml:"displayorder"`
	Parent               string   `xml:"parentDepartment"`
	UserVisibilityCustom int      `xml:"uservisibilitycustom"`
}

type Departments struct {
	XMLName   xml.Name     `xml:"departments"`
	Deparment []Department `xml:"department"`
}

func (kayako *Kayako) GetDepartmentID(name string) (int, error) {
	departments, err := kayako.GetDepartments()

	if err != nil {
		return -1, err
	}

	for _, d := range departments.Deparment {
		if d.Title == name {
			return d.ID, nil
		}
	}

	return -1, errors.New("Deparment not found.")
}

func (kayako *Kayako) GetDepartments() (*Departments, error) {

	var d Departments

	body, err := kayako.buildAndGetBody("/Base/Department", nil)
	if err != nil {
		return nil, err
	}

	err = xml.Unmarshal(body, &d)
	if err != nil {
		return nil, err
	}

	return &d, nil
}
