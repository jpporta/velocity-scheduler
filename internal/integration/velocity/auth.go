package velocity

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func (v *Velocity) getToken() {
	body := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{
		v.user.Email,
		v.user.Password,
	}
	out, err := json.Marshal(body)
	if err != nil {
		log.Fatalln("Failed to build login body", err)
	}
	res, err := http.Post(baseURL+"/jwt/token-auth/", "application/json", bytes.NewBuffer(out))
	if err != nil {
		log.Fatalln("Failed to login", err)
	}
	defer res.Body.Close()
	in, err := io.ReadAll(res.Body)

	var response struct{ Token string `json:"token"`}

	err = json.Unmarshal(in, &response)
	if err != nil {
		log.Fatalln("Failed to unmarshal response", err)
	}
	v.token = response.Token
}
