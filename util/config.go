package util

import (
    "os"
    "strconv"
    "fmt"
    "strings"
    "io/ioutil"
)

// GetStringParameter get a string parameter from secrets or environment variable
func GetStringParameter(name, defaultValue string) string {
    return getParameter(name, defaultValue)
}

// GetBoolParameter get a boolean parameter from secrets or environment variable
func GetBoolParameter(name string, defaultValue bool) bool {
    parameter := getParameter(name, defaultValue)
    value, err := strconv.ParseBool(parameter)

    if nil != err {
        Log.Error(err)
    }

    return value
}

// GetIntParameter get a integer parameter from secrets or environment variable
func GetIntParameter(name string, defaultValue int) int {
    parameter := getParameter(name, defaultValue)
    value, err := strconv.ParseInt(parameter, 10, 0)

    if nil != err {
        Log.Error(err)
    }

    return int(value)
}

// GetFloatParameter get a float parameter from secrets or environment variable
func GetFloatParameter(name string, defaultValue float32) float32 {
    parameter := getParameter(name, defaultValue)
    value, err := strconv.ParseFloat(parameter, 0)

    if nil != err {
        Log.Error(err)
    }

    return float32(value)
}

func getParameter(name string, defaultValue interface{}) string {
    var parameter interface{}
    parameter = getParameterFromSecrets(name)

    if "" == parameter {
        parameter = getParameterFromEnv(name)

        if "" == parameter {
            parameter = defaultValue
        }
    }

    return fmt.Sprintf("%v", parameter)
}

func getParameterFromSecrets(name string) string {
    name = strings.ToLower(name)
    path := fmt.Sprintf("/run/secrets/%s", name)
    secret, _ := ioutil.ReadFile(path)

    return string(secret)
}

func getParameterFromEnv(name string) string {
    name = strings.ToUpper(name)

    return os.Getenv(name)
}
