package main

import (
	"fmt"
)

type Client struct {
	Name     string
	LastName string
}

func (c Client) FullName() string{
	return fmt.Sprintf("%s %s", c.Name, c.LastName)
}

func main() {
	client := Client{"João", "Paixão"}
	fmt.Println(client.FullName())
}
