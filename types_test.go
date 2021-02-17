package branchio

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_Types_CustomInteger_MarshalJSONSuccess(t *testing.T) {
	c := CustomInteger{}
	c.Integer = 10
	result, err := c.MarshalJSON()
	expected := []byte(`10`)
	assert.Equal(t, expected, result)
	assert.Nil(t, err)
}

func Test_Types_CustomInteger_MarshalJSONEmpty(t *testing.T) {
	c := CustomInteger{}
	result, err := c.MarshalJSON()
	expected := []byte(`0`)
	assert.Equal(t, expected, result)
	assert.Nil(t, err)
}

func Test_Types_CustomInteger_UnmarshalCSVFilled(t *testing.T) {
	c := CustomInteger{}
	str := "12345"
	_ = c.UnmarshalCSV(str)
	assert.Equal(t, 12345, c.Value())
}

func Test_Types_CustomInteger_UnmarshalCSVEmpty(t *testing.T) {
	c := CustomInteger{}
	str := ""
	_ = c.UnmarshalCSV(str)
	assert.Equal(t, 0, c.Value())
}

func Test_Types_CustomInteger_UnmarshalCSVError(t *testing.T) {
	c := CustomInteger{}
	str := "foo"
	err := c.UnmarshalCSV(str)
	assert.Error(t, err)
	assert.Equal(t, `CustomInteger@UnmarshalCSV Parse int: strconv.Atoi: parsing "foo": invalid syntax`, err.Error())
}

func Test_Types_CustomFloat64_MarshalJSONSuccess(t *testing.T) {
	c := CustomFloat64{}
	c.Float64 = 10.10
	result, err := c.MarshalJSON()
	expected := []byte(`10.1`)
	assert.Equal(t, expected, result)
	assert.Nil(t, err)
}

func Test_Types_CustomFloat64_MarshalJSONEmpty(t *testing.T) {
	c := CustomFloat64{}
	result, err := c.MarshalJSON()
	expected := []byte(`0`)
	assert.Equal(t, expected, result)
	assert.Nil(t, err)
}

func Test_Types_CustomFloat64_UnmarshalCSVFilled(t *testing.T) {
	c := CustomFloat64{}
	str := "123.45"
	_ = c.UnmarshalCSV(str)
	assert.Equal(t, 123.44999694824219, c.Value())
}

func Test_Types_CustomFloat64_UnmarshalCSVEmpty(t *testing.T) {
	c := CustomFloat64{}
	str := ""
	_ = c.UnmarshalCSV(str)
	assert.Equal(t, float64(0), c.Value())
}

func Test_Types_CustomFloat64_UnmarshalCSVError(t *testing.T) {
	c := CustomFloat64{}
	str := "foo"
	err := c.UnmarshalCSV(str)
	assert.Error(t, err)
	assert.Equal(t, `CustomFloat64@UnmarshalCSV Parse float: strconv.ParseFloat: parsing "foo": invalid syntax`, err.Error())
}

func Test_Types_CustomTimestamp_UnmarshalCSVFilled(t *testing.T) {
	c := CustomTimestamp{}
	str := "2020-09-10 15:15:15+0000"
	_ = c.UnmarshalCSV(str)
	assert.Equal(t, "2020-09-10 15:15:15+0000", c.Value().Format(CustomTimestampFormatDefault))
}

func Test_Types_CustomTimestamp_UnmarshalCSVEmpty(t *testing.T) {
	c := CustomTimestamp{}
	str := ""
	_ = c.UnmarshalCSV(str)
	assert.True(t, c.Value().IsZero())
}

func Test_Types_CustomTimestamp_UnmarshalCSVError(t *testing.T) {
	c := CustomTimestamp{}
	str := "foo"
	err := c.UnmarshalCSV(str)
	assert.Error(t, err)
	assert.Equal(t, `CustomTimestamp@UnmarshalJSON ParseTime: parsing time "foo" as "2006-01-02 15:04:05-0700": cannot parse "foo" as "2006"`, err.Error())
}

func Test_Types_CustomTimestamp_MarshalJSONSuccess(t *testing.T) {
	c := CustomTimestamp{}
	c.Timestamp = time.Date(2020, time.Month(1), 20, 0, 0, 0, 0, time.UTC)
	result, err := c.MarshalJSON()
	assert.Equal(t, []byte(`"2020-01-20 00:00:00+0000"`), result)
	assert.Nil(t, err)
}

func Test_Types_CustomTimestamp_MarshalJSONEmpty(t *testing.T) {
	c := CustomTimestamp{}
	result, err := c.MarshalJSON()
	assert.Equal(t, []byte(`""`), result)
	assert.Nil(t, err)
}

func Test_Types_CustomDate_UnmarshalCSVFilled(t *testing.T) {
	c := CustomDate{}
	str := "2020-09-10"
	_ = c.UnmarshalCSV(str)
	assert.Equal(t, "2020-09-10", c.Value().Format(CustomDateFormatDefault))
}

func Test_Types_CustomDate_UnmarshalCSVFilledWithSlash(t *testing.T) {
	c := CustomDate{}
	str := "09/10/2020"
	_ = c.UnmarshalCSV(str)
	assert.Equal(t, "2020-09-10", c.Value().Format(CustomDateFormatDefault))
}

func Test_Types_CustomDate_UnmarshalCSVEmpty(t *testing.T) {
	c := CustomDate{}
	str := ""
	_ = c.UnmarshalCSV(str)
	assert.True(t, c.Value().IsZero())
}

func Test_Types_CustomDate_UnmarshalCSVError(t *testing.T) {
	c := CustomDate{}
	str := "foo"
	err := c.UnmarshalCSV(str)
	assert.Error(t, err)
	assert.Equal(t, `CustomDate@UnmarshalJSON ParseTime: parsing time "foo" as "2006-01-02": cannot parse "foo" as "2006"`, err.Error())
}

func Test_Types_CustomDate_MarshalJSONSuccess(t *testing.T) {
	c := CustomDate{}
	c.Date = time.Date(2020, time.Month(1), 20, 0, 0, 0, 0, time.UTC)
	result, err := c.MarshalJSON()
	assert.Equal(t, []byte(`"2020-01-20"`), result)
	assert.Nil(t, err)
}

func Test_Types_CustomDate_MarshalJSONEmpty(t *testing.T) {
	c := CustomDate{}
	result, err := c.MarshalJSON()
	assert.Equal(t, []byte(`""`), result)
	assert.Nil(t, err)
}

func Test_Types_CustomBoolean_UnmarshalCSVTrue(t *testing.T) {
	c := CustomBoolean{}
	str := "true"
	_ = c.UnmarshalCSV(str)
	assert.Equal(t, true, c.Value())
}

func Test_Types_CustomBoolean_UnmarshalCSVFalse(t *testing.T) {
	c := CustomBoolean{}
	str := "false"
	_ = c.UnmarshalCSV(str)
	assert.Equal(t, false, c.Value())
}

func Test_Types_CustomBoolean_UnmarshalCSVEmpty(t *testing.T) {
	c := CustomBoolean{}
	str := ""
	_ = c.UnmarshalCSV(str)
	assert.Equal(t, false, c.Value())
}

func Test_Types_CustomBoolean_UnmarshalCSVError(t *testing.T) {
	c := CustomBoolean{}
	str := "foo"
	_ = c.UnmarshalCSV(str)
	assert.Equal(t, false, c.Value())
}

func Test_Types_CustomBoolean_MarshalJSONSuccess(t *testing.T) {
	c := CustomBoolean{}
	c.Boolean = true
	result, err := c.MarshalJSON()
	assert.Equal(t, []byte(`true`), result)
	assert.Nil(t, err)
}

func Test_Types_CustomBoolean_MarshalJSONEmpty(t *testing.T) {
	c := CustomBoolean{}
	result, err := c.MarshalJSON()
	assert.Equal(t, []byte(`false`), result)
	assert.Nil(t, err)
}
