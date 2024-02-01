package models

import (
	"fmt"

	"github.com/17Peeyush/TinyURLClone-GO/utils"
)

type Meet struct {
	UniqueID string
	Url     string `binding: "required"`
}

var tempDB = []Meet{}

func NewEntry(link string) {

}

func (m *Meet) Save() error {
	// will save the hashed string in the db
	m.HashLink()
	return nil
}

func (m *Meet) HashLink() {
	// will hash the original string here with the help of a hashing function defined in util function
	//once the hashed string is stored in uniqueID then will try to save it in db and check if hash does not already exist

	hashedLink := utils.LinkHash(m.Url)
	fmt.Println(hashedLink)
	//First need to check if this hashed string is already present in db
	m.UniqueID = hashedLink
}

func (m *Meet) GetLink() {
	/*when we get something like tinyurl.com/xxxxxxx then will search this xxxxxxx in db and get it and return it
	where they will redirect it to original meeting link.
	*/
}