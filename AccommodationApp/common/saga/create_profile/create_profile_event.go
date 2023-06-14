package create_profile

import "accommodation_booking/common/domain"

type CreateProfileCommandType int8

const (
	CreateProfile = iota
	RollbackCreatedProfile
	UnknownCommand
)

type CreateProfileCommand struct {
	Profile domain.Profile
	Type    CreateProfileCommandType
}

type CreateProfileReplyType int8

const (
	ProfileNotCreated = iota
	ProfileCreated
	ProfileCreationRolledBack
	UnknownReply
)

type CreateProfileReply struct {
	Profile domain.Profile
	Type    CreateProfileReplyType
}
