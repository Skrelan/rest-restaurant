
INSERT INTO users (first_name, last_name, phone) Values ('Great','Gatsby','4087001111');
INSERT INTO users (first_name, last_name, phone) Values ('Jack','Sparrow','4201000000');
INSERT INTO users (first_name, last_name, phone) Values ('Obiwan','Kenobi','4375932021');

INSERT INTO restaurants (name, category) Values ('DOSA','Indian');
INSERT INTO restaurants (name, category) Values ('Green Barn','Thai');
INSERT INTO restaurants (name, category) Values ('Epic café','Continental');

INSERT INTO venues (street_address, city, state, restaurant_id) Values ('123 Bush St.','San Francisco','CA', 1);
INSERT INTO venues (street_address, city, state, restaurant_id) Values ('250 Brandon St.','San Jose','CA', 2);
INSERT INTO venues (street_address, city, state, restaurant_id) Values ('1 Hackerway','Menlo Park','CA', 3);

INSERT INTO ratings (cost, food, cleanliness_service, total_score, venue_id, user_id, comments) VALUES (3,3,3,3,1,1,'Serve''s the most overpriced food.');
INSERT INTO ratings (cost, food, cleanliness_service, total_score, venue_id, user_id, comments) VALUES (2,4,1,3,1,2,'Pricy, but nice.');
INSERT INTO ratings (cost, food, cleanliness_service, total_score, venue_id, user_id, comments) VALUES (5,5,5,3,2,1,'');
