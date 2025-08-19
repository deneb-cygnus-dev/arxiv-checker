package repositories

import "time"

type SearchOption struct {
	MaxPapers int
	StartDate time.Time
	EndDate   time.Time
}

// PaperRepository stores the cached papers
type PaperRepository struct {
	Papers []Paper
}

// NewPaperRepository returns an initialized PaperRepository
func NewPaperRepository() *PaperRepository {
	return &PaperRepository{}
}

// Search searches the arXiv.org and gets the papers
func (r PaperRepository) Search(searchWords []string, searchClasses []PaperClass, searchOption SearchOption) (int, bool, error) {
	return 0, true, nil
}

func (r PaperRepository) Save() error {
	return nil
}
