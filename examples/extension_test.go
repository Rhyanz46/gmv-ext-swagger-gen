package examples

import (
	"github.com/Rhyanz46/gmv-ext-swagger-gen/gmv_swagger_gen"
	"github.com/Rhyanz46/go-map-validator/map_validator"
	"testing"
)

func TestExtension(t *testing.T) {
	payload := map[string]interface{}{"hp": "+62567888", "email": "dev@ariansaputra.com"}
	validRole := map[string]map_validator.Rules{
		"hp":    {RegexString: `^\+(?:\d{2}[- ]?\d{6}|\d{11})$`},
		"email": {RegexString: `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`},
	}
	swaggerExt := gmv_swagger_gen.GMVSwaggerGenExt().GenerateSwagger(gmv_swagger_gen.SwaggerRules{
		Endpoint:    "/hello",
		HeaderRoles: nil,
		Method:      "",
		Description: "",
	})

	check, err := map_validator.
		NewValidateBuilder().
		AddExtension(swaggerExt).
		SetRules(validRole).
		Load(payload)
	if err != nil {
		t.Errorf("Expected not have error, but got error : %s", err)
	}
	_, err = check.RunValidate()
	if err != nil {
		t.Errorf("Expected not have error, but got error : %s", err)
	}
}

func TestInvalidExtension(t *testing.T) {
	payload := map[string]interface{}{"hp": "+62567888", "email": "dev@ariansaputra.com"}
	validRole := map[string]map_validator.Rules{
		"hp":    {RegexString: `^\+(?:\d{2}[- ]?\d{6}|\d{11})$`},
		"email": {RegexString: `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`},
	}
	swaggerExt := gmv_swagger_gen.GMVSwaggerGenExt()

	check, err := map_validator.
		NewValidateBuilder().
		AddExtension(swaggerExt).
		SetRules(validRole).
		Load(payload)
	if err != nil {
		t.Errorf("Expected not have error, but got error : %s", err)
	}
	res, err := check.RunValidate()
	if err != nil {
		t.Errorf("Expected not have error, but got error : %s", err)
	}
	if payload["hp"] != res.GetData()["hp"] {
		t.Errorf("Expected %s data, but we got : %s", payload["hp"], res.GetData()["hp"])
	}
}
