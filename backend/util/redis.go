package util

func GetUserEmailSigninCodeCacheKey(email string) string {
	return "sign_in_code:" + email
}
