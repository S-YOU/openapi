package openapi

// This file was automatically generated by genbuilders.go on 2018-05-20T20:06:19+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

func (v *oAuthFlows) Implicit() OAuthFlow {
	return v.implicit
}

func (v *oAuthFlows) Password() OAuthFlow {
	return v.password
}

func (v *oAuthFlows) ClientCredentials() OAuthFlow {
	return v.clientCredentials
}

func (v *oAuthFlows) AuthorizationCode() OAuthFlow {
	return v.authorizationCode
}