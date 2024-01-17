package routines

import "github.com/google/uuid"

type Routine struct {
	Id     uuid.UUID `json:"id"`
	Name   string    `binding:"required" json:"name"`
	UserId uuid.UUID `json:"userId"`
	Public bool      `json:"public"`
	Every  int       `ltfield:"InDays" json:"every" gt:0`
	InDays int       `json:"inDays" gt:0`
	Offset int       `ltfield:"InDays" json:"offset" gte:0`
}

var routines = []Routine{}

func Create(routine Routine, userId uuid.UUID) Routine {
	routine.Id = uuid.New()
	routine.UserId = userId
	routines = append(routines, routine)
	return routine
}

func GetAllByUser(userId uuid.UUID, publicOnly bool) []Routine {
	userRoutines := []Routine{}

	for _, routine := range routines {
		if routine.UserId == userId && (publicOnly == false || routine.Public == true) {
			userRoutines = append(userRoutines, routine)
		}
	}
	return userRoutines
}

func FindByUserIdAndId(userId uuid.UUID, routineId uuid.UUID) (Routine, bool) {
	for _, routine := range routines {
		if routine.UserId == userId && routine.Id == routineId {
			return routine, true
		}
	}
	return Routine{}, false

}

func DeleteByUserIdAndId(userId uuid.UUID, routineId uuid.UUID) bool {
	for i, routine := range routines {
		if routine.UserId == userId && routine.Id == routineId {
			routines = append(routines[:i], routines[i+1:]...)
			return true
		}
	}
	return false

}