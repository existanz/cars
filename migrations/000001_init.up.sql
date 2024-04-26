CREATE TABLE IF NOT EXISTS peoples (
  id serial PRIMARY KEY,
  name varchar(50) not null, 
  surname varchar(50)not null,
  patronymic varchar(50)
);



CREATE TABLE IF NOT EXISTS cars (
	id serial PRIMARY KEY,
	reg_num varchar(10) not null,
	mark varchar(50) not null,
	model varchar(50) not null,
	year int,
	owner_id int references peoples(id) not null
);

CREATE INDEX IF NOT EXISTS cars_reg_num_idx ON cars(reg_num);