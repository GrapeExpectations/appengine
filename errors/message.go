package errors

type message struct {
	Pkg  string `json:"package"`
	Fn   string `json:"function"`
	File string `json:"file"`
	Line int    `json:"line"`
	Msg  string `json:"message"`
}
