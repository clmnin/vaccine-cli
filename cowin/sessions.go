package cowin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Center struct {
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	BlockName string    `json:"block_name"`
	Sessions  []Session `json:"sessions"`
}

type Session struct {
	Date              string   `json:"date"`
	AvailableCapacity float64  `json:"available_capacity"`
	MinAge            int      `json:"min_age_limit"`
	Vaccine           string   `json:"vaccine"`
	Slots             []string `json:"slots"`
	AvailableDose1    int      `json:"available_capacity_dose1"`
	AvailableDose2    int      `json:"available_capacity_dose2"`
}

type CenterDetails struct {
	Name              string   `json:"name"`
	Address           string   `json:"address"`
	BlockName         string   `json:"block_name"`
	Date              string   `json:"date"`
	AvailableCapacity float64  `json:"available_capacity"`
	MinAge            int      `json:"min_age_limit"`
	Vaccine           string   `json:"vaccine"`
	Slots             []string `json:"slots"`
}

func GetAvailableSessions(districtID int, minAge int) string {
	centers := getSchedule(districtID, minAge)
	if len(centers) == 0 {
		return ""
	}
	details := getFormattedCenters(centers)
	return details
}

func getSchedule(districtID int, minAge int) []CenterDetails {
	date := today()
	req, err := http.NewRequest("GET", fmt.Sprintf(baseUrl+availabilityByDistrictUrl+"?district_id=%d&date=%s", districtID, date), nil)
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.93 Safari/537.36")
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		panic("Failed to fetch")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		panic("Failed to read response")
	}
	c := struct {
		Center []Center `json:"centers"`
	}{}
	err = json.Unmarshal(body, &c)
	if err != nil {
		log.Println(err)
		panic("Failed to parse json")
	}
	return filterCentersByAgeLimit(c.Center, minAge)
}

func filterCentersByAgeLimit(center []Center, minAge int) []CenterDetails {
	// We get a list of sessions (district was choosen in the prev fn)
	// and from this we get a list of sessions the users is requesting for
	availableCenters := make([]CenterDetails, 0)

	for _, center := range center {
		for _, session := range center.Sessions {
			if session.AvailableCapacity > 0 && minAge >= session.MinAge {
				centerDetails := CenterDetails{
					Name:              center.Name,
					Address:           center.Address,
					BlockName:         center.BlockName,
					AvailableCapacity: session.AvailableCapacity,
					Date:              session.Date,
					MinAge:            session.MinAge,
					Vaccine:           session.Vaccine,
					Slots:             session.Slots,
				}
				availableCenters = append(availableCenters, centerDetails)
			}
		}
	}

	return availableCenters
}

func getFormattedCenters(centers []CenterDetails) string {
	details := ""
	for _, center := range centers {
		details += fmt.Sprintln("Available:", center.AvailableCapacity)
		details += fmt.Sprintln("Name:", center.Name)
		details += fmt.Sprintln("Address:", center.Address)
		details += fmt.Sprintln("Block:", center.BlockName)
		details += fmt.Sprintln("Date:", center.Date)
		details += fmt.Sprintln("Vaccine:", center.Vaccine)
		details += fmt.Sprintln("Slots:", center.Slots)
		details += fmt.Sprintln("Min Age:", center.MinAge)
		details += fmt.Sprintln("=========================================")
		details += fmt.Sprint("\n\n")
	}
	return details
}
