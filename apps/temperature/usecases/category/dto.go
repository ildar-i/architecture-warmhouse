package category

type GetCategoryDto struct {
	CategoryId string `pg:"category_id"`
	Slug       string `pg:"slug"`
}

type AddCategoryDto struct {
	Slug         string `pg:"slug"`
	PluralName   string `pg:"plural_name"`
	SingularName string `pg:"singular_name"`
	Description  string `pg:"description"`
	Img          string `pg:"img"`
}
