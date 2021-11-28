package mate

var (
	_ Repo = (*Repository)(nil)
)

type Repo interface {
	FetchDB(db interface{})error
}

type Repository struct {

}

func (r *Repository) FetchDB(db interface{}) error {
	panic("implement me")
}

