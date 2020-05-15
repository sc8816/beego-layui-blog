package syserror

type Error interface {
	Code() int
	Error() string
	ReasonErr() error
}

func New(msg string, reason error) Error {
	return UnKnowError{
		msg:    msg,
		reason: reason,
	}
}
