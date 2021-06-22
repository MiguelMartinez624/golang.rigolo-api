package authentication

type AuthError struct {
	Msg string `json:"msg"`
}

func (e AuthError) Error() string {
	return e.Msg
}

var (
	ErrInvalidTokenID = &AuthError{"ID token is invalid."}
	ErrCouldNotParseClains = &AuthError{"ID token valid but couldn't parse claims."}
)
