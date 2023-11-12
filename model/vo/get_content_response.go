package vo

import "time"

type GetContentResponse struct {
	Id      uint      `json:"id"`
	Created time.Time `json:"created"`
	Name    string    `json:"name"`
}
