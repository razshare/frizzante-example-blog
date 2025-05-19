package sessions

import "time"

func InitialData() Data {
	return Data{
		Items: []Item{
			{Checked: false, Description: "Pet the cat."},
			{Checked: false, Description: "Do laundry"},
			{Checked: false, Description: "Pet the cat."},
			{Checked: false, Description: "Cook"},
			{Checked: false, Description: "Pet the cat."},
		},
		LastActivity: time.Now(),
		Expired:      false,
		Verified:     true,
	}
}

type Data struct {
	Items        []Item    `json:"items"`
	LastActivity time.Time `json:"lastActivity"`
	Verified     bool      `json:"verified"`
	Expired      bool      `json:"expired"`
	AccountId    string    `json:"accountId"`
}

type Item struct {
	Checked     bool   `json:"checked"`
	Description string `json:"description"`
}
