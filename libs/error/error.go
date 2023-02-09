package error

type Error struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
