package export

import (
	"fmt"
	"os"

	"github.com/kikiluvv/clockr/db"
	"github.com/kikiluvv/clockr/utils"
)

func ExportHours() error {
	dbData, err := db.Load()
	if err != nil {
		return err
	}

	filename := "hours.txt"
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	for _, s := range dbData.Sessions {
		breakHours := utils.BreakDurationHours(s.Breaks)
		line := fmt.Sprintf("%s: In %s, Out %s, Total %.2f hours (breaks %.2f)\n",
			s.Date, s.ClockIn, s.ClockOut, s.TotalHours-breakHours, breakHours)
		_, err := f.WriteString(line)
		if err != nil {
			return err
		}
	}

	return nil
}
