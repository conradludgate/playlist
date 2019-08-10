package main

func Redirect(status int, url string) response {
	headers := make(map[string]string)
	headers["Location"] = url

	return response{status: status, headers: headers}
}
