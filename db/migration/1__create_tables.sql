CREATE DATABASE rest_restaurants;

DROP TABLE IF EXISTS users CASCADE;
CREATE TABLE users (
	id SERIAL PRIMARY KEY,
	first_name VARCHAR(50) NOT NULL,
  last_name VARCHAR(50) NOT NULL,
	phone VARCHAR(10) NOT NULL,
  UNIQUE (first_name, last_name, phone)
) ;

DROP TABLE IF EXISTS restaurants CASCADE;
CREATE TABLE restaurants (
	id SERIAL PRIMARY KEY,
	name VARCHAR(50) NOT NULL,
  category VARCHAR(50) NOT NULL,
  UNIQUE (name, category)
);

DROP TABLE IF EXISTS venues CASCADE;
CREATE TABLE venues (
	id SERIAL PRIMARY KEY,
	street_address VARCHAR(95) NOT NULL,
	city VARCHAR(40) NOT NULL,
	state VARCHAR(4) NOT NULL,
	restaurant_id INT NOT NULL,
	constraint fk__venues__restaurants
	 foreign key (restaurant_id)
	 REFERENCES restaurants (id),
	UNIQUE (street_address, city, state, restaurant_id)
) ;

DROP TABLE IF EXISTS ratings CASCADE;
CREATE TABLE ratings (
	id SERIAL PRIMARY KEY,
	cost SMALLINT NOT NULL,
  food SMALLINT NOT NULL,
	cleanliness_service SMALLINT NOT NULL,
  total_score SMALLINT NOT NULL,
  venue_id int NOT NULL,
  user_id int NOT NULL,
	comments VARCHAR(400),
  date_time_created TIMESTAMP,
  date_time_updated TIMESTAMP,
  constraint fk__rating__venues
   foreign key (venue_id)
   REFERENCES venues (id),
  constraint fk__rating__users
   foreign key (user_id)
   REFERENCES users (id),
  UNIQUE ( user_id, venue_id )
);
