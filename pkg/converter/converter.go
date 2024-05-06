package converter

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"strconv"
	"strings"
)

const (
	MESSAGING_TYPE_XML  = "xml"
	MESSAGING_TYPE_JSON = "json"
	MESSAGING_TYPE_ISO  = "iso"
)

// ToString converts any value to string
func ToString(v interface{}) string {
	result := ""
	if v == nil {
		return ""
	}
	switch val := v.(type) {
	case string:
		result = val
	case int:
		result = strconv.Itoa(val)
	case int32:
		result = strconv.Itoa(int(val))
	case int64:
		result = strconv.FormatInt(val, 10)
	case bool:
		result = strconv.FormatBool(val)
	case float32:
		result = strconv.FormatFloat(float64(val), 'f', -1, 32)
	case float64:
		result = strconv.FormatFloat(val, 'f', -1, 64)
	case []byte:
		result = string(val)
	default:
		resultJSON, err := json.Marshal(val)
		if err == nil {
			result = string(resultJSON)
		} else {
			log.Println("[ToString] marshal error", err)
		}
	}

	return result
}

// ToInt64 convert any value to int64
func ToInt64(v interface{}) int64 {
	result := int64(0)
	switch v.(type) {
	case string:
		str := strings.TrimSpace(v.(string))
		result, _ = strconv.ParseInt(str, 10, 64)
	case int:
		result = int64(v.(int))
	case int32:
		result = int64(v.(int32))
	case int64:
		result = v.(int64)
	case float32:
		result = int64(v.(float32))
	case float64:
		result = int64(v.(float64))
	case []byte:
		var err error
		result, err = strconv.ParseInt(string(v.([]byte)), 10, 64)
		if err != nil {
			result = 0
		}
	case json.RawMessage:
		var num int64
		_ = json.Unmarshal(v.(json.RawMessage), &num)
		result = num
	default:
		result = int64(0)
	}

	return result
}

// ToInt32 converts any value to int32
func ToInt32(v interface{}) int32 {
	temp := int64(0)
	result := int32(0)

	switch val := v.(type) {
	case string:
		str := strings.TrimSpace(val)
		temp, _ = strconv.ParseInt(str, 10, 32)
		result = int32(temp)
	case int:
		result = int32(val)
	case int32:
		result = val
	case int64:
		result = int32(val)
	case float32:
		result = int32(val)
	case float64:
		result = int32(val)
	case []byte:
		temp, _ = strconv.ParseInt(string(val), 10, 32)
		result = int32(temp)
	}

	return result
}

// ToInt converts any value to int
func ToInt(v interface{}) int {
	result := 0
	switch v.(type) {
	case string:
		str := strings.TrimSpace(v.(string))

		// make sure that input is a valid int
		// if input > max int64
		// it's still treated as a valid query on db
		// but it will consume too much db resource
		// because it will scan all the existing rows
		var err error
		result, err = strconv.Atoi(str)
		if err != nil {
			result = 0
		}
	case int:
		result = v.(int)
	case int32:
		result = int(v.(int32))
	case int64:
		result = int(v.(int64))
	case float32:
		result = int(v.(float32))
	case float64:
		result = int(v.(float64))
	case []byte:
		var err error
		result, err = strconv.Atoi(string(v.([]byte)))
		if err != nil {
			result = 0
		}
	case json.RawMessage:
		var num int
		_ = json.Unmarshal(v.(json.RawMessage), &num)
		result = num
	default:
		result = 0
	}

	return result
}

// ToMarshal converts struct from json or xml to string
func ToMarshal(data interface{}, dataType string) string {
	result := ""

	// data already a string
	switch data.(type) {
	case string:
		result = data.(string)
	default:
		// continue
	}

	switch dataType {
	case MESSAGING_TYPE_XML:
		resultXML, err := xml.Marshal(data)
		if err == nil {
			result = string(resultXML)
		} else {
			log.Printf("Error: function ToMarshal, Data: %+v, Type: %s, Error: %+v", data, dataType, err)
		}
	case MESSAGING_TYPE_JSON:
		resultJSON, err := json.Marshal(data)
		if err == nil {
			result = string(resultJSON)
		} else {
			log.Printf("Error: function ToMarshal, Data: %+v, Type: %s, Error: %+v", data, dataType, err)
		}
	case MESSAGING_TYPE_ISO:
		result = fmt.Sprintf("%s", data)
	}

	return result
}

// ToFloat64 convert any value to float64
func ToFloat64(v interface{}) float64 {
	result := float64(0)
	switch v.(type) {
	case string:
		str := strings.TrimSpace(v.(string))
		result, _ = strconv.ParseFloat(str, 64)
	case int:
		result = float64(v.(int))
	case int32:
		result = float64(v.(int32))
	case int64:
		result = float64(v.(int64))
	case float32:
		result = float64(v.(float32))
	case float64:
		result = float64(v.(float64))
	case []byte:
		result, _ = strconv.ParseFloat(string(v.([]byte)), 64)
	case json.RawMessage:
		var num float64
		_ = json.Unmarshal(v.(json.RawMessage), &num)
		result = num
	default:
		result = float64(0)
	}

	return result
}

// ToBool convert any value to boolean
func ToBool(v interface{}) bool {
	var result bool
	switch v.(type) {
	case string:
		str := strings.TrimSpace(v.(string))
		result, _ = strconv.ParseBool(str)
	case int:
		result = (v != 0)
	default:
		// do nothing
	}

	return result
}

// ToArrayOfInt convert any value to []int
func ToArrayOfInt(v interface{}) []int {
	var result []int
	switch v.(type) {
	case string:
		_ = json.Unmarshal([]byte(v.(string)), &result)
	case []string:
		b := v.([]string)
		for _, vv := range b {
			result = append(result, ToInt(vv))
		}
	case [][]byte:

		b := v.([][]byte)
		for _, vv := range b {
			result = append(result, ToInt(vv))
		}
	case []int64:
		b := v.([]int64)
		for _, vv := range b {
			result = append(result, ToInt(vv))
		}
	default:
		// do nothing
	}

	return result
}

// ToArrayOfInt64 convert any value to []int64
func ToArrayOfInt64(v interface{}) []int64 {
	var result []int64
	switch v.(type) {
	case []int:
		b := v.([]int)
		for _, vv := range b {
			result = append(result, ToInt64(vv))
		}
	case string:
		_ = json.Unmarshal([]byte(v.(string)), &result)
	case [][]byte:

		b := v.([][]byte)
		for _, vv := range b {
			result = append(result, ToInt64(vv))
		}
	default:
		// do nothing
	}

	return result
}

// ToArrayOfString convert any value to []string
func ToArrayOfString(v interface{}) []string {
	var result []string
	switch val := v.(type) {
	case string:
		_ = json.Unmarshal([]byte(val), &result)
	case [][]byte:
		result = make([]string, len(val))
		for idx := range val {
			result[idx] = string(val[idx])
		}
	case []int64:
		result = make([]string, len(val))
		for idx := range val {
			result[idx] = strconv.FormatInt(val[idx], 10)
		}
	default:
		// do nothing
	}

	return result
}

// ToByteArr Convert any value to byte array
func ToByteArr(v interface{}) []byte {
	result := []byte("")
	if v == nil {
		return result
	}
	switch v.(type) {
	case string:
		result = []byte(v.(string))
	case int:
		result = []byte(strconv.Itoa(v.(int)))
	case int64:
		result = []byte(strconv.FormatInt(v.(int64), 10))
	case bool:
		result = []byte(strconv.FormatBool(v.(bool)))
	case float64:
		result = []byte(strconv.FormatFloat(v.(float64), 'E', -1, 64))
	case []uint8:
		result = v.([]uint8)
	default:
		log.Println("Unsupported data type")
	}

	return result
}

func ToCurrencyNumber(val float64) string {
	intVal := ToInt64(val)
	strVal := fmt.Sprintf("%.2f", val)
	decimalValue := (ToInt64(string(strVal[len(strVal)-2])) * 10) + ToInt64(string(strVal[len(strVal)-1]))

	sign := ""
	if intVal < 0 {
		sign = "-"
		intVal = 0 - intVal
	}

	parts := []string{"", "", "", "", "", "", ""}
	j := len(parts) - 1

	for intVal > 999 {
		parts[j] = strconv.FormatInt(intVal%1000, 10)
		switch len(parts[j]) {
		case 2:
			parts[j] = "0" + parts[j]
		case 1:
			parts[j] = "00" + parts[j]
		}
		intVal = intVal / 1000
		j--
	}
	parts[j] = strconv.Itoa(int(intVal))
	return fmt.Sprintf("%v.%v", sign+strings.Join(parts[j:], ","), decimalValue)
}
