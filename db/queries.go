package db

// GETALLUSERS is the Query to get all Users
const GETALLUSERS string = `
SELECT 	id, first_name, last_name, phone
FROM users
LIMIT %s
OFFSET %s`

// GETALLRESTAURANTS is the Query to get all the Restraunts
const GETALLRESTAURANTS string = `
SELECT
  r.id as id,
  r.name as name,
  r.category as category
FROM restaurants as r
LIMIT %s
OFFSET %s`

// GETALLVENUESBYRESTAURANT is the Query to get all Venues associated to a Restraunt
const GETALLVENUESBYRESTAURANT string = `
SELECT
  v.id as id,
  v.street_address as street_address,
  v.city as city,
  v.state as state
FROM venues as v
INNER JOIN restaurants as r
ON r.id = v.restaurant_id
WHERE r.id = %s`

// GETALLRATINGS is the Query all raitings for a given constraint. Default
const GETALLRATINGS string = `
SELECT
  rate.cost as cost,
  rate.food as food,
  rate.cleanliness_service as cleanliness_service,
  rate.total_score as total_score
  rate.date_time_created as date_time_created
  rate.date_time_updated as date_time_updated
FROM ratings as rate
%s
LIMIT %s
OFFSET %s
`
