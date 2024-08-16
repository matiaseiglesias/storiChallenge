package servicestest

import (
	"testing"

	"github.com/matiaseiglesias/storiChallenge/internal/services"
)

func TestDate(t *testing.T) {

	t.Run("Calculate every month name", func(t *testing.T) {

		month_1 := services.GetMonthName("15/01")
		month_2 := services.GetMonthName("15/02")
		month_3 := services.GetMonthName("15/03")
		month_4 := services.GetMonthName("15/04")
		month_5 := services.GetMonthName("15/05")
		month_6 := services.GetMonthName("15/06")
		month_7 := services.GetMonthName("15/07")
		month_8 := services.GetMonthName("15/08")
		month_9 := services.GetMonthName("15/09")
		month_10 := services.GetMonthName("15/10")
		month_11 := services.GetMonthName("15/11")
		month_12 := services.GetMonthName("15/12")

		if month_1 != "January" {
			t.Error("error while calculating January month name, result ", month_1)
		}
		if month_2 != "February" {
			t.Error("error while calculating February month name, result ", month_2)
		}
		if month_3 != "March" {
			t.Error("error while calculating March month name, result ", month_3)
		}
		if month_4 != "April" {
			t.Error("error while calculating April month name, result ", month_4)
		}
		if month_5 != "May" {
			t.Error("error while calculating May month name, result ", month_5)
		}
		if month_6 != "June" {
			t.Error("error while calculating June month name, result ", month_6)
		}
		if month_7 != "July" {
			t.Error("error while calculating July month name, result ", month_7)
		}
		if month_8 != "August" {
			t.Error("error while calculating August month name, result ", month_8)
		}
		if month_9 != "September" {
			t.Error("error while calculating September month name, result ", month_9)
		}
		if month_10 != "October" {
			t.Error("error while calculating October month name, result ", month_10)
		}
		if month_11 != "November" {
			t.Error("error while calculating November month name, result ", month_11)
		}
		if month_12 != "December" {
			t.Error("error while calculating December month name, result ", month_12)
		}
	})
}
