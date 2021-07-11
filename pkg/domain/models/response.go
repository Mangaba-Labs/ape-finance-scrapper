package models

// Response is a generic struct to our services
type Response struct {
	HttpCode int
	Status   string
	Message  string
}

// Set is our constructor to Response
func (c *Response) Set(httpCode int, status string, message string) {
	c.HttpCode = httpCode
	c.Message = message
	c.Status = status
}
