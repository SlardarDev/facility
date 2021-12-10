package others

type MapStringInterface map[string]interface{}

func (m MapStringInterface) Extend(m2 MapStringInterface) {
	for k, v := range m2 {
		if _, ok := m[k]; !ok {
			m[k] = v
		}
	}
}

func (m MapStringInterface) Overwrite(m2 MapStringInterface) {
	for k, v := range m2 {
		if _, ok := m[k]; ok {
			m[k] = v
		}
	}
}

func (m MapStringInterface) ExtendAndOverwrite(m2 MapStringInterface) {
	for k, v := range m2 {
		m[k] = v
	}
}

func (m MapStringInterface) Remove(key string) {
	delete(m, key)
}

func (m MapStringInterface) Contains(key string) bool {
	_, ok := m[key]
	return ok
}

func (m MapStringInterface) MinusInPlace(m2 MapStringInterface) {
	for k := range m2 {
		m.Remove(k)
	}
}

func (m MapStringInterface) Minus(m2 MapStringInterface) MapStringInterface {
	newMap := MapStringInterface{}
	for k, v := range m2 {
		if !m.Contains(k) {
			newMap[k] = v
		}
	}
	return newMap
}

func (m MapStringInterface) Put(key string, value interface{}) {
	m[key] = value
}
