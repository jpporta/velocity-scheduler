package velocity

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type VelocityClasses struct {
	Token      string    `json:"token"`
	Instructor int       `json:"instructor"`
	Duration   int       `json:"duration_time"`
	Start      time.Time `json:"start_time"`
}

type getClassesRes struct {
	Count   int               `json:"count"`
	Results []VelocityClasses `json:"results"`
}

func (v Velocity) GetClasses(unitId int, from time.Time) []VelocityClasses {
	// 30 Days from now
	nextMonth := from.Add(time.Hour * 24 * 30)
	u, err := url.Parse(baseURL + "/events/schedule")
	if err != nil {
		log.Fatalln("Failed to parse base URL", err)
	}
	params := url.Values{}
	params.Add("page", "1")
	params.Add("date_from", from.Format(time.DateOnly))
	params.Add("sort", "date_from")
	params.Add("is_canceled", "false")
	params.Add("date_to", nextMonth.Format(time.DateOnly))
	params.Add("unit_list", strconv.Itoa(unitId))
	params.Add("activity_list", "1")
	params.Add("timezone_from_unit", strconv.Itoa(unitId))

	u.RawQuery = params.Encode()

	log.Println(u)
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		log.Fatalln("Failed to created Request", err)
	}
	req.Header.Set("Authorization", "JWT "+v.token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln("Failed to get classes", err)
	}
	defer res.Body.Close()
	in, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatalln("Failed read body", err)
	}

	var response getClassesRes
	err = json.Unmarshal(in, &response)

	log.Println("Fetched " + strconv.Itoa(response.Count) + " classes.")

	return response.Results
}
