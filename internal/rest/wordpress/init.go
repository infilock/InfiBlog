package wordpress

// handler .
type handler struct{}

// NewHandler .
func NewHandler() Contract {
	return &handler{}
}
