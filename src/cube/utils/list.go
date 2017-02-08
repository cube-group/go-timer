package utils

type List struct {
	Index int
	Value []interface{}
}

func (list *List)Len() int {
	return len(list.Value)
}

func (list *List)Current() interface{} {
	if list.Len() > 0 {
		return list.Value[list.Index]
	}
	return nil
}

func (list *List)Next() interface{} {
	if list.Len() == list.Index + 1 {
		return false
	}
	list.Index++
	return list.Value[list.Index]
}

func (list *List)Pre() interface{} {
	if 0 == list.Index {
		return false
	}
	list.Index--
	return list.Value[list.Index]
}

func (list *List)Reset() {
	list.Index = 0
}

func (list *List)Remove(index int) interface{} {
	value := list.Value[index]
	list.Value = append(list.Value[:index], list.Value[index + 1:]...)
	return value
}

func (list *List)Push(value interface{}) {
	list.Value = append(list.Value, value)
}
