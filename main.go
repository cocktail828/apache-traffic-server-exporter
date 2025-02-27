package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/cocktail828/apache-traffic-server-exporter/prome"
	"github.com/cocktail828/go-tools/xlog/colorful"
	"github.com/cocktail828/go-tools/z/environ"
)

func getAtsMetrics(uri string) (*prome.GlobalMetrics, error) {
	resp, err := http.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var atsMetric prome.ATSMetrics
	if err := json.NewDecoder(resp.Body).Decode(&atsMetric); err != nil {
		return nil, err
	}

	return &atsMetric.Global, nil
}

func main() {
	go func() {
		uri := environ.String("ATS_STATS_URI")
		for {
			gm, err := getAtsMetrics(uri)
			if err != nil {
				colorful.Errorln(err)
				time.Sleep(time.Second)
				continue
			}

			prome.Update(gm)
			time.Sleep(10 * time.Second) // 每 10 秒更新一次
		}
	}()

	http.Handle("/debug/metrics", prome.Handler())
	http.ListenAndServe(":"+environ.String("ATS_LISTEN_PORT"), nil)
}
