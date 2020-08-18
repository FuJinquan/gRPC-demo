package to_json

import (
	"client/note"
	"encoding/json"
	"fmt"
)

type Response struct {
	Code    int32  `json:"code"`
	Data    Data   `json:"data"`
	Message string `json:"message"`
}
type Data struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

func OutPut(in *note.Response) string {
	res := Response{
		Code:    in.Code,
		Data:    Data{Title: in.Data.Title, Text: in.Data.Text},
		Message: in.Message,
	}
	resp, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err.Error())
	}
	return string(resp)

}
