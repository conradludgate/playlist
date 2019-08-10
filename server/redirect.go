package main

func redirect(status int, url string) Response {
	headers := make(map[string]string)
	headers["Location"] = url

	return Response{status: status, headers: headers}
}
