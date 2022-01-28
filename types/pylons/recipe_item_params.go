package types

import (
	"fmt"
	"math"
	"reflect"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type StringKeyValueList []StringKeyValue
type DoubleKeyValueList []DoubleKeyValue
type LongKeyValueList []LongKeyValue

type LongParamList []LongParam
type DoubleParamList []DoubleParam
type StringParamList []StringParam

type DoubleInputParamList []DoubleInputParam
type LongInputParamList []LongInputParam
type StringInputParamList []StringInputParam
type WeightedOutputsList []WeightedOutputs

type ItemList []Item
type ItemInputList []ItemInput

type DoubleWeightTable []DoubleWeightRange
type IntWeightTable []IntWeightRange

var floatType = reflect.TypeOf(float64(0))
var stringType = reflect.TypeOf("")

func getFloat(unk interface{}) (float64, error) {
	switch i := unk.(type) {
	case float64:
		return i, nil
	case float32:
		return float64(i), nil
	case int64:
		return float64(i), nil
	case int32:
		return float64(i), nil
	case int:
		return float64(i), nil
	case uint64:
		return float64(i), nil
	case uint32:
		return float64(i), nil
	case uint:
		return float64(i), nil
	case string:
		return strconv.ParseFloat(i, 64)
	default:
		v := reflect.ValueOf(unk)
		v = reflect.Indirect(v)
		// nolint: gocritic
		if v.Type().ConvertibleTo(floatType) {
			fv := v.Convert(floatType)
			return fv.Float(), nil
		} else if v.Type().ConvertibleTo(stringType) {
			sv := v.Convert(stringType)
			s := sv.String()
			return strconv.ParseFloat(s, 64)
		} else {
			return math.NaN(), fmt.Errorf("cannot convert type %v to float64", v.Type())
		}
	}
}

// Actualize creates a (key, value) list from ParamList
func (dpm DoubleParamList) Actualize(ec CelEnvCollection) (DoubleKeyValueList, error) {
	m := make([]DoubleKeyValue, 0, len(dpm))
	for _, param := range dpm {
		var valDec sdk.Dec
		var err error

		if len(param.Program) > 0 {
			var val float64
			val, err = ec.EvalFloat64(param.Program)
			valDec, _ = sdk.NewDecFromStr(fmt.Sprintf("%v", val))
		} else {
			valDec, err = DoubleWeightTable(param.WeightRanges).Generate()
		}
		if err != nil {
			return m, err
		}
		m = append(m, DoubleKeyValue{
			Key:   param.Key,
			Value: valDec,
		})
	}
	return m, nil
}

// Actualize builds the params
func (lpm LongParamList) Actualize(ec CelEnvCollection) (LongKeyValueList, error) {
	m := make([]LongKeyValue, 0, len(lpm))
	for _, param := range lpm {
		var val int64
		var err error

		if len(param.Program) > 0 {
			val, err = ec.EvalInt64(param.Program)
		} else {
			val, err = IntWeightTable(param.WeightRanges).Generate()
		}
		if err != nil {
			return m, err
		}
		m = append(m, LongKeyValue{
			Key:   param.Key,
			Value: val,
		})
	}
	return m, nil
}

// Actualize actualize string param using cel program
func (spm StringParamList) Actualize(ec CelEnvCollection) (StringKeyValueList, error) {
	m := make([]StringKeyValue, 0, len(spm))
	for _, param := range spm {
		var val string
		var err error

		if len(param.Program) > 0 {
			val, err = ec.EvalString(param.Program)
		} else {
			val = param.Value
		}
		if err != nil {
			return m, err
		}
		m = append(m, StringKeyValue{
			Key:   param.Key,
			Value: val,
		})
	}
	return m, nil
}

// Has check if an input is between double input param range
func (dp DoubleInputParam) Has(input sdk.Dec) bool {
	return input.GTE(dp.MinValue) && input.LTE(dp.MaxValue)
}

// Actualize creates a (key, value) list from ParamList
func (dpm DoubleInputParamList) Actualize() []DoubleKeyValue {
	m := make([]DoubleKeyValue, len(dpm))
	for i, param := range dpm {
		m[i] = DoubleKeyValue{
			Key:   param.Key,
			Value: param.MinValue.Add(param.MaxValue).QuoInt64(2),
		}
	}
	return m
}

// Has validate if input is between min max range
func (lp LongInputParam) Has(input int) bool {
	return int64(input) >= lp.MinValue && int64(input) <= lp.MaxValue
}

// Actualize generate a value from range
func (lpm LongInputParamList) Actualize() []LongKeyValue {
	// We don't have the ability to do random numbers in a verifiable way rn, so don't worry about it
	m := make([]LongKeyValue, len(lpm))
	for i, param := range lpm {
		m[i] = LongKeyValue{
			Key:   param.Key,
			Value: (param.MinValue + param.MaxValue) / 2,
		}
	}
	return m
}

// Actualize actualize string from StringInputParamList
func (lpm StringInputParamList) Actualize() []StringKeyValue {
	// We don't have the ability to do random numbers in a verifiable way rn, so don't worry about it
	m := make([]StringKeyValue, len(lpm))
	for i, param := range lpm {
		m[i] = StringKeyValue(param)
	}
	return m
}
