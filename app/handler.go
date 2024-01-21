package app

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"routine.nfitton.com/app/templates"
	"routine.nfitton.com/routines"
	"routine.nfitton.com/users"
)

var dayInSeconds = int64(60 * 60 * 24)

func dayNumToShortStr(i int) string {
	switch i {
	case 0:
		return "Sun"
	case 1:
		return "Mon"
	case 2:
		return "Tues"
	case 3:
		return "Wed"
	case 4:
		return "Thur"
	case 5:
		return "Fri"
	case 6:
		return "Sat"
	}
	return "-"
}

/*
bg-red-200
bg-green-200
bg-gray-200
bg-yellow-200
bg-gray-300
*/
func determineStatus(dayDiff int, isActive bool) string {
	if dayDiff < 0 {
		if isActive {
			return "bg-red-200"
			// return "bg-green-200"
		}
		return "bg-gray-200"
	} else if dayDiff == 0 {
		if isActive {
			return "bg-yellow-200"
		} else {
			return "bg-gray-300"
		}
	}
	if isActive {
		return "bg-yellow-200"
	}
	return "bg-gray-300"
}

func HandleRoot(c *gin.Context) {
	metadata := templates.Metadata{
		Title:       "GoRoutines",
		Description: "A routine management project written in Golang",
	}
	users := users.GetAll()
	content := templates.Root(users)
	component := templates.Wrapper(metadata, content)
	HtmlResponse(component)(c)
}

func HandleRoutines(c *gin.Context) {
	metadata := templates.Metadata{
		Title:       "Your routines",
		Description: "A routine management project written in Golang",
	}
	// userId, err := response.GetIdFromPath(c, "userId")
	// if err != nil {
	// 	c.HTML(http.StatusNotFound, "", templates.Wrapper(metadata, templates.NotFound("User ID invalid", err)))
	// 	return
	// }

	// user, found := users.FindById(userId)
	// if !found {
	// 	c.HTML(http.StatusNotFound, "", templates.Wrapper(metadata, templates.NotFound("User not found", err)))
	// 	return
	// }

	// Generate a random routine

	r := routines.Routine{
		Id:     uuid.New(),
		Name:   "Do Duolingo",
		UserId: uuid.New(),
		Public: false,
		InDays: 6,
		Every:  4,
		Offset: 0,
	}
	routines := [][]templates.TaskDay{}

	for off := 0; off < 6; off++ {
		daysRange := []templates.TaskDay{}
		for i := -7; i <= 7; i++ {
			date := time.Now().AddDate(0, 0, i)
			unixDays := int(date.Unix()/dayInSeconds) - off
			mod := unixDays % r.InDays
			isActive := mod < r.Every
			status := determineStatus(i, isActive)
			daysRange = append(daysRange, templates.TaskDay{
				Date:        date,
				Weekday:     dayNumToShortStr(int(date.Weekday())),
				Day:         date.Day(),
				StatusColor: status,
				Today:       i == 0,
			})
		}
		routines = append(routines, daysRange)
	}

	// userRoutines := routines.GetAllByUser(userId, false)
	content := templates.Routines(routines)
	component := templates.Wrapper(metadata, content)
	HtmlResponse(component)(c)
}
