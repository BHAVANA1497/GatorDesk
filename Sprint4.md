

![unit_tests](https://user-images.githubusercontent.com/46106443/164350900-b841619f-cb26-4c3a-ba4d-a9748a9f307b.gif)

The following video shows the code and execution for the enitr unit tests suite present for the backend

![unit_tests_demo](https://user-images.githubusercontent.com/46106443/164353014-d7b9aee1-e539-4cd4-8312-a8a07805363e.gif)






## REST API Documentation

### Session Management for Lost and Found APIs
### 1.	/postLostItem API:

Trying to post the item without any logged in user

##### Request

Method: `POST`

##### Target URL: `"localhost:8181/postLostItem"`

##### Example:

Sample Body:

```json
{
    "lost_type": "Jewellery",
    "description": "lost near pool",
    "image_path": "https://image"
}
```

Sample Response	:

```json
{
  "result": "User not loggedin"
}
```

##### Demonstration of the API in Postman:
<img width="1016" alt="image" src="https://user-images.githubusercontent.com/46106443/164355561-6e2f6053-438c-4124-a1bd-6a2e0fa4b418.png">



### 2.	/listAllLostItems API:

Trying to request without any logged in user

##### Request	

Method: `GET`

##### Target URL: `"localhost:8181/listAllLostItems"`

##### Example:

Sample Body: N/A

Sample Response	

```json
{
  "result": "User not loggedin"
}
 
```


##### Demonstration of the API in Postman:
<img width="1020" alt="image" src="https://user-images.githubusercontent.com/46106443/164355301-a3371819-1bdd-4ad3-a197-f36b422e7b5f.png">


Trying to request when the user is not admin

##### Request	

Method: `GET`

##### Target URL: `"localhost:8181/listAllLostItems"`

##### Example:

Sample Body: N/A

Sample Response	

```json
{
  "result": "login as admin to view this page"
}
 
```

##### Demonstration of the API in Postman:
<img width="1026" alt="image" src="https://user-images.githubusercontent.com/46106443/164355253-056be273-a831-43b6-ac23-0439a9299bd7.png">



### 3.	/listAllLostItemsByUserId API:

Trying to request without any logged in user

##### Request	

Method: `GET`

##### Target URL: `"localhost:8181/listAllLostItemsByUserId"`

##### Example:

Sample Body: N/A

Sample Response	

```json
{
  "result": "User not loggedin"
}
 
```


##### Demonstration of the API in Postman:
<img width="1008" alt="image" src="https://user-images.githubusercontent.com/46106443/164355906-b06efc33-e835-45bf-a7d1-aa459b552871.png">

### 4.	/listLostItemById/:id API:

Trying to request without any logged in user

##### Request	

Method: `GET`

##### Target URL: `"localhost:8181/listLostItemById/:id"`

##### Example:

Sample Body: N/A

Sample Response	

```json
{
  "result": "User not loggedin"
}
 
```


##### Demonstration of the API in Postman:
<img width="1010" alt="image" src="https://user-images.githubusercontent.com/46106443/164356234-cec3d93c-c5eb-4dc5-87c4-a756c01040c3.png">


Demo video for Lost items sessions management

![lost-sessions](https://user-images.githubusercontent.com/46106443/164356368-700ae9a0-b6ad-420b-a2f0-110bc6542fac.gif)


### 5.	\postFoundItem API:

Trying to request without any logged in user

##### Request

Method: `POST`

##### Target URL: `"localhost:8181/postFoundItem"`

##### Example:

Sample Body:

```json
{
    "found_type": "Jewellery",
    "description": "found somewhere",
    "image_path": "na"
}
```

Sample Response	:

```json
{
 "result": "User not loggedin"
}

```

##### Demonstration of the API in Postman:
<img width="1004" alt="image" src="https://user-images.githubusercontent.com/46106443/164357274-2190a35e-f4bd-4e8a-9501-fd815cefa063.png">


### 6.	\listAllFoundItemsByUserId API:

Trying to request without any logged in user

##### Request

Method: `GET`

##### Target URL: `"localhost:8181/listAllFoundItemsByUserId"`

##### Example:

Sample Body: N/A 

Sample Response	:

```json
{
   "result": "User not loggedin" 
}

```

##### Demonstration of the API in Postman:
<img width="1020" alt="image" src="https://user-images.githubusercontent.com/46106443/164357605-b960e48e-f36d-444d-828e-43b4455b11f0.png">

### 7.	\listFoundItemById/1 API:

Trying to request without any user logged in

##### Request

Method: `GET`

##### Target URL: `"localhost:8181/listFoundItemById/1"`

##### Example:

Sample Body: N/A

Sample Response	:

```json
{
 "result": "User not loggedin"
}

```

##### Demonstration of the API in Postman:
<img width="1022" alt="image" src="https://user-images.githubusercontent.com/46106443/164357963-a64ccdbb-8dd7-4ca4-9d2d-981680335b25.png">

Demo video for Lost items sessions management



![found-sessions](https://user-images.githubusercontent.com/46106443/164358192-209ecc6d-2ccb-48f6-9f61-e163d8bb5ccc.gif)



### 8.	/linkLostFound API:

##### Request

Method: `PUT`

##### Target URL: `"localhost:8181//linkLostFound/1"`

##### Example:


Sample Body	:

```json
{
    "ID": 1,
    "lost_type": "jewellery",
    "description": "found at pool",
    "image_path": "na",
    "is_found": true,
    "found_id":1
}
```

Sample Response	:

```json
{
    "ID": 1,
    "CreatedAt": "2022-04-20T02:26:48.886288-04:00",
    "UpdatedAt": "2022-04-20T22:52:39.634075-04:00",
    "DeletedAt": null,
    "UserId": 1,
    "lost_type": "jewellery",
    "description": "found at pool",
    "image_path": "na",
    "is_found": true,
    "found_id": 1
}

```

##### Demonstration of the API in Postman:

<img width="1013" alt="image" src="https://user-images.githubusercontent.com/46106443/164362525-ab7b0267-d1fa-4013-952b-9c8f3b1b97a2.png">



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
