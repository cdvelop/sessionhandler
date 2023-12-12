package sessionhandler

import (
	"github.com/cdvelop/model"
	"github.com/cdvelop/object"
)

const MODULE_NAME = "user"
const OBJECT_LOGIN = "login"
const TABLE_NAME = "session"

func Add(h *model.Handlers, c *Config) (s *Session, err string) {

	s = &Session{}

	m := &model.Module{
		ModuleName: MODULE_NAME,
		Title:      "Ingreso",
		IconID:     "icon-home",
		UI:         s,
		Areas:      map[string]string{},
		Objects:    []*model.Object{},
		Inputs:     []*model.Input{},
		Handlers:   h,
	}

	err = object.AddToHandlerFromStructs(m, s, h)
	if err != "" {
		return
	}

	s.HeaderAuthName = "Authorization"

	s.Config = c

	var fields []model.Field
	if c.FieldUser != nil {
		// fmt.Println("** c.FieldUser.Name:", c.FieldUser.Name)
		s.field_user = c.FieldUser.Name

		s.FieldUser = &model.Field{
			Name:                     "user",
			Legend:                   c.FieldUser.Legend,
			Input:                    c.FieldUser.Input,
			NotRequiredInDB:          true,
			Encrypted:                true,
			NotClearValueOnFormReset: true,
		}

		fields = append(fields, *s.FieldUser)

	}
	if c.FieldPassword != nil {
		// fmt.Println("** c.FieldPassword.Name:", c.FieldPassword.Name)
		s.field_password = c.FieldPassword.Name

		s.FieldPassword = &model.Field{
			Name:                     "password",
			Legend:                   c.FieldPassword.Legend,
			Input:                    c.FieldPassword.Input,
			NotRequiredInDB:          true,
			Encrypted:                true,
			NotClearValueOnFormReset: true,
		}

		fields = append(fields, *s.FieldPassword)
	}

	s.Form = &model.Object{
		ObjectName:      OBJECT_LOGIN,
		Table:           TABLE_NAME,
		NoAddObjectInDB: true,
		Fields:          fields,
		Module:          m,
		BackHandler: model.BackendHandler{
			BootResponse: nil,
			CreateApi:    nil,
			ReadApi:      nil,
			UpdateApi:    nil,
			DeleteApi:    nil,
		},
		FrontHandler: model.FrontendHandler{},
		FormData:     map[string]string{},
		// AlternativeValidateAdapter: s,
	}

	h.AddObjects(s.Form)

	return s, ""
}

func (s Session) NameOfAuthHandler() string {
	return OBJECT_LOGIN
}
