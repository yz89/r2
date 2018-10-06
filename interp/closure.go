package interp

import (
	"encoding/json"
	"r2/context"
	"strings"
)

// Closure is used to store lambda expression and its context
type Closure struct {
	param string
	exp   string
	env   context.Env
}

func (c *Closure) Serialize() (string, error) {
	buf := c.param
	buf += ";"
	buf += c.exp
	for _, dict := range c.env {
		buf += ";"
		dictStr, err := json.Marshal(dict)
		if err != nil {
			return "", err
		}
		buf += string(dictStr)
	}

	return buf, nil
}

func (c *Closure) Deserialize(data string) error {
	dataArray := strings.Split(data, ";")
	c.param = dataArray[0]
	c.exp = dataArray[1]
	for _, dict := range dataArray[2:] {
		newMap := make(map[string]string)
		if err := json.Unmarshal([]byte(dict), &newMap); err != nil {
			return err
		}
		c.env = append(c.env, newMap)
	}
	return nil
}
