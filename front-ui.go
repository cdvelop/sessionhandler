package sessionhandler

import "github.com/cdvelop/model"

func (s *Session) UserInterface(u *model.User) string {

	c := &model.TemplateModuleConfig{
		RenderAllSpaceCentered: true,
		Module:                 s.Object.Module,
		Form:                   s.Form,
		FormButtons: []*model.ButtonForm{
			{
				ObjectName: s.Form.ObjectName,
				ButtonName: "btn_" + s.Form.ObjectName,
				Title:      s.Title,
				IconID:     "icon-key",
				OnclickFun: "submitLoginForm(this)",
				Disabled:   true,
			},
		},
		AsideList: nil,
	}

	return s.ModuleTemplate(c)
}
