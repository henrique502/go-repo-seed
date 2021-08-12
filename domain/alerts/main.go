package alerts

import (
	"log"
	"time"

	"github.com/henrique502/go-repo-seed/infra"
	"github.com/henrique502/go-repo-seed/infra/opsgenie"
)

func syncAlerts(alerts []infra.OpsgenieListAlert) {
	for index, element := range alerts {

		log.Printf("Index: %d ID: %s", index, element.ID)
	}
}

func FetchDay(day time.Time) {

	log.Println("Fetch date: " + day.Format("2006-01-02"))
	aletsIterator(day)

}

func aletsIterator(day time.Time) {
	data := opsgenie.GetAlertList(time.Now())
	syncAlerts(data.Data)

	for len(data.Paging.Next) > 0 {
		data = opsgenie.GetAlertListPagination(data.Paging.Next)
		syncAlerts(data.Data)
	}

}
