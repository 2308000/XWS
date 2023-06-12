package update_profile

import "accommodation_booking/common/domain"

type UpdateProfileCommandType int8

const (
	UpdateProfile UpdateProfileCommandType = iota
	RollbackUpdatedProfile
	UnknownCommand
)

type UpdateProfileCommand struct {
	NewProfile domain.Profile
	OldProfile domain.Profile
	Type       UpdateProfileCommandType
}

type UpdateProfileReplyType int8

const (
	ProfileUpdated UpdateProfileReplyType = iota
	ProfileNotUpdated
	ProfileUpdateRolledBack
	UnknownReply
)

type UpdateProfileReply struct {
	NewProfile domain.Profile
	OldProfile domain.Profile
	Type       UpdateProfileReplyType
}
