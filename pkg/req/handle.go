package req

import (
	"go/adv-demo/pkg/res"
	"net/http"
)

func HandleBody[T any](w *http.ResponseWriter, req *http.Request) (*T, error) {

	body, err := Decode[T](req.Body)
	if err != nil {
		res.Json(*w, err.Error(), 402)
		return nil, err
	}

	err = IsValid(body)
	if err != nil {
		res.Json(*w, err.Error(), 402)
		return nil, err
	}

	return &body, nil

	//reg, _ := regexp.Compile(`[A-Za-z0-9._%+\-]+@[A-Za-z0-9._%+\-]+\.[a-z]{2,}`)
	//if !reg.MatchString(payload.Email) {
	//	res.Json(writer, "Wrong email", 403)
	//	return
	//}
}
