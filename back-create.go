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

	data_db, err := s.ReadSyncDataDB(s.UserTableName, map[string]string{"WHERE": where})
	if err != "" {
		return this + err
	}

	if len(data_db) != 1 {
		err = this + "credenciales de acceso incorrectas"
		return
	}

	fmt.Println("DATA RECIBIDA:", params)

	// for k, v := range res[0] {
	// params[0][k] = v
	// }
	// fmt.Println("RESULTADO CONSULTA:", params)

	// 1- CREAMOS EL OBJETO USUARIO

	new_user := model.User{
		Token:          token.BuildUniqueKey(16),
		Id:             data_db[0][s.FieldID],
		Ip:             u.Ip,
		Name:           data_db[0][s.FieldName],
		Area:           data_db[0][s.FieldArea],
		AccessLevel:    data_db[0][s.FieldAccessLevel],
		LastConnection: s.ToDay("2006-01-02 15:04:05"),
	}

	// fmt.Println("\nUSUARIO:", new_user)

	// 2- CONVERTIMOS LA DATA EN BYTES JSON
	encode_user, err := s.EncodeStruct(new_user)
	if err != "" {
		return this + err
	}

	//3- CIFRAMOS LA DATA DEL USUARIO
	session_encode, err := s.CipherAdapter.Encrypt(encode_user)
	if err != "" {
		return this + err
	}

	//4- CREAMOS EL OBJETO SESIÓN DEL LADO DEL CLIENTE
	new_session := SessionStore{
		Id_session:     new_user.Id,
		Session_status: "in",
		Session_encode: session_encode,
	}

	//5- CONVERTIMOS A JSON LA SESIÓN
	encode_session, err := s.EncodeStruct(new_session)
	if err != "" {
		return this + err
	}

	//6- CREAMOS UN NUEVO MAPA CON LA NUEVA SALIDA DE INFORMACIÓN
	response := map[string]string{
		"session": string(encode_session),
	}

	fmt.Println("\nnew_user:", new_user)

	out, err := s.BackendLoadBootData(&new_user)
	if err != "" {

		return this + err
	}
	response["boot"] = out

	//7- REMPLAZAMOS EL PRIMER ELEMENTO CON LA NUEVA INFORMACIÓN
	params[0] = response

	// fmt.Println("DATA ENVIADA:", params)

	return
}
