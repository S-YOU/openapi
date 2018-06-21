package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

// OAuthFlowsBuilder is used to build an instance of OAuthFlows. The user must
// call `Do()` after providing all the necessary information to
// build an instance of OAuthFlows
type OAuthFlowsBuilder struct {
	target *oauthFlows
}

// Do finalizes the building process for OAuthFlows and returns the result
func (b *OAuthFlowsBuilder) Do() OAuthFlows {
	return b.target
}

// NewOAuthFlows creates a new builder object for OAuthFlows
func NewOAuthFlows() *OAuthFlowsBuilder {
	return &OAuthFlowsBuilder{
		target: &oauthFlows{},
	}
}

// Implicit sets the Implicit field for object OAuthFlows.
func (b *OAuthFlowsBuilder) Implicit(v OAuthFlow) *OAuthFlowsBuilder {
	b.target.implicit = v
	return b
}

// Password sets the Password field for object OAuthFlows.
func (b *OAuthFlowsBuilder) Password(v OAuthFlow) *OAuthFlowsBuilder {
	b.target.password = v
	return b
}

// ClientCredentials sets the ClientCredentials field for object OAuthFlows.
func (b *OAuthFlowsBuilder) ClientCredentials(v OAuthFlow) *OAuthFlowsBuilder {
	b.target.clientCredentials = v
	return b
}

// AuthorizationCode sets the AuthorizationCode field for object OAuthFlows.
func (b *OAuthFlowsBuilder) AuthorizationCode(v OAuthFlow) *OAuthFlowsBuilder {
	b.target.authorizationCode = v
	return b
}

// Reference sets the $ref (reference) field for object OAuthFlows.
func (b *OAuthFlowsBuilder) Reference(v string) *OAuthFlowsBuilder {
	b.target.reference = v
	return b
}
