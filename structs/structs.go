package structs
// Response....
type Response struct{
	code int        `json:"code"`
	Body interface{}  `json := "Body"`
}