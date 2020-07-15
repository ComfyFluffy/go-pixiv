package pixiv

func IsInvalidCredentials(err error) bool {
	if er, ok := err.(*ErrAuth); ok {
		if er.Errors.System.Code == 1508 {
			return true
		}
	}
	return false
}
