package apierror

type Level string

func (e Level) String() string {
	return string(e)
}

const (
	LevelUser   Level = "user"
	LevelSystem Level = "system"
)
