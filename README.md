# RESTful Restaurant Review API
### Contents

1. Software Stack
2. API/Method mapping table
3. Endpoints
4. Database schema
5. Code distribution
6. Running
7. Constraints

----

### Software Stack

The software stack of this web-service can be broken down as follows:

* Backend : GOlang (v1.8)
* Database : Postgres DB
* Frontend : Javascript, HTML5, CSS3

------

### API/Method mapping table

|   | /user  | /restaurant   | /ratings      |
|---|---     |---            |---            |
|**GET**| `/users`<br> `/users?id=1,2` <br> `/users/{id}` | `/restaurants` <br> `/restaurants?id=1,2`<br> `/restaurants?name=dosa&city=San%20Jose` <br> `/restaurants/{id}`  | `/ratings` <br> `/ratings?user_id=1` <br> `/ratings?restaurant_id=1` <br> `/ratings/{id}` <br>|
|**POST**   | `/users`   | `/restaurants`  | `/ratings`   |
|**PUT**   | `/users?id=1` <br> `/users/{id}`   | `/restaurant?id=1` <br> `/restaurant/{id}`  | `/ratings?id=1` <br> `/ratings/{id}` <br> `/ratings/user_id=1&restaurant_id=2`   |  
|**DELETE**   | None  |  None  | None  |


----

### Endpoints
#### /v1
* `/users`:
  * Description : Get user(s) info
    * Method Type: `GET`
    * Input :
    * Output :

  * Description : Update a user
    * Method Type: `PUT`
    * Input :
    * Output :    

  * Description : Create a user
    * Method Type: `POST`
    * Input :
    * Output :

  * Description :
    * Method Type: `DELETE`
    * Input :
    * Output :


* `/restaurants`:

  * Description : Get restaurant(s) by name / city/ category/total score
    * Method Type: `GET`
    * Input :
    * Output :

  * Description : Update a restaurant
    * Method Type: `PUT`
    * Input :
    * Output :    

  * Description : Create a restaurant
    * Method Type: `POST`
    * Input :
    * Output :

  * Description :
    * Method Type: `DELETE`
    * Input :
    * Output :


* `/reviews`:
  * Description : Create a rating for a restaurant by a user
    * Method Type: `POST`
    * Input :
    * Output :

  * Description : Update a rating for a restaurant by a user
    * Method Type: `PUT`
    * Input :
    * Output :    

  * Description :
    * Method Type: `GET`
    * Input :
    * Output :

  * Description :
    * Method Type: `GET`
    * Input :
    * Output :

-----

### Database Schema:
The DB has 4 tables, they are as follows:
1. `users`
2. `restaurants`
3. `venues`
4. `ratings`

##### 1.users
|col_name   |type   |  
|---|---|
|id   | SERIAL   |   
|first_name   |  VARCHAR(50) |  
|last_name   |   VARCHAR(50)|   
|phone   | VARCHAR(10)  |    

constraints:
* ` UNIQUE (first_name, last_name, phone)`


##### 2. restaurant
|col_name   |type   |  
|---|---|
|id   | SERIAL   |   
|name   |  VARCHAR(50) |  
|category   |   VARCHAR(50)|   

constraints:
* ` UNIQUE (name, category)`

##### 3. venues
|col_name   |type   |  
|---|---|
|id   | SERIAL   |   
|restaurant_id   |   INT|
|street_address   |  VARCHAR(95) |  
|city   |   VARCHAR(40)|  
|state   |   VARCHAR(4)|
|zip_code   |   VARCHAR(5)|

constraints:
* ` UNIQUE (street_address, city, state, restaurant_id)`
* ` foreign key (restaurant_id)
 REFERENCES restaurants (id)`

##### 4. ratings

|col_name   |type   |  
|---|---|
|id   | SERIAL   |   
|cost   |  SMALLINT |  
|food   |  SMALLINT |   
|cleanliness_service   |  SMALLINT |  
|total_score   |   NUMERIC(5,2) |
| venue_id     |   INT          |
| user_id      |   INT          |
| comments     |   VARCHAR(400) |
| date_time_created | TIMESTAMP |
| date_time_updated | TIMESTAMP |


constraints:
* ` UNIQUE ( user_id, venue_id )`
* ` foreign key (venue_id)
  REFERENCES venues (id)`
* ` foreign key (user_id)
 REFERENCES users (id)`

 ------

 ### Code distribution
 The RESTful web-service is written GOlang and each independent chunk of code is bundled up into it's own individual package. The following packages exist in the code in this repository:

 * **db pkg** : This package contains the code for all the DB related functions, such as reading DB configs, setting up a DB connection, functions that are used to Query the DB, etc.

 * **middleware pkg** : This package contains all the code acts as the connective glue between the exposed endpoints and other packages involved

 * **models pkg** : This package contains all the models defined for the business logic of the web-application.

 * **utils pkg** : This package contains all the utility functions that help with making the code more maintainable. Functions that this package houses are ValidateNewUser, ValidateNewRestaurant, GenerateError, ResponseCodes etc.

 Other packages used are as follows:
 * [sqlx](http://jmoiron.github.io/sqlx/)
 * [logrusWrapper](https://github.com/Skrelan/LogrusWrapper)
 * [mux](http://www.gorillatoolkit.org/pkg/mux)
 * [net/http](https://golang.org/pkg/net/http/)
 * [encoding/json](https://golang.org/pkg/encoding/json/)
 * [fmt](https://golang.org/pkg/fmt/)
 * [strings](https://golang.org/pkg/strings/)

 ----
