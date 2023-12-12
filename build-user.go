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

func (s Session) DecodeUser(encode_user string) (u *model.User, err string) {

	session_decrypt, err := s.CipherAdapter.Decrypt(encode_user)
	if err != "" {
		return nil, err
	}
	// s.Log("Decrypt:", session_decrypt)

	u = &model.User{}

	err = s.DecodeStruct([]byte(session_decrypt), u)
	if err != "" {
		return nil, err
	}

	return u, ""
}
