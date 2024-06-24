package pkg

import (
	"log"
	"reflect"

	"github.com/arioprima/cari_kampus_api/schemas"
	"github.com/go-playground/validator/v10"
)

func ValidatorLogin(s interface{}, config []schemas.ErrorMetaConfig) (map[string]string, int) {
	v := validator.New()
	errResponse := make(map[string]string)
	errCount := 0

	// Validate the individual fields in the struct
	for _, cfg := range config {
		fieldValue, ok := getFieldValue(s, cfg.Field)
		if !ok {
			log.Printf("Field %s not found in struct", cfg.Field)
			continue
		}

		log.Printf("Validating field: %s with value: %v", cfg.Field, fieldValue)
		if _, exists := errResponse[cfg.Field]; !exists { // Only validate if no error exists for this field
			switch cfg.Tag {
			case "required":
				if err := v.Var(fieldValue, "required"); err != nil {
					errResponse[cfg.Field] = cfg.Message
					errCount++
				}
			case "email":
				if err := v.Var(fieldValue, "email"); err != nil {
					errResponse[cfg.Field] = cfg.Message
					errCount++
				}
			case "min":
				if err := v.Var(fieldValue, "min="+cfg.Value); err != nil {
					errResponse[cfg.Field] = cfg.Message
					errCount++
				}
			}
		}
	}

	return errResponse, errCount
}

// Helper function to get the field value from the struct
func getFieldValue(s interface{}, field string) (interface{}, bool) {
	r := reflect.ValueOf(s)
	f := reflect.Indirect(r).FieldByName(field)
	if !f.IsValid() {
		return nil, false
	}
	return f.Interface(), true
}
