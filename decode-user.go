package sessionhandler

import "github.com/cdvelop/model"

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
