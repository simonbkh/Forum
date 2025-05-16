
```
testinforum
├─ README.md
├─ build
│  └─ Dockerfile
├─ cmd
│  └─ main.go
├─ go.mod
├─ go.sum
├─ internal
│  ├─ data
│  │  ├─ database
│  │  │  ├─ createtables.go
│  │  │  └─ database.go
│  │  ├─ database.db
│  │  ├─ modles
│  │  │  └─ models.go
│  │  ├─ queries
│  │  │  └─ queries.go
│  │  └─ schema.sql
│  ├─ logic
│  │  ├─ services
│  │  │  ├─ auth_service.go
│  │  │  ├─ category_service.go
│  │  │  ├─ comment_service.go
│  │  │  ├─ home_service.go
│  │  │  ├─ like_service.go
│  │  │  └─ post_service.go
│  │  ├─ utils
│  │  │  ├─ Error.go
│  │  │  ├─ atoi.go
│  │  │  ├─ bcrypt.go
│  │  │  ├─ session.go
│  │  │  └─ utils.go
│  │  └─ validators
│  │     ├─ post_validator.go
│  │     └─ user_validator.go
│  └─ presentation
│     ├─ handlers
│     │  ├─ auth_handler.go
│     │  ├─ category_handler.go
│     │  ├─ comment_handler.go
│     │  ├─ home_handler.go
│     │  ├─ like_handler.go
│     │  ├─ post_handler.go
│     │  └─ static.go
│     ├─ middleware
│     │  └─ middleware.go
│     ├─ router
│     │  └─ router.go
│     ├─ static
│     │  ├─ css
│     │  │  ├─ create-post.css
│     │  │  ├─ login-register.css
│     │  │  └─ styles.css
│     │  ├─ images
│     │  │  ├─ logo.png
│     │  │  ├─ logout.png
│     │  │  ├─ post.png
│     │  │  └─ profile.png
│     │  └─ js
│     │     ├─ comment.js
│     │     ├─ login-register.js
│     │     ├─ main.js
│     │     └─ pagination.js
│     └─ templates
│        ├─ auth
│        │  └─ login.html
│        ├─ errors
│        │  └─ error.html
│        ├─ layouts
│        │  ├─ index.html
│        │  ├─ nav_bar.html
│        │  └─ side_bar.html
│        ├─ post
│        │  ├─ create_post.html
│        │  └─ mypost.html
│        └─ templates.go
└─ todo

```