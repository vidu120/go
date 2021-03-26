package calendar

import "errors"

type Date struct {
	year int
	month int
	date int
}

func (d *Date) Date() int {
	return d.date
}

func (d *Date) Month() int {
	return d.month
}

func (d *Date) Year() int {
	return d.year
}

func (d *Date) SetDate(date int) error{
	if date < 1|| date > 31{
		return errors.New("Invalid Date")
	}
	d.date = date
	return nil
}

func (d *Date) SetMonth(month int) error{
	if month < 1 || month > 12{
		return errors.New("Invalid year")
	}
	d.month = month
	return nil
}

//this is our setter for the year variable
func (d *Date) SetYear(year int) error{
	if year < 1{
		return errors.New("Invalid year")
	}
	d.year = year
	return nil
}



