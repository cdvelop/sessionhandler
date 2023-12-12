package sessionhandler

import "github.com/cdvelop/model"

func (s Session) BuildUserFromStoreData(data []map[string]string) (u *model.User, err string) {

	const this = "BuildUserFromStoreData error "
	if len(data) != 1 {
		return nil, this + "se esperaba un resultado"
	}

	// s.Log("info DECODIFICAR USUARIO DATA:", data)

	if encode_user, ok := data[0][s.Session_encode]; ok && encode_user != "" {
		// s.Log("encode_user:", encode_user)
		return s.DecodeUser(encode_user)
		// s.Log("info USUARIO:", u.Name)
	}

	return
}
