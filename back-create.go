package sessionhandler

import (
	"fmt"

	"github.com/cdvelop/model"
)

func (s Session) Create(u *model.User, params ...map[string]string) (err string) {

	s.Object.Log("CREATE SESSION:", params)

	fmt.Println("PARAMETROSS RECIBIDOS:", params)
	return
}
