package sessionhandler

import (
	"github.com/cdvelop/input"
	"github.com/cdvelop/model"
	"github.com/cdvelop/object"
)

const MODULE_NAME = "user"
const OBJECT_LOGIN = "login"
const TABLE_NAME = "session"

func Add(h *model.Handlers) (s *Session, err string) {

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

	s.Form = &model.Object{
		ObjectName:      OBJECT_LOGIN,
		Table:           TABLE_NAME,
		NoAddObjectInDB: true,
		Fields: []model.Field{
			// {Name: "user", Legend: "Usuario", Input: input.Mail(), NotClearValueOnFormReset: true},
			{Name: "password", Legend: "Contraseña", Input: input.Rut("hide-typing"), NotClearValueOnFormReset: true},
		},
		Module: m,
		BackHandler: model.BackendHandler{
			BootResponse: nil,
			CreateApi:    s,
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
	return s.Form.ObjectName
}
