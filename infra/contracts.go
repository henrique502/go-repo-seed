package infra

import "time"

type Postgre interface {
}

type Opsgenie interface {
}

type OpsgenieListResponse struct {
	Paging struct {
		Next  string `json:"next"`
		First string `json:"first"`
		Last  string `json:"last"`
	} `json:"paging"`
	Took      float64 `json:"took"`
	RequestID string  `json:"requestId"`
}

type OpsgenieListAlert struct {
	ID             string    `json:"id"`
	TinyID         string    `json:"tinyId"`
	Alias          string    `json:"alias"`
	Message        string    `json:"message"`
	Status         string    `json:"status"`
	Acknowledged   bool      `json:"acknowledged"`
	IsSeen         bool      `json:"isSeen"`
	Tags           []string  `json:"tags"`
	Snoozed        bool      `json:"snoozed"`
	SnoozedUntil   time.Time `json:"snoozedUntil"`
	Count          int       `json:"count"`
	LastOccurredAt time.Time `json:"lastOccurredAt"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	Source         string    `json:"source"`
	Owner          string    `json:"owner"`
	Priority       string    `json:"priority"`
	Responders     []struct {
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"responders"`
	Integration struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"integration"`
	Report struct {
		AckTime        int    `json:"ackTime"`
		CloseTime      int    `json:"closeTime"`
		AcknowledgedBy string `json:"acknowledgedBy"`
		ClosedBy       string `json:"closedBy"`
	} `json:"report"`
}

type OpsgenieListAlertResponse struct {
	OpsgenieListResponse
	Data []OpsgenieListAlert `json:"data"`
}
