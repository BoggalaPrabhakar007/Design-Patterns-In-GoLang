//Single responsibility principle
/*The single responsibility principle suggests that two separate aspects of a problem need to be handled by a different module. 
In other word, changes in a module should be originated from only one reason.*/

package main

import (
	"fmt"
	"io/ioutil"
)

var entryCount = 0

type Journal struct {
	Entries []string
}

func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d:%s", entryCount, text)
	j.Entries = append(j.Entries, entry)
	return entryCount
}

func (j *Journal) RemoveEntry(index int)  {
	//Removes the entry                                                                                                                                                                       	
}

// Separation of concerns

func (j *Journal) Save(fileName string){
	_=ioutil.WriteFile(fileName,[]byte(fileName),0644)
}

func (j *Journal) LoadFile(fileName string){
	//Load the file
}

func main(){
	
}
