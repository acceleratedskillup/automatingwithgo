package omnientity

import (
	"fmt"
	"sync"
)

type omnientity struct {
	name   string //take note of small letters
	height float64
}

var lock = &sync.Mutex{}

var singleInstance *omnientity

func GetInstance() *omnientity {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creating single instance now.")
			singleInstance = &omnientity{"OmniEntity", 10.00}
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstance
}

//no setters, only getters
func (omnientity *omnientity) GetName() string {
	return omnientity.name
}

func (omnientity *omnientity) GetHeight() float64 {
	return omnientity.height
}

func (omnientity *omnientity) GetString() string {
	return fmt.Sprintf("omnientity name:%s height:%v", omnientity.name, omnientity.height)
}
