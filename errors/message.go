package errors

type Message struct {
	Pkg string `json:"package"`
	Fn  string `json:"function"`
	Msg string `json:"message"`
}
