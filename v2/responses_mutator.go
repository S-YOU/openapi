package openapi

func (m *ResponsesMutator) Response(key ResponseMapKey, value Response) *ResponsesMutator {
	m.proxy.setResponse(key, value)
	return m
}

