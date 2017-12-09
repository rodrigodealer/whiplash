package es

func config(url string) string {
	if url != "" {
		return url
	}
	return "http://127.0.0.1:9200"
}
