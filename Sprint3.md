# Sprint 3 – Report on the Tasks Accomplished

## Overview of the Project
Our project ‘Gator Desk’ is an online portal that aims at increasing the efficiency of handling basic maintenance tasks and fostering communication among residents in any residential community. Few of the key features of this project include - raising of complaints and maintenance requests by the residents, lost and found portal, portal for announcements of various community events and invitations for gatherings, and a discussion forum for the residents. It is basically a go-to place for the residents of the Gator Community, to cater for all their needs.

## Backend Developers:
- Nitya Vaishnavi Mamillapalle
- Akhil Maddukuri

## Frontend Developers:
- Bhavana Chinthalapally
- Harshitha Myadam

## Overview of the Tasks done in this sprint
- Lost and Found APIs in the Backend (Golang) - Nitya and Akhil
- Admin Management APIs in the Backend (Golang) - Nitya and Akhil
- Unit tests for all the APIs in Backend (Golang) - Nitya and Akhil
- Session Management code testing for User and Admin Testing in Backend - Nitya and Akhil
- Lost and Found Page for Admin in Frontend - Bhavana and Harshitha
- Lost and Found Page for User in Frontend - Bhavana and Harshitha
- Cypress tests and Unit tests - Bhavana and Harshitha

## Links pertaining to the project
- [GatorDesk Repo Link](https://github.com/BHAVANA1497/GatorDesk) 
- [Jira Board tasks for Sprint 3](https://github.com/BHAVANA1497/GatorDesk/projects/4)
- [GitHub – Discussion Forum for our project](https://github.com/BHAVANA1497/GatorDesk/discussions/76)
- [Hyperlink to the user stories in this sprint](https://github.com/BHAVANA1497/GatorDesk/issues?q=is%3Aissue+is%3Aclosed)
- [Actions and Workflows tab for our project on Git](https://github.com/BHAVANA1497/GatorDesk/actions)

## Backend tasks achieved - (GoLang) - [Akhil Maddukuri and Nitya Vaishnavi Mamillapalle]
- Developed and tested Lost&Found API - \postLostItem
- Developed and tested Lost&Found API - \listAllLostItems
- Developed and tested Lost&Found API - \postFoundItem
- Developed and tested Lost&Found API - \listAllFoundItems
- Developed and tested Lost&Found API - \listAllLostItemsByUserId
- Developed and tested Lost&Found API - \listAllLostItemsByUserId
- Developed and tested Admin Management API - \addAdmin
- Developed and tested Admin Management API - \deleteAdmin
- Developed and tested Admin Management API - \adminLogin
- Developed and tested Admin Management API - \listAllAdmins
- Developed and tested Admin Management API - \logout

## Backend Unit Tests - (GoLang)
We have written the following unit test cases:


Demo of the above mentioned Tests are after the REST API Documentation section.

## REST API Documentation

### User Authentication API
### 1.	\postLostItem API:

##### Request

Method: `POST`

##### Target URL: `"localhost:8181/postLostItem"`

##### Example:

Sample Body:

```json
{
	"lost_type": "jewellery",
	"description": "found near bus stop"
	"image_path": "https://dummy"
	"is_found": false
}
```

Sample Response	:

```json
{
 "result": "Lost item uploaded successfully!"
}
```

##### Demonstration of the API in Postman:
![postlostitem](https://user-images.githubusercontent.com/43131417/161362302-94765ae5-8b88-4631-bced-35a157d8f51f.jpg)


### 2.	\listAllLostItems API:

##### Request	

Method: `GET`

##### Target URL: `"localhost:8181/listAllLostItems"`

##### Example:

Sample Body: N/A

Sample Response	

```json
{
   "data": [
       {
           "ID": 0,
           "CreatedAt": "2022-04-01T21:21:44.043948-04:00",
           "UpdatedAt": "2022-04-01T21:21:44.043948-04:00",
           "DeletedAt": null,
           "lost_id": 0,
           "UserId": 0,
           "lost_type": "jewellery",
           "description": "found near busstop",
           "image_path": "https://dummy",
           "is_found": false,
           "found_id": 0
       },
       {
           "ID": 0,
           "CreatedAt": "2022-04-01T21:26:59.836429-04:00",
           "UpdatedAt": "2022-04-01T21:26:59.836429-04:00",
           "DeletedAt": null,
           "lost_id": 1,
           "UserId": 0,
           "lost_type": "book",
           "description": "found near office",
           "image_path": "https://dummy",
           "is_found": false,
           "found_id": 0
       }
   ]
}
 
```

##### Demonstration of the API in Postman:
![ListAllLostItems](https://user-images.githubusercontent.com/43131417/161362482-95edb751-85cc-4ab7-ac83-dedf886ec611.jpg)


### 3.	\postFoundItem API:

##### Request

Method: `POST`

##### Target URL: `"localhost:8181/postFoundItem"`

##### Example:

Sample Body:

```json
{
	"found_id": 1,
	"found_type": "jewellery",
	"description": "found on Sunday",
	"image_path": "https://dummy"
}
```

Sample Response	:

```json
{
 "result": "Found item has been updated successfully in the database!"
}

```

##### Demonstration of the API in Postman:
![postFoundItem](https://user-images.githubusercontent.com/43131417/161362991-91efeb2b-1b52-4ce7-a883-4f1d11f7d0d4.jpg)

### 4.	\listAllFoundItems API:

##### Request

Method: `GET`

##### Target URL: `"localhost:8181/listAllFoundItems"`

##### Example:

Sample Body: N/A 

Sample Response	:

```json
{
    "data": [
        {
            "ID": 0,
            "CreatedAt": "2022-04-01T22:13:58.785071-04:00",
            "UpdatedAt": "2022-04-01T22:13:58.785071-04:00",
            "DeletedAt": null,
            "found_id": 1,
            "UserId": 0,
            "found_type": "JEwellery",
            "description": "Found on sunday",
            "image_path": "https://"
        },
        {
            "ID": 0,
            "CreatedAt": "2022-04-01T22:15:56.009151-04:00",
            "UpdatedAt": "2022-04-01T22:15:56.009151-04:00",
            "DeletedAt": null,
            "found_id": 2,
            "UserId": 0,
            "found_type": "JEwellery",
            "description": "Found on sunday",
            "image_path": "https://"
        }
    ]
}

```

##### Demonstration of the API in Postman:
![LIST ALL FOUND ITEMS](https://user-images.githubusercontent.com/43131417/161363188-a0bcffc5-c81e-49d9-adb1-e32771207a56.jpg)

### 7.	\addAdmin API:

##### Request

Method: `POST`

##### Target URL: `"localhost:8181/addAdmin"`

##### Example:

Sample Body: 

```json
{
    "adminname": "test_user1",
    "password": "test",
    "firstname": "test",
    "lastname": "user",
    "phone": "+15714643661"   
}
```

Sample Response	:

```json
{
 "result": "Successfully added as admin"
}

```

##### Demonstration of the API in Postman:
![add admin1](https://user-images.githubusercontent.com/43131417/161363827-32f6cb8f-a1c8-4405-a8ae-4af8ccfa808a.jpg)

In Database:

![add admin 2](https://user-images.githubusercontent.com/43131417/161363843-165ffabb-dd01-470d-a864-14de7f450288.jpg)


### 8.	\deleteAdmin API:

##### Request

Method: `DELETE`

##### Target URL: `"localhost:8181/deleteAdmin/1"`

##### Example:

Sample Body: N/A

Sample Response	:

```json
{
 "data": true
}

```

##### Demonstration of the API in Postman:

![delete admin](https://user-images.githubusercontent.com/43131417/161364050-b481d339-b9d7-49e9-b1b6-e398b1ed04d0.jpg)



### 9.	\adminLogin API:

##### Request

Method: `POST`

##### Target URL: `"localhost:8181/adminLogin"`

##### Example:

The following 3 scenarios for admin login has been tested:

1. Admin login - direct case

![admin succ login](https://user-images.githubusercontent.com/43131417/161364193-3ec9ce26-14c7-4a93-a523-1fc7a3d8473c.jpg)


2. Hitting Admin Login when already logged in

![admin already logged in](https://user-images.githubusercontent.com/43131417/161364240-0401a55d-b41f-4db5-95b4-277e864247e2.jpg)

### 10.	\listAllAdmins API:

##### Request

Method: `GET`

##### Target URL: `"localhost:8181/listAllAdmins"`

##### Example:

Sample Body: N/A

Sample Response	:

```json
{
   "data": [
       {
           "ID": 1,
           "CreatedAt": "2022-04-01T22:26:49.372947-04:00",
           "UpdatedAt": "2022-04-01T22:26:49.372947-04:00",
           "DeletedAt": null,
           "adminname": "Akhil",
           "password": "password",
           "firstname": "Akhil",
           "lastname": "Maddukuri",
           "phone": "+15714643662"
       }
   ]
}
```
### 11.	\logout API:

##### Request

Method: `POST`

##### Target URL: `"localhost:8181/logout"`

##### Example:

The following scenarios for logout has been tested for User and Admin:

1. Successful logout:

![logout succ](https://user-images.githubusercontent.com/43131417/161364666-f2a73801-94a8-4cf7-acd0-f304c857ade5.jpg)

2. Trying to logout without logging in:

![logout 2](https://user-images.githubusercontent.com/43131417/161364770-f45f4bf5-447e-47ee-898f-c9277413ff87.jpg)

 

### Backend Demo:

- Admin Already logged in:

![admin_login_already_loggedin](https://user-images.githubusercontent.com/43131417/161365108-fab922c7-9d44-46e4-8619-2b9760ea9102.gif)

- Delete Admin:

![delete_admin](https://user-images.githubusercontent.com/43131417/161365117-3e11888e-3ad6-4c84-af2e-21ffa748948d.gif)

- Admin Login:

![admin_login](https://user-images.githubusercontent.com/43131417/161365122-44d3f038-3079-4b00-b4d0-31f28a2ca4bf.gif)

- Admin login when another user already logged in:

![admin_login_when_another_loggedin](https://user-images.githubusercontent.com/43131417/161365138-6219e3dd-16a9-485c-9f07-73e431044769.gif)

- Add Admin:

![add_admin](https://user-images.githubusercontent.com/43131417/161365141-b13542c8-2b44-4616-a63e-418977246fc6.gif)

- Post Lost Item:

![post_lostitem](https://user-images.githubusercontent.com/43131417/161365148-4ded7ebd-0c8d-43b6-98d0-ba860a69aa7f.gif)


## UI tasks accomplished - (Tech Stack: Angular 13, HTML, CSS, Type Script, JavaScript)

1. Created an admin component that displays all the lost and found items
2. Developed a backend call that fetches the entire data and displays in the UI to the end user
3. Developed an additional feature for admin where the admin can link the lost item with the found item.
4. Developed a webpage where the user can request his/her lost or found items to be displayed.
5. Built unit test cases for all the angular components using Karma and Jasmine
6. Performed automated testing using cypress for all the functionalitied developed in this sprint.
7. Integrated the frontend and backend of the functionality developed in this sprint.

### Steps to run 

#### Running the project
1. Clone the project from https://github.com/BHAVANA1497/GatorDesk.git
2. run the below steps in the terminal
    -- npm i
    -- ng serve --o

#### Cypress Tests
Steps:-


npm install cypress


In package.json :-

"scripts": {
    [....]
    "cypress:open": "cypress open",
    "cypress:run": "cypress run",
    "cypress": "cypress open"
  },
  
To Run :-


npm run cypress


#### Jest tests


npm install jest jest-preset-angular --save-dev


To Run:-


npm run test


## UI Demo

1. UI for Lost and Found for Admin:

![lost-found-admin](https://user-images.githubusercontent.com/43131417/161365014-d18c7186-99be-406b-970a-9446779d9275.gif)

2. UI for Lost and Found for User:

![lost-found-user](https://user-images.githubusercontent.com/43131417/161365028-182091eb-9f67-43ab-b2b9-da2e67355dae.gif)

3. Lost and Found Cypress Tests:

![lost-found-cypress-test](https://user-images.githubusercontent.com/43131417/161365049-31a45f03-8b5e-4c41-a66b-9652870be43d.gif)

4. Lost and Found Unit Tests:

![lost-found-unit-test](https://user-images.githubusercontent.com/43131417/161365061-ae148c4e-9a2f-4ae4-9686-6200880b6fad.gif)


### Video Demo Link:
https://youtu.be/tDzvsLLHSPU

