package domain

func ReadConfig(name string) (map[string]interface{}, error) {
	return map[string]interface{}{
		"Hello": 2,
	}, nil
}
