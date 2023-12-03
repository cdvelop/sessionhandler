package sessionhandler

import "github.com/cdvelop/model"

func (s *Session) UserInterface(opt ...string) string {

	c := &model.TemplateModuleConfig{
		Module:      s.Object.Module,
		Form:        s.Form,
		AsideList:   nil,
		ButtonLogin: true,
	}

	return s.ModuleTemplate(c)
}
