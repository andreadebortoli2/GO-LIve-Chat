##### A web app with live chat

This is a basic app where you can register, login and live chat, the admin can also edit user to give or remove superior access permission.

Once started you can chat simultaneously between different browsers.

To try the app:
- in editor terminal start the app with the command ```air``` 
- open multiple browsers at [localhost](http://localhost:8080/)
- register an account or login with one of the already existing profiles
  - username: admin
    - email: admin@admin.com
    - password: password
  - username: moderator
    - email: moderator@moderator.com
    - password: password
  - username: user
    - email: user@user.com
    - password: password
- chat simultaneously between the browsers
- admin can access to users list to delete user or modify the access level



---


- built in GO version 1.23.1
- use [Air - Live reload for Go apps](https://github.com/air-verse/air) dev util for live-reload
- use the [chi router](https://go-chi.io/#/) to manage routing
- use [NoSurf](https://github.com/justinas/nosurf) to manage CSRF Tokens
- use [SQLite3](https://sqlite.org/) as database
- use [GORM](https://gorm.io/) ORM library for GO
- use [Pure-Go SQLite driver for GORM](https://github.com/glebarez/sqlite) SQLite3 driver for GORM written all in GO
- use [govalidator](https://github.com/asaskevich/govalidator) to easily manage input validation
- use [Bootstrap](https://getbootstrap.com/) for front end styling
- use [</> htmx](https://htmx.org/) to execute form requests without reloading the page
- use [Gorilla WebSocket](https://github.com/gorilla/websocket) for chat live load
- use [Gorilla Sessions](https://github.com/gorilla/sessions) to manage the session