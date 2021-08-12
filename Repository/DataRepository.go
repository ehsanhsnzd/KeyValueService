package Repository


type Data struct {
	Content map[string]string
}


func NewData() *Data  {
	return &Data{
		Content: make(map[string]string),
	}
}

func (data *Data) SetValue(key string,value string) bool  {
	 data.Content[key]=value
	 return true
}


func (data *Data) GetValue(key string)(string,bool)  {
	if value,found := data.Content[key]; found {
		return value,true
	}
	return "",false
}
