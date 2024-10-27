package environment

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
	"reflect"
	"strings"
	"time"
)

type Environment struct {
}

func (e *Environment) LoadEnv(config any) error {
	// loading ../../.env when running tests
	godotenv.Load("./.env")
	godotenv.Load("../.env")
	godotenv.Load("../../.env")

	value := reflect.ValueOf(config)
	if value.Kind() != reflect.Ptr || value.Elem().Kind() != reflect.Struct {
		return errors.New("config must be a pointer to a struct")
	}

	value = value.Elem()
	valueType := value.Type()

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		fieldType := valueType.Field(i)
		envTag := fieldType.Tag.Get("env")

		if envTag == "" {
			continue
		}

		envValue, ok := os.LookupEnv(envTag)
		if !ok {
			return errors.New("missing environment variable: " + envTag)
		}

		if !field.CanSet() {
			return errors.New("cannot set field value: " + fieldType.Name)
		}

		switch field.Type().Kind() {
		case reflect.String:
			field.SetString(envValue)

		case reflect.Slice:
			if fieldType.Type.Elem().Kind() != reflect.String {
				return errors.New("slice elements must be of type strings")
			}

			envValues := strings.Split(envValue, ",")
			field.Set(reflect.ValueOf(envValues))

		case reflect.Int64:
			switch field.Type().String() {
			case "time.Duration":
				dur, err := time.ParseDuration(envValue)
				if err != nil {
					return err
				}
				field.Set(reflect.ValueOf(dur))
			}
		}
	}

	return nil
}
