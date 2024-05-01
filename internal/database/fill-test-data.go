package database

import (
	"cars/internal/models"
)

func FillTestData(db CardyB) {
	peoples := []models.People{
		{Id: 1, Name: "Ivan", Surname: "Ivanov", Patronymic: "Ivanovich"},
		{Id: 2, Name: "Petr", Surname: "Petrov", Patronymic: "Petrovich"},
		{Id: 3, Name: "Sidor", Surname: "Sidorov", Patronymic: "Sidorovich"},
		{Id: 4, Name: "Vasiliy", Surname: "Vasiliev", Patronymic: "Vasilievich"},
		{Id: 5, Name: "Vladimir", Surname: "Vladimirov", Patronymic: "Vladimirovich"},
		{Id: 6, Name: "Alexey", Surname: "Alexeev", Patronymic: "Alexeevich"},
		{Id: 7, Name: "Alexey", Surname: "Alexeev", Patronymic: "Alexeevich"},
		{Id: 8, Name: "Ivan", Surname: "Ivanov", Patronymic: "Ivanovich"},
		{Id: 9, Name: "Petr", Surname: "Petrov", Patronymic: "Petrovich"},
		{Id: 10, Name: "Sidor", Surname: "Sidorov", Patronymic: "Sidorovich"},
		{Id: 11, Name: "Vasiliy", Surname: "Vasiliev", Patronymic: "Vasilievich"},
		{Id: 12, Name: "Vladimir", Surname: "Vladimirov", Patronymic: "Vladimirovich"},
		{Id: 13, Name: "Alexey", Surname: "Alexeev", Patronymic: "Alexeevich"},
		{Id: 14, Name: "Alexey", Surname: "Alexeev", Patronymic: "Alexeevich"},
		{Id: 15, Name: "Alex Under", Surname: "Pussi can", Patronymic: "Sir Gay Evil Each"},
	}

	for _, people := range peoples {
		db.AddNewPeople(people)
	}
	cars := []models.Car{
		{RegNum: "X123XX150", Mark: "Lada", Model: "Vesta", Year: 2002, Owner: models.People{Id: 1}},
		{RegNum: "X123XX151", Mark: "Lada", Model: "Granta", Year: 2012, Owner: models.People{Id: 2}},
		{RegNum: "X123XX152", Mark: "Lada", Model: "Priora", Year: 2012, Owner: models.People{Id: 3}},
		{RegNum: "X123XX153", Mark: "Lada", Model: "Xray", Year: 2014, Owner: models.People{Id: 4}},
		{RegNum: "X123XX154", Mark: "Lada", Model: "Vesta", Year: 2012, Owner: models.People{Id: 5}},
		{RegNum: "X123XX155", Mark: "Lada", Model: "Granta", Year: 2014, Owner: models.People{Id: 6}},
		{RegNum: "X123XX156", Mark: "Lada", Model: "Priora", Year: 2014, Owner: models.People{Id: 7}},
		{RegNum: "X123XX157", Mark: "Lada", Model: "Xray", Year: 2015, Owner: models.People{Id: 8}},
		{RegNum: "X123XX158", Mark: "Lada", Model: "Vesta", Year: 2014, Owner: models.People{Id: 9}},
		{RegNum: "X123XX159", Mark: "Lada", Model: "Granta", Year: 2016, Owner: models.People{Id: 10}},
		{RegNum: "X123XX160", Mark: "Lada", Model: "Priora", Year: 2016, Owner: models.People{Id: 11}},
		{RegNum: "X123XX161", Mark: "Lada", Model: "Xray", Year: 2017, Owner: models.People{Id: 12}},
		{RegNum: "X123XX162", Mark: "Lada", Model: "Vesta", Year: 2016, Owner: models.People{Id: 13}},
		{RegNum: "X123XX163", Mark: "Lada", Model: "Granta", Year: 2018, Owner: models.People{Id: 14}},
		{RegNum: "X123XX164", Mark: "VAZ", Model: "Priora", Year: 2018, Owner: models.People{Id: 5}},
		{RegNum: "X123XX165", Mark: "VAZ", Model: "Xray", Year: 2019, Owner: models.People{Id: 6}},
		{RegNum: "X123XX166", Mark: "VAZ", Model: "Vesta", Year: 2018, Owner: models.People{Id: 7}},
		{RegNum: "X123XX167", Mark: "VAZ", Model: "Granta", Year: 2020, Owner: models.People{Id: 8}},
		{RegNum: "X123XX168", Mark: "VAZ", Model: "Priora", Year: 2020, Owner: models.People{Id: 9}},
		{RegNum: "X123XX169", Mark: "VAZ", Model: "Xray", Year: 2021, Owner: models.People{Id: 10}},
		{RegNum: "X123XX170", Mark: "VAZ", Model: "Vesta", Year: 2020, Owner: models.People{Id: 1}},
		{RegNum: "X123XX171", Mark: "VAZ", Model: "Granta", Year: 2022, Owner: models.People{Id: 2}},
		{RegNum: "X123XX172", Mark: "VAZ", Model: "Priora", Year: 2022, Owner: models.People{Id: 3}},
	}

	for _, car := range cars {
		db.AddNewCar(car)
	}
}
