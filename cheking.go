package sessionhandler

import "github.com/cdvelop/model"

func (s Session) Checking(u *model.User, params []map[string]string) (user_db map[string]string, err string) {
	const this = "Checking "

	if len(params) != 1 {
		return nil, this + "el numero de parámetros recibidos es incorrecto"
	}

	var where = map[string]string{}

	for _, data := range params {

		if s.FieldUser != nil {
			// s.Log("s.FieldUser.Name", s.FieldUser.Name)
			if value, exist := data[s.FieldUser.Name]; exist && value != "" {
				where[s.field_user] = value
			}
		}

		if s.FieldPassword != nil {
			// s.Log("s.FieldPassword.Name", s.FieldPassword.Name)
			if value, exist := data[s.FieldPassword.Name]; exist && value != "" {
				where[s.field_password] = value
			}
		}
	}

	data_db, err := s.ReadSyncDataDB(&model.ReadParams{
		FROM_TABLE:    s.UserTableName,
		WHERE:         []map[string]string{where},
		AND_CONDITION: true})
	if err != "" {
		return nil, this + err
	}

	if len(data_db) != 1 {
		return nil, this + "credenciales de acceso incorrectas"
	}

	return data_db[0], ""

}
