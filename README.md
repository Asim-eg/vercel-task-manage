# Asim Task Manager

## Overview
[asim-task-manage](https://asim-task-manage.vercel.app/) is a web-based task management application designed for efficient task organization and tracking. Developed using Go and the Gin framework, with MongoDB for data storage, it provides functionalities such as task creation, updating, deletion, and viewing. The application is deployed on [Vercel](https://asim-task-manage.vercel.app/) and offers an intuitive user interface for managing tasks.

## Features
- **Task Creation:** Users can add new tasks with details like name, description, and author.
- **Task Viewing:** Provides a view of all tasks in a neatly organized list.
- **Task Updating:** Allows users to update task details, including completion status.
- **Task Deletion:** Users can delete tasks that are no longer needed.
- **Comments:** Users can add comments to tasks, enhancing collaboration.
- **Real-time Updates:** The application updates tasks in real-time as changes are made.

## Technology Stack
- **Frontend:** Vanilla JavaScript, HTML, CSS
- **Backend:** Go (Gin framework)
- **Database:** MongoDB
- **Deployment:** [Vercel](https://asim-task-manage.vercel.app/)

## Setup and Installation
1. Clone the Repository:
   ```bash
   git clone https://github.com/Asim-eg/vercel-task-manage.git
   cd vercel-task-manage
   ```
# MongoDB Setup
1. Set up a MongoDB Atlas account and create a cluster.
2. Add your MongoDB URI in a .env file as MONGODB_URI.

# Running the Application
1. Start the server using `go run api/entrypoint.go`.
2. Open `public/index.html` in a web browser to access the frontend.

# API Endpoints
- **POST /api/create:** Create a new task.
- **GET /api/get:** Retrieve all tasks.
- **PUT /api/update/:id:** Update an existing task.
- **DELETE /api/delete/:id:** Delete a task.
- **GET /api/get/:id:** Get a single task by ID.

# Directory Structure
- `/api`: Backend Go files.
- `/public`: Frontend HTML, CSS, JavaScript files.
- `vercel.json`: Vercel deployment configuration.

# Deployment
Deployed on Vercel, the application automatically updates upon code changes pushed to the repository.

# Contributing
Contributions to enhance asim-task-manage are welcome. Follow these steps:
1. Fork the repo.
2. Create a feature branch (`git checkout -b feature/AmazingFeature`).
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`).
4. Push to the branch (`git push origin feature/AmazingFeature`).
5. Open a pull request.

# License
Distributed under the MIT License. See [LICENSE](https://github.com/Asim-eg/vercel-task-manage/blob/main/LICENSE) for more information.

# Contact
Project Repository: [GitHub](https://github.com/Asim-eg/vercel-task-manage)
