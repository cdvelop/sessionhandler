package sessionhandler

import "github.com/cdvelop/model"

type Session struct {
	*model.Object
	Form           *model.Object
	Id_session     string `Legend:"Id"`
	Session_status string `Legend:"Estado"`
	Session_encode string `Legend:"Sesión"`
}
