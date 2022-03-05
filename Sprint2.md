# Sprint 2 – Report on the Tasks Accomplished

## Overview of the Project
Our project ‘Gator Desk’ is an online portal that aims at increasing the efficiency of handling basic maintenance tasks and fostering communication among residents in any residential community. Few of the key features of this project include - raising of complaints and maintenance requests by the residents, lost and found portal, portal for announcements of various community events and invitations for gatherings, and a discussion forum for the residents. It is basically a go-to place for the residents of the Gator Community, to cater for all their needs.

## Overview of the Tasks done in this sprint
- Integration of Frontend and Backend for Announcement APIs
- User Authgentication APIs in the Backend
- Unit tests for all the APIs in Backend (Golang)
- Session Management code in Golang
- Code for Cookies in backend (Golang)
- Cypress and Jest for Frontent

## Links pertaining to the project
- [GatorDesk Repo Link](https://github.com/BHAVANA1497/GatorDesk) 
- [Jira Board tasks for Sprint 2](https://github.com/BHAVANA1497/GatorDesk/projects/2)
- [GitHub – Discussion Forum for our project](https://github.com/BHAVANA1497/GatorDesk/discussions)
- [Actions and Workflows tab for our project on Git](https://github.com/BHAVANA1497/GatorDesk/actions)
- [Hyperlink to the user stories in this sprint](https://github.com/BHAVANA1497/GatorDesk/issues?q=is%3Aissue+is%3Aclosed)

## Backend tasks achieved - (GoLang)
- Developed and tested User Authentication API - \login
- Developed and tested User Authentication API - \signUp
- Developed and tested User Authentication API - \deleteUser
- Developed and tested User Authentication API - \logout

## Backend Unit Tests - (GoLang)
We have written the following unit test cases:
1)	TestListAllAnnouncementsView
2)	TestAdminCreateAnnouncementPassCase
3)	TestAdminCreateAnnoncementFailCase
4)	TestAdminEditAnnouncementPassCase
5)  TestLoginPassCase
6)  TestLoginFailCase
7)  TestRegisterPassCase
8)  TestRegisterFailCase
9)  TestDeleteUserPassCase
10) TestDeleteUserFailCase

Demo of the above mentioned Tests after the REST API Documentation section.

## REST API Documentation

### User Authentication API
### 1.	\login API:

##### Request

Method: `POST`

##### Target URL: `"localhost:8181/login"`

##### Example:

Sample Body:

```json
{
    "username": "nitya_v",
    "password": "IForgotMyPassword"
}
```

Sample Response	:

```json
{
 "result": "Successful Login!"
}
```

##### Demonstration of the API in Postman:

<img width="452" alt="image" src="https://user-images.githubusercontent.com/43131417/156868271-5a49a52a-7f85-4343-98be-bb1d0c53f897.png">


### 2.	\signUp API:

##### Request	

Method: `POST`

##### Target URL: `"localhost:8181/signUp"`

##### Example:

Sample Body:

```json
{
    "username": "nitya_v",
    "password": "IForgotMyPassword",
    "firstname": "Nitya Vaishnavi",
    "lastname": "Mamillapalle",
    "phone": "1234567890",
    "aptNo": "GG-333"
}
```

Sample Response	

```json
{
    "result": "Succesfully signed up the user!"
}
```

##### Demonstration of the API in Postman:

<img width="452" alt="image" src="https://user-images.githubusercontent.com/43131417/156868265-a3f25e9d-ff19-4411-974b-07252e136b2d.png">


### 3. \deleteUser API:

##### Request	

Method: `DELETE`

##### Target URL: `"http://localhost:8181/deleteUser/3"`

##### Example:

Sample Body: N/A

Sample Response	

```json
{
    "data": true
}
```

##### Demonstration of the API in Postman:
<img width="478" alt="image" src="https://user-images.githubusercontent.com/43131417/156868289-e19dce28-13dd-4038-af00-14ace41240e3.png">


### Backend Demo - Unit Test Cases in GoLang:

![delete_api-2](https://user-images.githubusercontent.com/43131417/156868531-a396f812-e788-4f9e-9116-b5914cb8c638.gif)


## UI tasks accomplished - (Tech Stack: Angular 13, HTML, CSS, Type Script, JavaScript)

1) Refactored UI using material UI and created login page .
2) Integrated frontend and backend for annoucement and admin annoucement page
3) Added Jest unit testcases for UI
4) Using Cypress framework added frontend unit test cases.

### UI Demo

1. Home Page:

![home_page_new](https://user-images.githubusercontent.com/43131417/156868427-005d4657-dfcd-4f06-80de-29bdbe1e3ca9.gif)


2. Cypress:

![cypress](https://user-images.githubusercontent.com/43131417/156868423-201c3d73-bcb9-4bd2-bbc2-9c35c89660a2.gif)


3. Jest-UI-Angular:
![jest-UI-angular](https://user-images.githubusercontent.com/43131417/156868468-bb12bd47-c83c-496c-b3e1-b04e8d67accc.gif)


### Video Demo Link:
https://youtu.be/tDzvsLLHSPU

### Steps to run:

#### Cypress Tests
Steps:-

```
npm install cypress
```

In package.json :-

"scripts": {
    [....]
    "cypress:open": "cypress open",
    "cypress:run": "cypress run",
    "cypress": "cypress open"
  },
  
To Run :-

```
npm run cypress
```

#### Jest tests

```
npm install jest jest-preset-angular --save-dev
```

To Run:-

```
ng test
```

#### Go tests:

```
cd backend
cd webapp
go mod tidy
go test -v
```
