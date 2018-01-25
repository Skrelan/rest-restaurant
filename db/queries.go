package db

// GETALLUSERS is the Query to get all users
const GETALLUSERS string = `
SELECT 	id, first_name, last_name, phone
FROM users;`
