package models

type User struct {
	ID            int
	Name          string
	PreferedBike  int
	Email         string
	Password      string
	Users []int
}
