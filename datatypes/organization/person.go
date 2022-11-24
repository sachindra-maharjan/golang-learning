package organization

import (
	"errors"
	"fmt"
	"strings"
)

type organization interface {
	ID() string
}

type Person struct {
	firstName string
	lastName string
	twitterHandle string
}

func NewPerson(firstname, lastname string) Person{
	return Person{firstName: firstname, lastName: lastname}
}

func (p *Person) ID() string {
	return "12345"
}

func (p *Person) FullName() string {
	return fmt.Sprintf("%v %v", p.firstName, p.lastName)
}

func (p *Person) TwitterHandle() string {
	return p.twitterHandle
}

func (p *Person) SetTwitterHandler(twitterHandle string) error {
	if len(twitterHandle) == 0 {
		p.twitterHandle = twitterHandle
	} else if !strings.HasPrefix(twitterHandle, "@") {
		return errors.New("Twitter handlers must start with @")
	}

	p.twitterHandle = twitterHandle
	return nil
}


