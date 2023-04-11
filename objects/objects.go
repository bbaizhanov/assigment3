package objects

type Book struct {
	Id    string "json:\"Id\""
	Title string "json:\"Title\""
	Desc  string "json\"Desc\""
	Cost  int    "json:\"Cost\""
}
