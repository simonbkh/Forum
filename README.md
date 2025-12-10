# Web Forum Application

## 🎯 Project Overview

This project is a complete web forum application built using **Go** for the backend and **vanilla HTML/CSS/JavaScript** for the frontend. The forum allows for user communication, post categorization, and an interactive rating system (likes/dislikes). A core focus is on implementing secure **authentication** using cookies and managing data persistence with **SQLite**.

## ✨ Key Features

* **🔑 Authentication & Sessions:**
    * User **Registration** (Email, Username, Password).
    * User **Login** with session management using secure **HTTP Cookies**.
    * Password stored securely (Bonus: Hashed using **bcrypt**).
    * Session cookies include an expiration date.
* **💬 Communication:**
    * Registered users can create **Posts** and **Comments**.
    * Posts can be associated with one or more user-defined **Categories**.
    * All posts and comments are visible to all users (registered and non-registered).
* **👍 Rating System:**
    * Registered users can **Like** or **Dislike** any post or comment.
    * The total count of likes and dislikes is visible to all users.
* **⚙️ Filtering:** Users can filter the displayed posts by:
    * **Categories** (acting as sub-forums).
    * **Created Posts** (only available to the logged-in user).
    * **Liked Posts** (only available to the logged-in user).
* **💾 Data Management:**
    * Database stored using **SQLite**.
    * Database structured based on an **Entity Relationship Diagram (ERD)** to ensure performance and integrity.
    * Mandatory use of `SELECT`, `CREATE`, and `INSERT` SQL queries.
* **🐳 Containerization:** The entire application is packaged and deployed using **Docker**.

## 🛠️ Technology Stack

| Component | Technology | Constraints/Notes |
| :--- | :--- | :--- |
| **Backend** | **Go** | Core server-side logic, routing, and database interaction. |
| **Frontend** | **HTML, CSS, JavaScript** | **Strictly NO** frontend frameworks/libraries (React, Vue, Angular, etc.). |
| **Database** | **SQLite3** | Embedded database for storing users, posts, comments, and related data. |
| **Authentication** | **Cookies** | Used for session management. |
| **Security (Bonus)** | **bcrypt** | For password hashing. |
| **Containerization** | **Docker** | Used to create and run the application image. |

## 🏗️ Database Structure (Conceptual ERD)

To efficiently handle posts, comments, user interactions (likes/dislikes), and categories, the database schema will be designed around the following core entities:

* **Users:** Stores credentials and profile information.
* **Posts:** Stores the main content, creation details, and associated user.
* **Comments:** Stores replies to posts and associated user.
* **Categories:** Stores the available topics for posts.
* **PostCategories:** A many-to-many junction table linking Posts and Categories.
* **Reactions:** A single table or set of tables to track likes/dislikes on both Posts and Comments.
* **Sessions:** Stores active user sessions/cookies.



## 🚀 Getting Started (Docker)

The application is containerized for easy deployment.

### Prerequisites

* [Docker](https://docs.docker.com/get-docker/) installed.

### 1. Build the Docker Image

Navigate to the directory containing your `Dockerfile` (typically the project root or the backend directory) and build the image:

```bash
docker build -t go-forum-app .
