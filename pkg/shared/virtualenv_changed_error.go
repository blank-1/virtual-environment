package shared

type VirtualEnvChangedError struct {
}

func (v VirtualEnvChangedError) Error() string {
	return "detected virtual environment instance change"
}
