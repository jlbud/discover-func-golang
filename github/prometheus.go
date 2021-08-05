package github

import (
	"bytes"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func PushGauge() {
	gauge := prometheus.NewGaugeFunc(prometheus.GaugeOpts{
		Name: "cn_n2t_http_cur",
		Help: "The timestamp of the last successful completion of a DB backup.",
		ConstLabels: map[string]string{
			"interface": "/v1/core",
		},
	}, func() float64 {
		return 100
	})
	err := push.New("http://192.168.5.25:9091/", "cn-n2t").
		Collector(gauge).
		Grouping("instance", "request").
		Push()
	if err != nil {
		fmt.Println("Could not push completion time to Pushgateway:", err)
	}
}

func PushCount() {
	gauge := prometheus.NewCounterFunc(prometheus.CounterOpts{
		Name: "v1_core_count",
		Help: "The timestamp of the last successful completion of a DB backup.",
		ConstLabels: map[string]string{
			"interface": "/v1/core",
			"code":      "200",
		},
	}, func() float64 {
		return 100
	})
	err := push.New("http://192.168.5.25:9091/", "http_stat").
		Collector(gauge).
		Grouping("http", "request").
		Push()
	if err != nil {
		fmt.Println("Could not push completion time to Pushgateway:", err)
	}
}

func main() {
	params := RequestParams{
		Name:  "some_metric_v3",
		Type:  Counter,
		Value: 10.3,
		Labels: map[string]string{
			"interface": "/v1/core",
		},
	}
	Push(params)
}

const (
	Gauge   = "gauge"
	Counter = "counter"
)

type RequestParams struct {
	Name   string
	Type   string  // metric type
	Value  float64 // number that need to be counted
	Labels map[string]string

	JobName       string
	ServerAddress string // read from apollo config center
}

func Push(params RequestParams) error {
	buf := &bytes.Buffer{}
	buf.WriteString(buildMetricType(params.Name, params.Type))
	buf.WriteString(buildMetricData(params.Name, params.Labels, params.Value))

	inst, err := instance()
	if err != nil {
		return err
	}
	reqURI := fmt.Sprintf("/metrics/job/%s/instance/%s",
		url.QueryEscape(params.JobName),
		url.QueryEscape(inst))

	reqURL := fmt.Sprintf("%s%s", params.ServerAddress, reqURI)

	req, err := http.NewRequest("PUT", reqURL, buf)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "encoding=delimited")
	r, err := (&http.Client{}).Do(req)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return nil
}

func buildMetricType(name, metricType string) string {
	return fmt.Sprintf("# TYPE %s %s\n", name, metricType)
}

func buildMetricData(name string, labels map[string]string, value float64) string {
	buf := bytes.Buffer{}
	if len(labels) > 0 {
		buf.WriteString("{")
		for k, v := range labels {
			buf.WriteString(fmt.Sprintf("%s=\"%s\"", k, v))
		}
		buf.WriteString("}")
	}
	return fmt.Sprintf("%s %s %v\n", name, buf.String(), value)
}

func normalizeURL(url string) string {
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}
	return strings.TrimSuffix(url, "/")
}

func instance() (string, error) {
	host := os.Getenv("HOSTNAME")
	if host == "" {
		return "", fmt.Errorf("hostname not be empty, please contact system admin config environment")
	}
	return normalizeURL(host), nil
}
