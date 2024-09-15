package extract

// type DocumentHtml struct {
// 	Title string
// 	Links []string
// }

// func FetchHtml(url string) (DocumentHtml, error) {
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return DocumentHtml{}, err
// 	}
// 	defer resp.Body.Close()
// 	doc, err := html.Parse(resp.Body)
// 	if err != nil {
// 		return DocumentHtml{}, err
// 	}
// 	return extract(doc), nil
// }
