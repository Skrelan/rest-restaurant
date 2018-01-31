import requests
import pprint
import json
import time
"""
API demo code:

Usage:
- Create DB (queries to do so in project-root/db/migrations)
- Start up rest-restaurant on localhost:8000
- $ cd local
- $ python api-tests.py


NOTE:

- Yes, GOlang has built features and libraries (httptest,etc)
for integration and live testing.

- Yes, that is a better way to implement these
tests.

- But, since this is a demo code repository and
since one would manually test these endpoints for the demo,
why not have python do that for us?
"""

PORT = 8000
base_url = "http://localhost:{}{}"


def test_data(type_requested):
    users = [{
        "first_name": "Tom",
        "last_name": "Jerry",
        "phone": "4118717325"
    }, {
        "first_name": "Thomas",
        "last_name": "Jerico",
        "phone": "4118717325"
    }]
    restaurants = [{
        "name": "Sweet Tomato",
        "category": "Buffet",
        "venue": {
            "city": "San Mateo",
            "state": "CA",
            "street_address": "1234 Abc st.",
            "zip_code": "94123"
        }
    }, {
        "name": "DOSA",
        "category": "Indian",
        "venue": {
            "city": "San Mateo",
            "state": "CA",
            "street_address": "1234 Abc st.",
            "zip_code": "94123"
        }
    }, {
        "name": "Tacobell",
        "category": "Tex-Mex",
        "venue": {
            "city": "San Mateo",
            "state": "CA",
            "street_address": "1234 Abc st.",
            "zip_code": "94123"
        }
    }]
    ratings = [{
        "cost": 1,
        "food": 2,
        "cleanliness": 3,
        "service": 4,
        "restaurant_id": 3,
        "user_id": 4,
        "comments": "A post hangover miracle!"
    }, {
        "cost": 1,
        "food": 1,
        "cleanliness": 1,
        "service": 1
    }, {
        "cost": 1,
        "food": 1,
        "cleanliness": 1,
        "service": 1,
        "comments": "A brilliant take on fine fast food!"
    }]
    if type_requested == "users":
        return users
    if type_requested == "restaurants":
        return restaurants
    if type_requested == "ratings":
        return ratings


def GET(urls):
    for i, url in enumerate(urls):
        print("running GET request against {}".format(url))
        if raw_input("Run? (y/n) ").lower() == 'n':
            continue
        response = requests.get(url=url)
        print("\n Response : {0}".format(response))
        pprint.pprint(response.json())
        print("\n\n")


def POST(urls, data):
    if not data:
        return "ERROR: invalid test data"
    for i, url in enumerate(urls):
        print("running POST request against {}".format(url))
        pprint.pprint(data[i])
        if raw_input("Run?").lower() == 'n':
            continue
        response = requests.post(url=url, data=json.dumps(data[i]))
        print("\n Response : {0}".format(response))
        pprint.pprint(response.json())
        print("\n\n")


def PUT(urls, data):
    if not data:
        return "ERROR: invalid test data"
    for i, url in enumerate(urls):
        print("running PUT request against {}".format(url))
        pprint.pprint(data[i])
        if raw_input("Run?").lower() == 'n':
            continue
        response = requests.put(url=url, data=json.dumps(data[i]))
        print("\n Response : {0}".format(response))
        pprint.pprint(response.json())
        print("\n\n")


def test_GET_user():
    urls = [
        base_url.format(PORT, "/v1/users"),
        base_url.format(PORT, "/v1/users?id=1"),
        base_url.format(PORT, "/v1/users/1")
    ]
    GET(urls)


def test_GET_restaurants():
    urls = [
        base_url.format(PORT, "/v1/restaurants"),
        base_url.format(PORT, "/v1/restaurants?id=1"),
        base_url.format(PORT, "/v1/restaurants/1"),
        base_url.format(PORT, "/v1/restaurants?total_score=3.0"),
        base_url.format(PORT, "/v1/restaurants?name=DOSA")
    ]
    GET(urls)


def test_GET_ratings():
    urls = [
        base_url.format(PORT, "/v1/ratings"),
        base_url.format(PORT, "/v1/ratings/1"),
        base_url.format(PORT, "/v1/ratings?id=1"),
        base_url.format(PORT, "/v1/ratings?user_id=1"),
        base_url.format(PORT, "/v1/ratings?restaurant_id=1"),
        base_url.format(PORT, "/v1/ratings?user_id=1&restaurant_id=1")
    ]
    GET(urls)


def test_POST_user():
    urls = [base_url.format(PORT, "/v1/users")]
    users = test_data("users")
    POST(urls, users)


def test_POST_restaurant():
    urls = [base_url.format(PORT, "/v1/restaurants")]
    data = test_data("restaurants")
    POST(urls, data)


def test_POST_rating():
    urls = [base_url.format(PORT, "/v1/ratings")]
    data = test_data("ratings")
    POST(urls, data)


def test_PUT_user():
    urls = [
        base_url.format(PORT, "/v1/users"),
        base_url.format(PORT, "/v1/users/4")
    ]
    users = test_data("users")
    PUT(urls, [users[1] for x in range(len(urls))])


def test_PUT_restaurant_single_venue():
    urls = [
        base_url.format(PORT, "/v1/restaurants"),
        base_url.format(PORT, "/v1/restaurants/4")
    ]
    data = test_data("restaurants")
    PUT(urls, [data[1] for x in range(len(urls))])
    GET([base_url.format(PORT, "/v1/restaurants/4")])


def test_PUT_restaurant_all_venues():
    urls = [
        base_url.format(PORT, "/v1/restaurants?update_parent=true"),
        base_url.format(PORT, "/v1/restaurants/4?update_parent=true")
    ]
    data = test_data("restaurants")
    GET([base_url.format(PORT, "/v1/restaurants?name=DOSA")])
    PUT(urls, [data[2] for x in range(len(urls))])
    GET([base_url.format(PORT, "/v1/restaurants?name=Tacobell")])


def test_PUT_rating():
    urls = [
        base_url.format(PORT, "/v1/ratings"),
        base_url.format(PORT, "/v1/ratings/4"),
        base_url.format(PORT, "/v1/ratings/4"),
        base_url.format(PORT, "/v1/ratings?user_id=4&restaurant_id=3"),
        base_url.format(PORT, "/v1/ratings?user_id=4&restaurant_id=3")
    ]
    data = test_data("ratings")
    PUT(urls, [data[1], data[1], data[2], data[1], data[2]])


def run_tests():
    funcs = [
        test_GET_user, test_POST_user, test_PUT_user, test_GET_restaurants,
        test_POST_restaurant, test_PUT_restaurant_single_venue,
        test_PUT_restaurant_all_venues, test_GET_ratings, test_POST_rating,
        test_PUT_rating
    ]
    for func in funcs:
        if raw_input("Begin next test/demo? (y/n)\n Check : {} \t".format(
                func.__name__)).lower() == "n":
            continue
        func()
        print("-" * 50)


run_tests()
