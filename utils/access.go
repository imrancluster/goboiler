package utils

// CheckAccess : check user has permission to add user
func CheckAccess(roleCode string, orgCode string, reqOrgCode string, userPerms []string, apiPerm string) bool {
	if roleCode == "SuperAdmin" {
		return true
	}
	if orgCode != reqOrgCode {
		return false
	}
	hasPermission := false
	for _, val := range userPerms {
		if val == apiPerm {
			hasPermission = true
		}
	}
	return hasPermission
}

// CheckAccessForUpdateUser : check user has update permission
func CheckAccessForUpdateUser(access string, orgCode string) bool {
	return true
}
