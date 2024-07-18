package mocks

import "github.com.com/crudBook/pkg/models"

var Books = []models.Book{
	{Id: 1,
		Title:  "Golang",
		Author: "Gopher",
		Desc:   "Golang tour"},

	{Id: 2,
		Title:  "C++",
		Author: "System devs",
		Desc:   "C++ tutorial"},

	{Id: 3,
		Title:  "Java",
		Author: "Java devs",
		Desc:   "About Java"},
		
	{Id: 4,
		Title:  "Scala",
		Author: "Scala prog",
		Desc:   "About Scala"},
}
