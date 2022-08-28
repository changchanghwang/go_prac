package mydict

import "errors"

// Dictionary type
type Dictionary map[string]string

var errNotFound = errors.New("Not Found")
var errDuplicated = errors.New("The word is exists")

func (d Dictionary) Search(word string)(string,error){
	value, exists := d[word]
	if exists{
		return value,nil
	}
	return "", errNotFound
}

// add a word to dictionary
func (d Dictionary) Add(word, def string) error{
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errDuplicated
	}
	return nil
}

func (d Dictionary) Update(word, def string) error{
	_, err := d.Search(word)
	switch err{
	case errNotFound:
		return err
	case nil:
		d[word] = def
	}
	return nil
}

func (d Dictionary) Delete(word string){
	delete(d, word)
}