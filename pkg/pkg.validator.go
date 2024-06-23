package pkg

import (
	"github.com/arioprima/cari_kampus_api/schemas"
	"github.com/go-playground/validator/v10"
)

func ValidatorLogin(s interface{}, config []schemas.ErrorMetaConfig) (interface{}, int) {
	v := validator.New()
	errResponse := make(map[string]string)
	errCount := 0

	// Validasi setiap konfigurasi error
	for _, cfg := range config {
		switch cfg.Tag {
		case "required":
			if err := v.Var(s, "required"); err != nil {
				errResponse[cfg.Field] = cfg.Message
				errCount++
			}
		case "email":
			if err := v.Var(s, "email"); err != nil {
				errResponse[cfg.Field] = cfg.Message
				errCount++
			}
		case "min":
			if err := v.Var(s, "min="+cfg.Value); err != nil {
				errResponse[cfg.Field] = cfg.Message
				errCount++
			}
		}
	}

	return errResponse, errCount
}
