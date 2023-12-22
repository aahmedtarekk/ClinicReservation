package APIs

type Response struct {
	ResponseStatus  bool
	ResponseMessage string
	UserUUID        string
	ResponseData    interface{}
}
