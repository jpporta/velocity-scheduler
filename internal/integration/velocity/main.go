package velocity

import "github.com/jpporta/velocity-classes-scheduler/internal/models"

const baseURL = "https://studiovelocity.com.br/api/v1"

type Velocity struct {

	token string
	user models.User
}

func Login(user models.User) Velocity {
	v := Velocity{
		token: "",
		user:	user,
	}

	v.getToken()

	return v
}


