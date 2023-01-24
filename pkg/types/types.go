package types;

type Value struct {
	Value float32 `json:"value"`
	Unit string   `json:"unit"`
}

type Ingredients struct {
	Malt []struct {
		Name string  `json:"name"`
		Amount Value `json:"amount"`
	} 						 `json:"malt"`
	Hops []struct {
		Name string  		 `json:"name"`
		Add string   		 `json:"add"`
		Attribute string `json:"attribute"`
		Amount Value 		 `json:"amount"`
	} 					 `json:"hops"`
	Yeast string `json:"yeast"`
}

type Beer struct {
	Id 							 int32   			`json:"id"`
	Name 						 string  			`json:"name"`
	Ingredients			 Ingredients  `json:"ingredients"`
}