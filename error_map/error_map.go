package error_map

import "errors"

func Register(code interface{}, message string) {
	switch MODE {
	case MODE_INT:
		intMap[code.(int)] = message
	case MODE_STR:
		strMap[code.(string)] = message
	}
}

func Message(code interface{}) string {
	switch MODE {
	case MODE_INT:
		return intMap[code.(int)]
	case MODE_STR:
		return strMap[code.(string)]
	}
	return ""
}

func Error(code interface{}) error {
	return errors.New(Message(code))
}
