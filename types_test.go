package branchio

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type TypesTestSuite struct {
	suite.Suite
}

func (suite *TypesTestSuite) TestCustomIntegerMarshalJSONSuccess() {
	c := CustomInteger{}
	c.Integer = 10
	result, err := c.MarshalJSON()
	expected := []byte(`10`)
	assert.Equal(suite.T(), expected, result)
	assert.Nil(suite.T(), err)
}

func (suite *TypesTestSuite) TestCustomIntegerMarshalJSONEmpty() {
	c := CustomInteger{}
	result, err := c.MarshalJSON()
	expected := []byte(`0`)
	assert.Equal(suite.T(), expected, result)
	assert.Nil(suite.T(), err)
}

func (suite *TypesTestSuite) TestCustomIntegerUnmarshalCSVFilled() {
	c := CustomInteger{}
	str := "12345"
	err := c.UnmarshalCSV(str)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 12345, c.Value())
}

func (suite *TypesTestSuite) TestCustomIntegerUnmarshalCSVEmpty() {
	c := CustomInteger{}
	str := ""
	err := c.UnmarshalCSV(str)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 0, c.Value())
}

func (suite *TypesTestSuite) TestCustomIntegerUnmarshalCSVError() {
	c := CustomInteger{}
	str := "foo"
	err := c.UnmarshalCSV(str)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), `CustomInteger@UnmarshalCSV Parse int: strconv.Atoi: parsing "foo": invalid syntax`, err.Error())
}

func (suite *TypesTestSuite) TestCustomFloat64MarshalJSONSuccess() {
	c := CustomFloat64{}
	c.Float64 = 10.10
	result, err := c.MarshalJSON()
	expected := []byte(`10.1`)
	assert.Equal(suite.T(), expected, result)
	assert.Nil(suite.T(), err)
}

func (suite *TypesTestSuite) TestCustomFloat64MarshalJSONEmpty() {
	c := CustomFloat64{}
	c.Float64 = 10.10
	result, err := c.MarshalJSON()
	expected := []byte(`10.1`)
	assert.Equal(suite.T(), expected, result)
	assert.Nil(suite.T(), err)
}

func (suite *TypesTestSuite) TestCustomFloat64UnmarshalCSVFilled() {
	c := CustomFloat64{}
	str := "123.45"
	err := c.UnmarshalCSV(str)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 123.44999694824219, c.Value())
}

func (suite *TypesTestSuite) TestCustomFloat64_UnmarshalCSVEmpty() {
	c := CustomFloat64{}
	str := ""
	err := c.UnmarshalCSV(str)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), float64(0), c.Value())
}

func (suite *TypesTestSuite) TestCustomFloat64UnmarshalCSVError() {
	c := CustomFloat64{}
	str := "foo"
	err := c.UnmarshalCSV(str)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), `CustomFloat64@UnmarshalCSV Parse float: strconv.ParseFloat: parsing "foo": invalid syntax`, err.Error())
}

func (suite *TypesTestSuite) TestCustomTimestampUnmarshalCSVFilled() {
	c := CustomTimestamp{}
	str := "2020-09-10 15:15:15+0000"
	err := c.UnmarshalCSV(str)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "2020-09-10 15:15:15+0000", c.Value().Format(CustomTimestampFormatDefault))
}

func (suite *TypesTestSuite) TestCustomTimestampUnmarshalCSVEmpty() {
	c := CustomTimestamp{}
	str := ""
	err := c.UnmarshalCSV(str)
	assert.NoError(suite.T(), err)
	assert.True(suite.T(), c.Value().IsZero())
}

func (suite *TypesTestSuite) TestCustomTimestampUnmarshalCSVError() {
	c := CustomTimestamp{}
	str := "foo"
	err := c.UnmarshalCSV(str)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), `CustomTimestamp@UnmarshalJSON ParseTime: parsing time "foo" as "2006-01-02 15:04:05-0700": cannot parse "foo" as "2006"`, err.Error())
}

func (suite *TypesTestSuite) TestCustomTimestampMarshalJSONSuccess() {
	c := CustomTimestamp{}
	c.Timestamp = time.Date(2020, time.Month(1), 20, 0, 0, 0, 0, time.UTC)
	result, err := c.MarshalJSON()
	assert.Equal(suite.T(), []byte(`"2020-01-20 00:00:00+0000"`), result)
	assert.Nil(suite.T(), err)
}

func (suite *TypesTestSuite) TestCustomTimestampMarshalJSONEmpty() {
	c := CustomTimestamp{}
	result, err := c.MarshalJSON()
	assert.Equal(suite.T(), []byte(`""`), result)
	assert.Nil(suite.T(), err)
}

func (suite *TypesTestSuite) TestCustomBooleanUnmarshalCSVTrue() {
	c := CustomBoolean{}
	str := "true"
	err := c.UnmarshalCSV(str)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), true, c.Value())
}

func (suite *TypesTestSuite) TestCustomBooleanUnmarshalCSVFalse() {
	c := CustomBoolean{}
	str := "false"
	err := c.UnmarshalCSV(str)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), false, c.Value())
}

func (suite *TypesTestSuite) TestCustomBooleanUnmarshalCSVEmpty() {
	c := CustomBoolean{}
	str := ""
	err := c.UnmarshalCSV(str)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), false, c.Value())
}

func (suite *TypesTestSuite) TestCustomBooleanUnmarshalCSVError() {
	c := CustomBoolean{}
	str := "foo"
	err := c.UnmarshalCSV(str)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), false, c.Value())
}

func (suite *TypesTestSuite) TestCustomBooleanMarshalJSONSuccess() {
	c := CustomBoolean{}
	c.Boolean = true
	result, err := c.MarshalJSON()
	assert.Equal(suite.T(), []byte(`true`), result)
	assert.Nil(suite.T(), err)
}

func (suite *TypesTestSuite) TestCustomBooleanMarshalJSONEmpty() {
	c := CustomBoolean{}
	result, err := c.MarshalJSON()
	assert.Equal(suite.T(), []byte(`false`), result)
	assert.Nil(suite.T(), err)
}

func TestTypesTestSuite(t *testing.T) {
	suite.Run(t, new(TypesTestSuite))
}
