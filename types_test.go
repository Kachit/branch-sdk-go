package branchio

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type TypesCustomIntegerTestSuite struct {
	suite.Suite
	testable *CustomInteger
}

func (suite *TypesCustomIntegerTestSuite) SetupTest() {
	suite.testable = &CustomInteger{}
}

func (suite *TypesCustomIntegerTestSuite) TestMarshalJSONSuccess() {
	suite.testable.Integer = 10
	result, err := suite.testable.MarshalJSON()
	expected := []byte(`10`)
	assert.Equal(suite.T(), expected, result)
	assert.Nil(suite.T(), err)
}

func (suite *TypesCustomIntegerTestSuite) TestMarshalJSONEmpty() {
	result, err := suite.testable.MarshalJSON()
	expected := []byte(`0`)
	assert.Equal(suite.T(), expected, result)
	assert.Nil(suite.T(), err)
}

func (suite *TypesCustomIntegerTestSuite) TestUnmarshalCSVFilled() {
	str := "12345"
	err := suite.testable.UnmarshalCSV(str)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 12345, suite.testable.Value())
}

func (suite *TypesCustomIntegerTestSuite) TestUnmarshalCSVEmpty() {
	str := ""
	err := suite.testable.UnmarshalCSV(str)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 0, suite.testable.Value())
}

func (suite *TypesCustomIntegerTestSuite) TestUnmarshalCSVError() {
	str := "foo"
	err := suite.testable.UnmarshalCSV(str)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), `CustomInteger@UnmarshalCSV Parse int: strconv.Atoi: parsing "foo": invalid syntax`, err.Error())
}

type TypesCustomFloat64TestSuite struct {
	suite.Suite
	testable *CustomFloat64
}

func (suite *TypesCustomFloat64TestSuite) SetupTest() {
	suite.testable = &CustomFloat64{}
}

func (suite *TypesCustomFloat64TestSuite) TestMarshalJSONSuccess() {
	suite.testable.Float64 = 10.10
	result, err := suite.testable.MarshalJSON()
	expected := []byte(`10.1`)
	assert.Equal(suite.T(), expected, result)
	assert.Nil(suite.T(), err)
}

func (suite *TypesCustomFloat64TestSuite) TestMarshalJSONEmpty() {
	suite.testable.Float64 = 10.10
	result, err := suite.testable.MarshalJSON()
	expected := []byte(`10.1`)
	assert.Equal(suite.T(), expected, result)
	assert.Nil(suite.T(), err)
}

func (suite *TypesCustomFloat64TestSuite) TestUnmarshalCSVFilled() {
	str := "123.45"
	err := suite.testable.UnmarshalCSV(str)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 123.44999694824219, suite.testable.Value())
}

func (suite *TypesCustomFloat64TestSuite) TestUnmarshalCSVEmpty() {
	str := ""
	err := suite.testable.UnmarshalCSV(str)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), float64(0), suite.testable.Value())
}

func (suite *TypesCustomFloat64TestSuite) TestUnmarshalCSVError() {
	str := "foo"
	err := suite.testable.UnmarshalCSV(str)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), `CustomFloat64@UnmarshalCSV Parse float: strconv.ParseFloat: parsing "foo": invalid syntax`, err.Error())
}

type TypesCustomTimestampTestSuite struct {
	suite.Suite
	testable *CustomTimestamp
}

func (suite *TypesCustomTimestampTestSuite) SetupTest() {
	suite.testable = &CustomTimestamp{}
}

func (suite *TypesCustomTimestampTestSuite) TestUnmarshalCSVFilled() {
	str := "2020-09-10 15:15:15+0000"
	err := suite.testable.UnmarshalCSV(str)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "2020-09-10 15:15:15+0000", suite.testable.Value().Format(CustomTimestampFormatDefault))
}

func (suite *TypesCustomTimestampTestSuite) TestUnmarshalCSVEmpty() {
	str := ""
	err := suite.testable.UnmarshalCSV(str)
	assert.NoError(suite.T(), err)
	assert.True(suite.T(), suite.testable.Value().IsZero())
}

func (suite *TypesCustomTimestampTestSuite) TestUnmarshalCSVError() {
	str := "foo"
	err := suite.testable.UnmarshalCSV(str)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), `CustomTimestamp@UnmarshalJSON ParseTime: parsing time "foo" as "2006-01-02 15:04:05-0700": cannot parse "foo" as "2006"`, err.Error())
}

func (suite *TypesCustomTimestampTestSuite) TestMarshalJSONSuccess() {
	suite.testable.Timestamp = time.Date(2020, time.Month(1), 20, 0, 0, 0, 0, time.UTC)
	result, err := suite.testable.MarshalJSON()
	assert.Equal(suite.T(), []byte(`"2020-01-20 00:00:00+0000"`), result)
	assert.Nil(suite.T(), err)
}

func (suite *TypesCustomTimestampTestSuite) TestMarshalJSONEmpty() {
	result, err := suite.testable.MarshalJSON()
	assert.Equal(suite.T(), []byte(`""`), result)
	assert.Nil(suite.T(), err)
}

type TypesCustomBooleanTestSuite struct {
	suite.Suite
	testable *CustomBoolean
}

func (suite *TypesCustomBooleanTestSuite) SetupTest() {
	suite.testable = &CustomBoolean{}
}

func (suite *TypesCustomBooleanTestSuite) TestUnmarshalCSVTrue() {
	str := "true"
	err := suite.testable.UnmarshalCSV(str)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), true, suite.testable.Value())
}

func (suite *TypesCustomBooleanTestSuite) TestUnmarshalCSVFalse() {
	str := "false"
	err := suite.testable.UnmarshalCSV(str)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), false, suite.testable.Value())
}

func (suite *TypesCustomBooleanTestSuite) TestUnmarshalCSVEmpty() {
	str := ""
	err := suite.testable.UnmarshalCSV(str)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), false, suite.testable.Value())
}

func (suite *TypesCustomBooleanTestSuite) TestUnmarshalCSVError() {
	str := "foo"
	err := suite.testable.UnmarshalCSV(str)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), false, suite.testable.Value())
}

func (suite *TypesCustomBooleanTestSuite) TestMarshalJSONSuccess() {
	suite.testable.Boolean = true
	result, err := suite.testable.MarshalJSON()
	assert.Equal(suite.T(), []byte(`true`), result)
	assert.Nil(suite.T(), err)
}

func (suite *TypesCustomBooleanTestSuite) TestMarshalJSONEmpty() {
	result, err := suite.testable.MarshalJSON()
	assert.Equal(suite.T(), []byte(`false`), result)
	assert.Nil(suite.T(), err)
}

func TestTypesCustomIntegerTestSuite(t *testing.T) {
	suite.Run(t, new(TypesCustomIntegerTestSuite))
}

func TestTypesCustomFloat64TestSuite(t *testing.T) {
	suite.Run(t, new(TypesCustomFloat64TestSuite))
}

func TestTypesCustomTimestampTestSuite(t *testing.T) {
	suite.Run(t, new(TypesCustomTimestampTestSuite))
}

func TestTypesCustomBooleanTestSuite(t *testing.T) {
	suite.Run(t, new(TypesCustomBooleanTestSuite))
}
