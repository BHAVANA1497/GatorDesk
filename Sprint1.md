# Sprint 1 – Report on the Tasks Accomplished

## Overview of the Project
Our project ‘Gator Desk’ is an online portal that aims at increasing the efficiency of handling basic maintenance tasks and fostering communication among residents in any residential community. Few of the key features of this project include - raising of complaints and maintenance requests by the residents, lost and found portal, portal for announcements of various community events and invitations for gatherings, and a discussion forum for the residents. It is basically a go-to place for the residents of the Gator Community, to cater for all their needs.

## Links pertaining to the project
- [GatorDesk Repo Link](https://github.com/BHAVANA1497/GatorDesk) 
- [GitHub – Discussion Forum for our project](https://github.com/BHAVANA1497/GatorDesk/discussions)
- [Actions and Workflows tab for our project on Git](https://github.com/BHAVANA1497/GatorDesk/actions)
- [Hyperlink to the user stories in this sprint](https://github.com/BHAVANA1497/GatorDesk/issues?q=is%3Aissue+is%3Aclosed)

## Backend tasks achieved - (GoLang)
- Developed and tested Announcement API - \createAnnouncement
-	Developed and tested Announcement API - \editAnnouncement
-	Developed and tested Announcement API - \listAllAnnouncements
-	Developed and tested Announcement API - \deleteAnnouncement

## REST API Documentation

### Announcement API
### 1.	\createAnnouncement API:

##### Request

Method: `POST`

##### Target URL: `"localhost:8181/createAnnouncement"`

##### Example:

Sample Body:

```json
{
    "announcementId": 2,
    "adminId":1,
    "announcementCategory":"Event Announcement",
    "announcementTitle":"New Year Celebrations",
    "venue":"Stoneridge Leasing Office",
    "eventDescription":"We are celebrating New Year's!",
    "createdAt":null,
    "updatedAt": null
}
```

Sample Response	:

```json
{
 "result": "Announcement created successfully"
}
```

##### Demonstration of the API in Postman:
![Create_api](https://user-images.githubusercontent.com/43131417/152629001-f84ffb54-99a4-46c4-b4fb-d43231e7ee6f.gif)
 


### 2.	\editAnnouncement API:

##### Request	

Method: `PUT`

##### Target URL: `"localhost:8181/editAnnouncement"`

##### Example:

Sample Body:

```json
{
    "announcementId":2,
    "adminId":1,
    "announcementCategory":"Event Update",
    "announcementTitle":"Announcement title edit",
    "venue":"Stoneridge Apartments",
    "eventDescription":"Event Description edit",
    "createdAt":null,
    "updatedAt": null
}
```

Sample Response	

```json
{
    "announcementId": 2,
    "adminId": 1,
    "announcementCategory": "Event Update",
    "announcementTitle": " Announcement title edit",
    "venue": "Stoneridge Apartments",
    "eventDescription": "Event Description edit",
    "createdAt": "2022-02-03T19:46:36.194787-05:00",
    "updatedAt": "2022-02-04T16:00:34.949808-05:00"
}
```

##### Demonstration of the API in Postman:
![edit_api](https://user-images.githubusercontent.com/43131417/152629010-2a6f35b1-268c-44a6-b085-353ca3cb9e9d.gif)
 

### 3.	\listAllAnnouncements API:

##### Request	

Method: `GET`

##### Target URL: `"http://localhost:8181/listAllAnnouncements"`

##### Example:

Sample Body: N/A

Sample Response	

```json
{
    "data": [
        {
            "announcementId": 1,
            "adminId": 1,
            "announcementCategory": "Event Announcement",
            "announcementTitle": "New Year Celebrations",
            "venue": "Stoneridge Leasing Office",
            "eventDescription": "We are celebrating New Year’s!",
            "createdAt": "2022-02-02T19:17:12.878628-05:00",
            "updatedAt": "2022-02-02T19:17:12.878628-05:00"
        },
        {
            "announcementId": 3,
            "adminId": 3,
            "announcementCategory": "Event Announcement-100",
            "announcementTitle": "Gator Farewell",
            "venue": "Stephen O’conel",
            "eventDescription": "Free food and drinks available until supply lasts",
            "createdAt": "2022-02-03T19:52:57.02995-05:00",
            "updatedAt": "2022-02-04T19:28:28.37535-05:00"
        },
        {
            "announcementId": 5,
            "adminId": 3,
            "announcementCategory": "Event Announcement-101",
            "announcementTitle": "Gator Resonance",
            "venue": "Stephen O&#39;conel",
            "eventDescription": "Free food and drinks available until supply lasts",
            "createdAt": "2022-02-04T20:13:57.047228-05:00",
            "updatedAt": "2022-02-04T20:13:57.047228-05:00"
        },
        {
            "announcementId": 6,
            "adminId": 3,
            "announcementCategory": "Event",
            "announcementTitle": "Gator Resonance",
            "venue": "Stephen O&#39;conel",
            "eventDescription": "Free food and drinks available until supply lasts",
            "createdAt": "2022-02-04T20:20:25.80535-05:00",
            "updatedAt": "2022-02-04T20:20:25.80535-05:00"
        },
        {
            "announcementId": 7,
            "adminId": 3,
            "announcementCategory": "Event",
            "announcementTitle": "Gator meetup",
            "venue": "Stephen O&#39;conel",
            "eventDescription": "Free food and drinks available until supply lasts",
            "createdAt": "2022-02-04T20:22:40.258026-05:00",
            "updatedAt": "2022-02-04T20:22:40.258026-05:00"
        },
        {
            "announcementId": 8,
            "adminId": 3,
            "announcementCategory": "Event",
            "announcementTitle": "Gator meetup-2",
            "venue": "Reitz Union",
            "eventDescription": "Free food and drinks available until supply lasts",
            "createdAt": "2022-02-04T20:25:05.9024-05:00",
            "updatedAt": "2022-02-04T20:39:48.215568-05:00"
        },
        {
            "announcementId": 9,
            "adminId": 3,
            "announcementCategory": "Event",
            "announcementTitle": "Gator meetup-2",
            "venue": "Reitz Union",
            "eventDescription": "Free food and drinks available until supply lasts",
            "createdAt": "2022-02-04T22:03:12.687761-05:00",
            "updatedAt": "2022-02-04T22:03:12.687761-05:00"
        }
    ]
}
```

##### Demonstration of the API in Postman:
![getall_api](https://user-images.githubusercontent.com/43131417/152629019-7b6c1204-62bc-4a42-818e-f964e581d4f5.gif)
 

### 4.	\deleteAnnouncement API:

##### Request	

Method: `DELETE`

##### Target URL: `"http://localhost:8181/deleteAnnouncement/4"`

##### Example:	

Sample Body:	N/A

Sample Response	

```json
{
   "data": true
}
```

##### Demonstration of the API in Postman:
![delete_api](https://user-images.githubusercontent.com/43131417/152629025-7759c2a2-3d94-4b3c-b1ae-eb2dd6ece6ee.gif)


## UI tasks accomplished - (Tech Stack: Angular 13, HTML, CSS, Type Script, JavaScript)

1.	As a part of the initial setup required for the frontend, we performed the following tasks:
  - Installed Node Package Manager 
  - Installed Node 14
  - Installed VSCode editor
2. Created a Landing Page. Created and enhanced the following components:
  -	Header component
  -	Footer component
  -	Carousal component
3. Created the following 3 pages for the following 3 functionalities:
  -	Announcements
  -	Lost and found
  -	Maintenance Requests 

The above functionalities have been implemented for both the use cases i.e., for user and admin roles. These pages can be reached my selecting the option on the Menu bar on the Home page.

4. Login, Signup buttons have been created for existing and new users.
5. Added routings for different pages.
6. Used Material Design to enhance the make and feel of the UI.

### UI Demo

