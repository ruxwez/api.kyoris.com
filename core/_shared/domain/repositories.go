package sharedDomain

import (
	authDomain "api.kyoris.com/core/auth/domain"
	userDomain "api.kyoris.com/core/user/domain"
)

type Repositories struct {
	UserRepository userDomain.UserRepository
	AuthRepository authDomain.AuthRepository
}
