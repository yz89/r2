package context

import "errors"

type Env []map[string]string

func ExtEnv(x, v string, env Env) Env {
	newEnv := make([]map[string]string, 0)

	newKV := make(map[string]string)
	newKV[x] = v

	newEnv = append(newEnv, newKV)
	newEnv = append(newEnv, env...)

	return newEnv
}

func Lookup(x string, env Env) (string, error) {
	for _, c := range env {
		if v, ok := c[x]; ok {
			return v, nil
		}
	}
	return "", errors.New("symbol not found")
}
