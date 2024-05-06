package usecase

type Schema struct {
	UserUsecaseItf
}

func New(
	userUC UserUsecaseItf,
) *Schema {
	return &Schema{
		UserUsecaseItf: userUC,
	}
}
