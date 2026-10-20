package main
import ("fmt")


type command struct {
	name string
	args []string
}

type Commands struct {
	 com map[string]func(*state,command) error
}




func (c *Commands) run(s *state, cmd command) error {
	value, ok := c.com[cmd.name]
	if ok != true {
		return fmt.Errorf("Not found")
	}
	err := value(s,cmd)
	
	return err
}

func (c *Commands) register(name string, f func(*state, command) error) {
	
	c.com[name] = f
	
}
