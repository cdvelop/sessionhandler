package sessionhandler

import (
	"fmt"

	"github.com/cdvelop/model"
	"github.com/cdvelop/token"
)

func (s Session) Create(u *model.User, params ...map[string]string) (err string) {
	const this = "Create session error "

	if len(params) != 1 {
		return this + "el numero de parámetros recibidos es incorrecto"
	}

	var where string
	var and string

	for _, data := range params {

		if s.FieldUser != nil {
			fmt.Println("s.FieldUser.Name", s.FieldUser.Name)
			if value, exist := data[s.FieldUser.Name]; exist && value != "" {

				where += and + s.field_user + ` = '` + value + `'`
				and = ` AND `
			}
		}

		if s.FieldPassword != nil {
			fmt.Println("s.FieldPassword.Name", s.FieldPassword.Name)
			if value, exist := data[s.FieldPassword.Name]; exist && value != "" {

				where += and + s.field_password + ` = '` + value + `'`
				and = ` AND `
			}
		}
	}

	res, err := s.ReadObjectsInDB(s.UserTableName, map[string]string{"WHERE": where})
	if err != "" {
		return this + err
	}

	if len(res) != 1 {
		err = this + "credenciales de acceso incorrectas"
		return
	}

	for k, v := range res[0] {
		params[0][k] = v
	}
	// fmt.Println("RESULTADO CONSULTA:", params)

	new_user := model.User{
		Token:          token.BuildUniqueKey(16),
		Id:             params[0][s.FieldID],
		Ip:             u.Ip,
		Name:           params[0][s.FieldName],
		Area:           params[0][s.FieldArea],
		AccessLevel:    params[0][s.FieldAccessLevel],
		LastConnection: s.ToDay("2006-01-02 15:04:05"),
	}

	fmt.Println("\nnew_user:", new_user)

	out, err := s.BackendLoadBootData(&new_user)
	if err != "" {

		return this + err
	}
	params[0]["boot_data"] = out

	// fmt.Println("\nBOOT DATA:", out)

	return
}
