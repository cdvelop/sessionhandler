package sessionhandler

import "github.com/cdvelop/model"

type Session struct {
	*model.Object
	Form *model.Object
	SessionStore
	Config
}

type SessionStore struct {
	Id_session     string `json:"i,omitempty" Legend:"Id"`
	Session_status string `json:"s,omitempty" Legend:"Estado"`
	Session_encode string `json:"e,omitempty" Legend:"Sesión"`
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
