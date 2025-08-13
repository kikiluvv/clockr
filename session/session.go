package session

import (
	"fmt"
	"time"

	"github.com/kikiluvv/clockr/db"
	"github.com/kikiluvv/clockr/utils"
)

func ClockIn() {
	dbData, _ := db.Load()
	now := utils.NowTimeString()
	date := utils.TodayDateString()

	// check if already clocked in today
	for _, s := range dbData.Sessions {
		if s.Date == date && s.ClockIn != "" {
			fmt.Println("Already clocked in today at", s.ClockIn)
			return
		}
	}

	newSession := db.Session{
		Date:    date,
		ClockIn: now,
	}
	dbData.Sessions = append(dbData.Sessions, newSession)
	db.Save(dbData)
	fmt.Println("Clocked in at", now)
}

func ClockOut() {
	dbData, _ := db.Load()
	date := utils.TodayDateString()
	now := utils.NowTimeString()

	for i, s := range dbData.Sessions {
		if s.Date == date && s.ClockOut == "" {
			dbData.Sessions[i].ClockOut = now
			hours, _ := utils.DurationInHours(s.ClockIn, now)
			dbData.Sessions[i].TotalHours = hours
			db.Save(dbData)
			fmt.Println("Clocked out at", now, "- Total hours:", hours)
			return
		}
	}
	fmt.Println("No active session to clock out today.")
}

func Status() {
	dbData, _ := db.Load()
	date := utils.TodayDateString()
	for _, s := range dbData.Sessions {
		if s.Date == date {
			breakHours := utils.BreakDurationHours(s.Breaks)
			fmt.Printf("Today: In: %s Out: %s Total: %.2f (breaks: %.2f)\n", s.ClockIn, s.ClockOut, s.TotalHours-breakHours, breakHours)
			return
		}
	}
	fmt.Println("No session today.")
}

func BreakStart() {
	dbData, _ := db.Load()
	date := utils.TodayDateString()
	now := utils.NowTimeString()

	for i, s := range dbData.Sessions {
		if s.Date == date && s.ClockOut == "" {
			dbData.Sessions[i].Breaks = append(dbData.Sessions[i].Breaks, db.Break{Start: now})
			db.Save(dbData)
			fmt.Println("Break started at", now)
			return
		}
	}
	fmt.Println("No active session to start a break.")
}

func BreakEnd() {
	dbData, _ := db.Load()
	date := utils.TodayDateString()
	now := utils.NowTimeString()

	for i, s := range dbData.Sessions {
		if s.Date == date && s.ClockOut == "" {
			if len(s.Breaks) == 0 || s.Breaks[len(s.Breaks)-1].End != "" {
				fmt.Println("No break in progress.")
				return
			}
			dbData.Sessions[i].Breaks[len(s.Breaks)-1].End = now
			db.Save(dbData)
			fmt.Println("Break ended at", now)
			return
		}
	}
	fmt.Println("No active session to end a break.")
}

// ShowSummary prints hours worked today, this week, and current pay period
func ShowSummary() {
	dbData, err := db.Load()
	if err != nil {
		fmt.Println("Error loading database:", err)
		return
	}

	now := time.Now()
	today := now.Format(utils.DateLayout)

	var todayHours, weekHours, periodHours float64

	// calculate start of week (Monday)
	weekday := int(now.Weekday())
	if weekday == 0 {
		weekday = 7 // sunday = 7
	}
	monday := now.AddDate(0, 0, -weekday+1)
	weekStart := monday.Format(utils.DateLayout)

	periodStart := dbData.Config.StartOfPeriod
	periodEnd := dbData.Config.EndOfPeriod

	for _, s := range dbData.Sessions {
		if s.Date == today {
			todayHours += s.TotalHours - utils.BreakDurationHours(s.Breaks)
		}
		if s.Date >= weekStart && s.Date <= today {
			weekHours += s.TotalHours - utils.BreakDurationHours(s.Breaks)
		}
		if s.Date >= periodStart && s.Date <= periodEnd {
			periodHours += s.TotalHours - utils.BreakDurationHours(s.Breaks)
		}
	}

	fmt.Printf("Summary:\n")
	fmt.Printf("Today: %.2f hours\n", todayHours)
	fmt.Printf("This Week: %.2f hours\n", weekHours)
	fmt.Printf("Current Pay Period (%s to %s): %.2f hours\n", periodStart, periodEnd, periodHours)
}

// AdjustSession lets you modify a session's in/out time and recalculates total hours
func AdjustSession(date, field, newTime string) {
	dbData, err := db.Load()
	if err != nil {
		fmt.Println("Error loading database:", err)
		return
	}

	found := false
	for i, s := range dbData.Sessions {
		if s.Date == date {
			found = true
			switch field {
			case "in":
				dbData.Sessions[i].ClockIn = newTime
			case "out":
				dbData.Sessions[i].ClockOut = newTime
			default:
				fmt.Println("Invalid field:", field)
				fmt.Println("Use 'in' or 'out'")
				return
			}

			// recalc total hours if both in/out exist
			if dbData.Sessions[i].ClockIn != "" && dbData.Sessions[i].ClockOut != "" {
				hours, err := utils.DurationInHours(dbData.Sessions[i].ClockIn, dbData.Sessions[i].ClockOut)
				if err != nil {
					fmt.Println("Error calculating total hours:", err)
					return
				}
				dbData.Sessions[i].TotalHours = hours
			}

			db.Save(dbData)
			fmt.Printf("Session on %s updated: %s = %s\n", date, field, newTime)
			return
		}
	}

	if !found {
		fmt.Println("No session found on date:", date)
	}
}

// StatusString returns today's status as a string instead of printing
func StatusString() string {
	dbData, _ := db.Load()
	date := utils.TodayDateString()
	for _, s := range dbData.Sessions {
		if s.Date == date {
			breakHours := utils.BreakDurationHours(s.Breaks)
			return fmt.Sprintf("Today: In: %s Out: %s Total: %.2f (breaks: %.2f)", s.ClockIn, s.ClockOut, s.TotalHours-breakHours, breakHours)
		}
	}
	return "No session today."
}

// SummaryString returns the summary as a string instead of printing
func SummaryString() string {
	dbData, err := db.Load()
	if err != nil {
		return "Error loading database: " + err.Error()
	}

	now := time.Now()
	today := now.Format(utils.DateLayout)

	var todayHours, weekHours, periodHours float64

	weekday := int(now.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	monday := now.AddDate(0, 0, -weekday+1)
	weekStart := monday.Format(utils.DateLayout)

	periodStart := dbData.Config.StartOfPeriod
	periodEnd := dbData.Config.EndOfPeriod

	for _, s := range dbData.Sessions {
		breakHours := utils.BreakDurationHours(s.Breaks)
		if s.Date == today {
			todayHours += s.TotalHours - breakHours
		}
		if s.Date >= weekStart && s.Date <= today {
			weekHours += s.TotalHours - breakHours
		}
		if s.Date >= periodStart && s.Date <= periodEnd {
			periodHours += s.TotalHours - breakHours
		}
	}

	return fmt.Sprintf("Today: %.2f hours\nThis Week: %.2f hours\nCurrent Pay Period (%s to %s): %.2f hours", todayHours, weekHours, periodStart, periodEnd, periodHours)
}
