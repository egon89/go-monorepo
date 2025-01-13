package jwt

import (
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gookit/goutil"
)

type JwtHelper struct {
	claims       jwt.MapClaims
	realmRoles   []string
	accountRoles []string
	scopes       []string
}

func NewJwtHelper(claims jwt.MapClaims) *JwtHelper {
	return &JwtHelper{
		claims:       claims,
		realmRoles:   parseRealRoles(claims),
		accountRoles: parseAccountRoles(claims),
		scopes:       parseScopes(claims),
	}
}

func (j *JwtHelper) GetUserId() (string, error) {
	return j.claims.GetSubject()
}

func (j *JwtHelper) HasRole(role string) bool {
	return goutil.Contains(j.realmRoles, role)
}

func (j *JwtHelper) TokenHasScope(scope string) bool {
	return goutil.Contains(j.scopes, scope)
}

func parseRealRoles(claims jwt.MapClaims) []string {
	var realmRoles []string = make([]string, 0)

	if claim, ok := claims["realm_access"]; ok {
		if roles, ok := claim.(map[string]interface{})["roles"]; ok {
			for _, role := range roles.([]interface{}) {
				realmRoles = append(realmRoles, role.(string))
			}
		}
	}

	return realmRoles
}

func parseAccountRoles(claims jwt.MapClaims) []string {
	var accountRoles []string = make([]string, 0)

	if claim, ok := claims["resource_access"]; ok {
		if acc, ok := claim.(map[string]interface{})["account"]; ok {
			if roles, ok := acc.(map[string]interface{})["roles"]; ok {
				for _, role := range roles.([]interface{}) {
					accountRoles = append(accountRoles, role.(string))
				}
			}
		}
	}

	return accountRoles
}

func parseScopes(claims jwt.MapClaims) []string {
	scopeStr, err := parseString(claims, "scope")
	if err != nil {
		return make([]string, 0)
	}

	scopes := strings.Split(scopeStr, " ")

	return scopes
}

func parseString(claims jwt.MapClaims, key string) (string, error) {
	var (
		ok    bool
		raw   interface{}
		value string
	)

	raw, ok = claims[key]
	if !ok {
		return "", nil
	}

	value, ok = raw.(string)
	if !ok {
		return "", fmt.Errorf("key %s is invalid", key)
	}

	return value, nil
}
