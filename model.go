package sessionhandler

import "github.com/cdvelop/model"

type Session struct {
	*model.Object
	Form           *model.Object
	Id_session     string `Legend:"Id"`
	Session_status string `Legend:"Estado"`
	Session_encode string `Legend:"Sesión"`

	Config
}

type Config struct {
	UserTableName string // ej: staff,user,client
	FieldUser     *model.Field
	FieldPassword *model.Field

	FieldID          string // ej: id_staff, id_user
	FieldName        string //ej: staff_name, user_name
	FieldArea        string //ej: staff_area, user_area
	FieldAccessLevel string // ej: staff_credentials, user_level

	field_user     string
	field_password string
}
