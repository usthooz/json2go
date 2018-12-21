package json2go


type Data struct { 
	B float64 `json:"b"` 
	A float64 `json:"a"` 
}
 

type Json2go struct { 
	A string `json:"a"` 
	B string `json:"b"` 
	Data Data `json:"data"` 
}
 

type Json2GoAutoGenerate struct {
	Json2go Json2go `json:"json2go"`
}

