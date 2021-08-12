package Models


// Content interface
type Content interface {
	SetValue(string,string) bool
	GetValue(string) (string ,bool)
}

