package models

type Response struct {
	HttpCode int
	Status string
	Message string
}


func (c *Response) Set(httpCode int, status string, message string) {
	c.HttpCode = httpCode
	c.Message = message
	c.Status = status
}