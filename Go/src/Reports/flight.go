package Reports

import (
	"sort"
)

type Flight struct {
	sourceCity         string
	destinationCity    string
	sourceCountry      string
	destinationCountry string
	sourceAirport      string
	destinationAirport string
	departureDate      string
	returnDate         string
	price              int
	tripLength         int
}

type By func(f1, f2 *Flight) bool

func (by By) Sort(flights []Flight) {
	fs := &flightSorter{
		flights: flights,
		by:      by, // The Sort method's receiver is the function (closure) that defines the sort order.
	}
	sort.Sort(fs)
}

type flightSorter struct {
	flights []Flight
	by      func(p1, p2 *Flight) bool // Closure used in the Less method.
}

// Len is part of sort.Interface.
func (s *flightSorter) Len() int {
	return len(s.flights)
}

// Swap is part of sort.Interface.
func (s *flightSorter) Swap(i, j int) {
	s.flights[i], s.flights[j] = s.flights[j], s.flights[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *flightSorter) Less(i, j int) bool {
	return s.by(&s.flights[i], &s.flights[j])
}

var s_City = func(f1, f2 *Flight) bool {
	return f1.sourceCity < f2.sourceCity
}

var d_City = func(f1, f2 *Flight) bool {
	return f1.destinationCity < f2.destinationCity
}

var s_Country = func(f1, f2 *Flight) bool {
	return f1.sourceCountry < f2.sourceCountry
}

var d_Country = func(f1, f2 *Flight) bool {
	return f1.destinationCountry < f2.destinationCountry
}

var s_Airport = func(f1, f2 *Flight) bool {
	return f1.sourceAirport < f2.sourceAirport
}

var d_Airport = func(f1, f2 *Flight) bool {
	return f1.destinationAirport < f2.destinationAirport
}

var d_Date = func(f1, f2 *Flight) bool {
	return f1.departureDate < f2.departureDate
}

var r_Date = func(f1, f2 *Flight) bool {
	return f1.returnDate < f2.returnDate
}

var trip_Price = func(f1, f2 *Flight) bool {
	return f1.price < f2.price
}

var trip_Length = func(f1, f2 *Flight) bool {
	return f1.tripLength < f2.tripLength
}
