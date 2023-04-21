package main

import (
	"fmt"
	"reflect"
	"time"

	"github.com/mitchellh/mapstructure"
)

func TimeToStringHookFunc() mapstructure.DecodeHookFuncType {
	return func(from reflect.Type, to reflect.Type, data any) (any, error) {
		if reflect.DeepEqual(to, reflect.TypeOf(time.Time{})) {
			tt, ok := tryParseTime(data.(string))
			if ok {
				return tt, nil
			}
		}
		return data, nil
	}
}

func main() {
	type Person struct {
		Name   string
		Age    int
		Emails []string
		Extra  map[string]string
		C      time.Time
	}

	input := map[string]interface{}{
		"name":   "Mitchell",
		"age":    91,
		"emails": []string{"one", "two", "three"},
		"extra": map[string]string{
			"twitter": "mitchellh",
		},
		"c": "2006-01-02 15:04:05-07:00",
	}

	var result Person
	err := Decode(input, &result)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v", result)
}

func Decode(input any, result any) error {
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Metadata:   nil,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(TimeToStringHookFunc()),
		Result:     result,
	})
	if err != nil {
		return err
	}
	return decoder.Decode(input)
}

func tryParseTime(candidate string) (time.Time, bool) {
	var ret time.Time
	var found bool
	timeFormats := [...]string{
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		time.Kitchen,
		time.RFC3339,
		time.RFC3339Nano,
		"2006-01-02",                         // RFC 3339
		"2006-01-02 15:04",                   // RFC 3339 with minutes
		"2006-01-02 15:04:05",                // RFC 3339 with seconds
		"2006-01-02 15:04:05-07:00",          // RFC 3339 with seconds and timezone
		"2006-01-02T15Z0700",                 // ISO8601 with hour
		"2006-01-02T15:04Z0700",              // ISO8601 with minutes
		"2006-01-02T15:04:05Z0700",           // ISO8601 with seconds
		"2006-01-02T15:04:05.999999999Z0700", // ISO8601 with nanoseconds
	}

	for _, format := range timeFormats {
		ret, found = tryParseExactTime(candidate, format)
		if found {
			return ret, true
		}
	}
	return time.Now(), false
}

func tryParseExactTime(candidate string, format string) (time.Time, bool) {
	var ret time.Time
	var err error
	ret, err = time.ParseInLocation(format, candidate, time.Local)
	if err != nil {
		return time.Now(), false
	}
	return ret, true
}
