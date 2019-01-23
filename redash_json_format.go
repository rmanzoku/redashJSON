package redash_json_format

import (
	"encoding/json"

	"github.com/pkg/errors"
)

type RedashFormattedJSON struct {
	Columns []*Column                `json:"columns"`
	Rows    []map[string]interface{} `json:"rows"`
}

type Column struct {
	Name         string `json:"name"`
	ColumnType   string `json:"type"`
	FriendlyName string `json:"friendly_name"`
}

func NewRedashFormattedJSON() (*RedashFormattedJSON, error) {
	return new(RedashFormattedJSON), nil
}

func NewColumn(name string, columnType string, friendlyName string) (*Column, error) {
	ret := &Column{
		name,
		columnType,
		friendlyName,
	}
	return ret, nil
}

func (r *RedashFormattedJSON) AddColumn(column *Column) (err error) {
	r.Columns = append(r.Columns, column)
	return nil
}

func (r *RedashFormattedJSON) AddRow(row map[string]interface{}) (err error) {
	if len(r.Columns) != len(row) {
		return errors.New("row elements must be equal column elements")
	}
	r.Rows = append(r.Rows, row)
	return nil
}

func UnmarshalRedashFormattedJSON(data []byte) (RedashFormattedJSON, error) {
	var r RedashFormattedJSON
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *RedashFormattedJSON) Marshal() ([]byte, error) {
	return json.Marshal(r)
}
