package redash_json_format

import (
	"fmt"
	"testing"

	"github.com/cheekybits/is"
)

var desiredJSONString = `
{
  "columns": [
    {
      "name": "date",
      "type": "date",
      "friendly_name": "date"
    },
    {
      "name": "day_number",
      "type": "integer",
      "friendly_name": "day_number"
    },
    {
      "name": "value",
      "type": "integer",
      "friendly_name": "value"
    },
    {
      "name": "total",
      "type": "integer",
      "friendly_name": "total"
    }
  ],
  "rows": [
    {
      "value": 40832,
      "total": 53141,
      "day_number": 0,
      "date": "2014-01-30"
    },
    {
      "value": 27296,
      "total": 53141,
      "day_number": 1,
      "date": "2014-01-30"
    },
    {
      "value": 22982,
      "total": 53141,
      "day_number": 2,
      "date": "2014-01-30"
    }
  ]
}
`

func TestUnmarshalRedashFormattedJSON(t *testing.T) {
	is := is.New(t)
	var err error

	r, err := NewRedashFormattedJSON()
	is.NoErr(err)

	c, err := NewColumn("date", "date", "date")
	is.NoErr(err)
	err = r.AddColumn(c)
	is.NoErr(err)

	c2, err := NewColumn("day_number", "integer", "day_number")
	is.NoErr(err)
	err = r.AddColumn(c2)
	is.NoErr(err)
	c3, err := NewColumn("value", "integer", "value")
	is.NoErr(err)
	err = r.AddColumn(c3)
	is.NoErr(err)
	c4, err := NewColumn("total", "integer", "total")
	is.NoErr(err)
	err = r.AddColumn(c4)
	is.NoErr(err)

	row := map[string]interface{}{
		"date":       "2014-01-30",
		"day_number": 0,
		"total":      53141,
		"value":      40832,
	}

	err = r.AddRow(row)
	is.NoErr(err)

	row2 := map[string]interface{}{
		"date":       "2014-01-30",
		"day_number": 0,
		"total":      53141,
		"value":      40832,
	}

	err = r.AddRow(row2)
	is.NoErr(err)

	out, err := r.Marshal()
	is.NoErr(err)

	fmt.Println(string(out))

}
