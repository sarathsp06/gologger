package logger

import (
	"encoding/json"
	"fmt"
)

//Metric format for metrics
type Metric struct {
	Name   string                 `json:"name"`
	Tags   map[string]string      `json:"tags"`
	Fields map[string]interface{} `json:"fields"`
}

//NewMetric creates and initialises the metric
func NewMetric(name string) *Metric {
	m := new(Metric)
	m.Name = name
	m.Tags = map[string]string{"process_name": processName, "process_id": fmt.Sprint(processID), "host_name": hostName}
	m.Fields = make(map[string]interface{})
	return m
}

//Tag sets a tag
func (m *Metric) Tag(key, value string) *Metric {
	m.Tags[key] = value
	return m
}

//Field sets field
func (m *Metric) Field(key string, value interface{}) *Metric {
	m.Fields[key] = value
	return m
}

//String implements Stringer interface
//json encode the object and returns
func (m Metric) String() string {
	m.Name = processName + "-" + m.Name
	marshalledData, err := json.Marshal(&m)
	if err != nil {
		return fmt.Sprintf(`{name:"%s-error-creating-metricjson",fields:{value:1}}`, processName)
	}
	return string(marshalledData) + "\n"
}
