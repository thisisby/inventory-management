package records

type Stores struct {
	Id       string `db:"id"`
	Name     string `db:"name"`
	Address  string `db:"address"`
	Contacts string `db:"contacts"`
}
