package sessionhandler

import (
	"fmt"

	"github.com/cdvelop/model"
)

func (s Session) Submit() {

	const this = "submitLoginForm error "

	s.Log("DATA:", s.Form.FormData)

	s.Log(s.Session_status)

	// s.CreateObjectsInDB(s.Table, true, map[string]string{
	// 	s.Id_session:     "123",
	// 	s.Session_status: "in",
	// 	s.Session_encode: "xx",
	// })

	s.ReadAsyncDataDB(model.ReadParams{
		FROM_TABLE:      s.Table,
		WHERE:           []string{s.Session_status},
		SEARCH_ARGUMENT: "in",
	}, func(r model.ReadResult) {

		if r.Error != "" {
			s.UserMessage(r.Error)
			return
		}

		if len(r.DataString) == 1 {
			s.Log("hay usuario en db local:", r.DataString)

		} else {
			s.Log("no hay usuario en local. enviando data al backend:", s.Form.FormData)

			s.SendOneRequest("POST", "create", s.Form.ObjectName, s.Form.FormData, func(result []map[string]string, err string) {

				if err != "" {
					s.UserMessage(err)
					return
				}

				s.Log("RESULTADO SESIÓN:", result)

				if len(result) != 1 {
					s.UserMessage("error se esperaba data para inicio de sesión")
					return
				}

				// SI NO HAY DATA DE ARRANQUE NO DETENGO EL FLUJO
				err = s.FrontendLoadBootData(result[0]["boot"])
				if err != "" {
					s.Log(err)
				}

				// DECODIFICAMOS LA SESIÓN PARA ALMACENARLA
				var session SessionStore

				err = s.DecodeStruct([]byte(result[0]["session"]), &session)
				if err != "" {
					s.UserMessage(this + err)
					return
				}

				fmt.Println("LA SESSION:", session.Session_status)

				// ALMACENAMOS LA SESIÓN EN LA DB DEL NAVEGADOR
				err = s.CreateObjectsInDB(s.Table, false, map[string]string{
					s.Id_session:     session.Id_session,
					s.Session_status: session.Session_status,
					s.Session_encode: session.Session_encode,
				})
				if err != "" {
					s.UserMessage(this + err)
					return
				}

				// EJECUTAMOS LA CONSTRUCCIÓN DE LA UI
				err = s.BuildFrontendUI()
				if err != "" {
					s.UserMessage(this + err)
					return
				}
				//

			})

		}

	})
	// form_name := f.html_form.Get("name").String()

}
