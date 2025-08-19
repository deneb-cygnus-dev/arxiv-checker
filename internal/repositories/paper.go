package repositories

// Paper is the data holding the basic paper information
type Paper struct {
	Author   Author
	Title    string
	Abstract string
	Class    string
	Keywords []string
	URL      string
}

// Author is the data holding the basic author information
type Author struct {
	Name         string
	Organization string
}

// PaperClass is the data holding the arXiv class identifiers
type PaperClass struct {
	Category string
	Field    string
}
