package prome

import (
	"net/http"
	"reflect"
	"strings"

	"github.com/cocktail828/go-tools/xlog/colorful"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	registerer = prometheus.NewRegistry()
	metrics    = []*Metric{}
)

type Metric struct {
	metric any
	getter func(reflect.Value) float64
}

func init() {
	rt := reflect.TypeOf(GlobalMetrics{})
	for i := 0; i < rt.NumField(); i++ {
		rsf := rt.Field(i)
		metrics = append(metrics, &Metric{
			promauto.With(registerer).NewGauge(prometheus.GaugeOpts{
				Name: strings.ReplaceAll(rsf.Tag.Get("json"), ".", "_"),
				Help: rsf.Tag.Get("desc"),
			}),
			func(rv reflect.Value) float64 {
				switch v := rv.Field(i).Interface().(type) {
				case int64:
					return float64(v)
				case float64:
					return v
				default:
					colorful.Errorf("unknow type %t", v)
				}
				return 0
			},
		})
	}
}

func Handler() http.Handler {
	return promhttp.HandlerFor(registerer, promhttp.HandlerOpts{Registry: registerer})
}

func Update(gm *GlobalMetrics) {
	rv := reflect.ValueOf(gm).Elem()
	for _, m := range metrics {
		val := m.getter(rv)

		switch v := m.metric.(type) {
		case prometheus.Gauge:
			v.Set(val)
		case prometheus.Histogram:
			v.Observe(val)
		default:
			colorful.Errorf("unknow type %t", v)
		}
	}
}
