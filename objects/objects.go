package objects

type Book struct {
	Id      string "json:\"Id\""
	Name    string "json:\"Name\""
	Cost    int    "json:\"cost\""
	Content string "json:\"content\""
}
