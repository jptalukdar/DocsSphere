package response

type Response struct {
	ReturnCode int
	Msg        string
	Origin     string
	ErrorCode  string
}

func (r *Response) Print() {

}
