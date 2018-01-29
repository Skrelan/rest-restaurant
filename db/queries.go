package db

// GETALLUSERS is the Query to get all Users
const GETALLUSERS string = `
SELECT 	id, first_name, last_name, phone
FROM users
LIMIT %s
OFFSET %s`

// GETUSERBYIDS is the Query to get User info, by id
const GETUSERBYIDS string = `
SELECT 	id, first_name, last_name, phone
FROM users
WHERE users.id IN (%s)
`

// GETALLVENUES is the Query to get all Venues
const GETALLVENUES string = `
SELECT
  v.id as id,
  r.name as name,
  r.category as category,
  v.street_address as "a.street_address",
  v.city as "a.city",
  v.state as "a.state"
FROM venues as v
INNER JOIN restaurants as r
ON r.id = v.restaurant_id
LIMIT %s
OFFSET %s`

// GETVENUESBYIDS is the Query to get Venue(s) by id(s)
const GETVENUESBYIDS string = `
SELECT
  v.id as id,
  r.name as name,
  r.category as category,
  v.street_address as "a.street_address",
  v.city as "a.city",
  v.state as "a.state"
FROM venues as v
INNER JOIN restaurants as r
ON r.id = v.restaurant_id
WHERE v.id IN (%s)`

// GETVENUESWHERE is the Query to get Venue(s) WHERE
const GETVENUESWHERE string = `
SELECT
  v.id as id,
  r.name as name,
  r.category as category,
  v.street_address as "a.street_address",
  v.city as "a.city",
  v.state as "a.state"
FROM venues as v
INNER JOIN restaurants as r
ON r.id = v.restaurant_id
WHERE %s
LIMIT %s
OFFSET %s`

// GETALLRATINGS is the Query all raitings for a given constraint.
const GETALLRATINGS string = `
SELECT
  v.id as "r.id",
  r.name as "r.name",
  u.id as "u.id",
  u.first_name as "u.first_name",
  u.last_name as "u.last_name",
  v.street_address as "r.a.street_address",
  v.city as "r.a.city",
  v.state as "r.a.state",
  rate.cost as "rate.cost",
  rate.food as "rate.food",
  rate.cleanliness_service as "rate.cleanliness_service",
  rate.total_score as "rate.total_score",
  rate.comments as "rate.comments",
  rate.date_time_created as "rate.date_time_created",
  rate.date_time_updated as "rate.date_time_updated"
FROM ratings as rate
INNER JOIN venues as v
ON rate.venue_id = v.id
INNER JOIN restaurants as r
ON r.id = v.restaurant_id
INNER JOIN users as u
ON u.id = rate.user_id
ORDER by "rate.date_time_updated" DESC
LIMIT %s
OFFSET %s
`

// GETRATINGSWHERE is the Query that returns Ratings where certain constraints are satisfied
const GETRATINGSWHERE string = `
SELECT
  v.id as "r.id",
  r.name as "r.name",
  u.id as "u.id",
  u.first_name as "u.first_name",
  u.last_name as "u.last_name",
  v.street_address as "r.a.street_address",
  v.city as "r.a.city",
  v.state as "r.a.state",
  rate.cost as "rate.cost",
  rate.food as "rate.food",
  rate.cleanliness_service as "rate.cleanliness_service",
  rate.total_score as "rate.total_score",
  rate.comments as "rate.comments",
  rate.date_time_created as "rate.date_time_created",
  rate.date_time_updated as "rate.date_time_updated"
FROM ratings as rate
INNER JOIN venues as v
ON rate.venue_id = v.id
INNER JOIN restaurants as r
ON r.id = v.restaurant_id
INNER JOIN users as u
ON u.id = rate.user_id
WHERE %s
ORDER by "rate.date_time_updated" DESC
LIMIT %s
OFFSET %s
`

//INSERTINTOUSERS is the query used to create a new user
const INSERTINTOUSERS = `
INSERT INTO users (first_name, last_name, phone)
VALUES ('%s', '%s', '%s')
`

// INSERTINTORESTAURANTS inserts new restaurant info into the restaurant table
const INSERTINTORESTAURANTS = `
INSERT INTO restaurants (name, category)
VALUES ('%s', '%s')
ON CONFLICT DO NOTHING`

// INSERTINTOVENUES inserts new venue info into the venues table
const INSERTINTOVENUES = `
INSERT INTO venues (street_address, city, state, zip_code, restaurant_id)
VALUES ('%s','%s','%s','%s',
  (SELECT id FROM restaurants as r
    WHERE r.name = '%s' AND r.category = '%s'))`

// CHECKUSERRATINGS checks if a user has given a restraunt (not just a venue), a rating in the past 30 days
const CHECKUSERRATINGS = `
SELECT count(rate.id) as count
FROM ratings as rate
INNER JOIN users as u
ON rate.user_id = u.id
INNER JOIN venues as v
ON rate.venue_id = v.id
INNER JOIN restaurants as r
ON v.restaurant_id = r.id
WHERE u.id = %d AND r.id = %d
AND rate.date_time_created >= '%s'`

// INSERTINTORATINGS inserts a new rating into the ratings table.
const INSERTINTORATINGS = `
INSERT INTO ratings
(cost,
 food,
 cleanliness_service,
 total_score,
 venue_id,
 user_id,
 comments,
 date_time_created,
 date_time_updated)
VALUES
( %d, %d, %d, %f, %d, %d, '%s', '%s', '%s')`
