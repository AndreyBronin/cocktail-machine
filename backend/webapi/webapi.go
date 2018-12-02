package webapi

type Ingredient struct {
	Name string
}

type Recipe struct {
	Ingredients map[string]uint8
}

// WebApi is backend REST API
type WebApi interface {
	MakeCocktail(contract string)
}

type WebAPIComponent struct {
	Port int
	Host int
}

func (wa WebAPIComponent) Start() {

}
