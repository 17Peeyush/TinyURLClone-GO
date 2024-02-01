package models

type Meet struct {
	UniqueID string
	Link     string
}

var tempDB = []Meet{}

func (m *Meet) Save() {
	// will save the hashed string in the db
}

func (m *Meet) HashLink() {
	// will hash the original string here with the help of a hashing function defined in util function
	//once the hashed string is stored in uniqueID then will try to save it in db and check if hash does not already exist
}

func (m *Meet) GetLink() {
	/*when we get something like tinyurl.com/xxxxxxx then will search this xxxxxxx in db and get it and return it
	where they will redirect it to original meeting link.
	*/
}