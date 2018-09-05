package auth

type Scope struct {
	Type    string
	Name    string
	Actions []string
}
