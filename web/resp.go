package web

type ErrResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (e ErrResp) Error() string {
	return e.Msg
}

var ErrInvalidToken = ErrResp{Code: 41001, Msg: "invalid token"}
