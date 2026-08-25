package config

type Policy struct {
	MaxSeverity     int
	RequireAssignee bool
	AllowedStatuses []string
}

func DefaultPolicy() Policy {
	return Policy{MaxSeverity: 5, RequireAssignee: false, AllowedStatuses: []string{"new", "assigned", "in_progress", "resolved", "closed", "archived"}}
}
func (p Policy) AllowsStatus(s string) bool {
	for _, v := range p.AllowedStatuses {
		if v == s {
			return true
		}
	}
	return false
}
func (p Policy) NormalizeSeverity(s int) int {
	if s < 1 {
		return 1
	}
	if s > p.MaxSeverity {
		return p.MaxSeverity
	}
	return s
}
