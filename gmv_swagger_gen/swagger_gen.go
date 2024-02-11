package gmv_swagger_gen

import "github.com/Rhyanz46/go-map-validator/map_validator"

type extBase struct {
	rules       *map[string]map_validator.Rules
	headerRules *map[string]map_validator.Rules
	data        interface{}
	extraData   *map_validator.ExtraOperationData
	dontUseBody bool
}

type SwaggerMeta struct {
	BaseUrl     string
	Description string
	Contact     string
	SwaggerURL  string
}

type SwaggerRules struct {
	Endpoint    string
	HeaderRoles *map[string]map_validator.Rules
	Method      string
	Description string
}

func (e *extBase) SetRoles(rules *map[string]map_validator.Rules) {
	e.rules = rules
}

func (e *extBase) BeforeLoad(data interface{}) error {
	//TODO implement me
	//panic("implement me")
	return nil
}

func (e *extBase) AfterLoad(data *map[string]interface{}) error {
	//TODO implement me
	//panic("implement me")
	return nil
}

func (e *extBase) BeforeValidation(data *map[string]interface{}) error {
	//TODO implement me
	//panic("implement me")
	return nil
}

func (e *extBase) AfterValidation(data *map[string]interface{}) error {
	return nil
}

func (e *extBase) SetExtraData(data *map_validator.ExtraOperationData) map_validator.ExtensionType {
	e.extraData = data
	return e
}

func (e *extBase) DontUseBody() *extBase {
	e.dontUseBody = true
	return e
}

func (e *extBase) GenerateSwagger(meta SwaggerRules) *extBase {
	return e
}

func GMVSwaggerGenExt() *extBase {
	return &extBase{}
}
