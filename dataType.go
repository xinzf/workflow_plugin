package workflow_plugin

import (
	"github.com/xinzf/dataframe/series"
)

type DataType string

const (
	// INT =================================
	// integer
	// =================================
	INT       DataType = "INT"
	TINYINT   DataType = "TINYINT"
	MEDIUMINT DataType = "MEDIUMINT"
	SMALLINT  DataType = "SMALLINT"
	BIGINT    DataType = "BIGINT"

	// DECIMAL =================================
	// float
	// =================================
	DECIMAL DataType = "DECIMAL"

	// DATETIME =================================
	// datetime
	// =================================
	DATETIME DataType = "DATETIME"

	// TEXT =================================
	// string
	// =================================
	TEXT       DataType = "TEXT"
	LONGTEXT   DataType = "LONGTEXT"
	MEDIUMTEXT DataType = "MEDIUMTEXT"
	TINYTEXT   DataType = "TINYTEXT"
	VARCHAR    DataType = "VARCHAR"
	CHAR       DataType = "CHAR"
	JSON       DataType = "JSON"
)

var dataTypes = []DataType{
	INT, TINYINT, MEDIUMINT, SMALLINT, BIGINT, DECIMAL,
	DATETIME, TEXT, LONGTEXT, MEDIUMTEXT, TINYTEXT, VARCHAR, CHAR, JSON,
}

func (d DataType) Is() bool {
	if len(string(d)) == 0 {
		return false
	}
	for _, dt := range dataTypes {
		if dt == d {
			return true
		}
	}
	return false
}

func (d DataType) String() string {
	return string(d)
}

//
//func (d DataType) ParseValue(val interface{}) (interface{}, error) {
//    switch d {
//    case INT, TINYINT, MEDIUMINT, SMALLINT:
//        return cast.ToIntE(val)
//    case BIGINT:
//        return cast.ToInt64E(val)
//    case DECIMAL:
//        return cast.ToFloat64E(val)
//    case TEXT, VARCHAR, LONGTEXT, TINYTEXT, CHAR, DATETIME:
//        return cast.ToStringE(val)
//    case JSON:
//        return cast.ToStringMapE(val)
//    default:
//        return nil, fmt.Errorf("Unknow dataType: %s\n", d.String())
//    }
//}

func (d DataType) SeriesType() series.Type {
	switch d {
	case INT, TINYINT, MEDIUMINT, SMALLINT, BIGINT:
		return series.Int
	case DECIMAL:
		return series.Float
	case TEXT, VARCHAR, LONGTEXT, TINYTEXT, CHAR:
		return series.String
	case DATETIME:
		return series.Time
	default:
		return series.String
	}
}
