# TaskRocket

TaskRocket is being developed as a demonstration microservices application for managing personal tasks. This side
project serves as an opportunity for me to enhance my Golang skills and gain a deeper understanding of microservices
architecture.

As a DevOps engineer, I'm currently using my free time to learn Golang, a language that has intrigued me immensely.
While my coding expertise might not yet reflect that of a Google software engineer with two decades of experience,
a senior Golang engineer will be regularly reviewing the project to ensure its quality.

The aim for TaskRocket is to establish a microservices application that can be easily deployed on Kubernetes using
helm charts or kustomize. The application is designed to leverage the capabilities of the Istio service mesh fully.
Various design and technology choices have been made to familiarize myself with different technologies and concepts.

Feedback, suggestions, or contributions in any form are greatly appreciated.  
Feel free to contact me via email at github@sebapacuk.com

Below you'll find the current plan for the microservices in TaskRocket.

## Microservices Plan

### User Management Service

- Handles user registration, login, and profile management
- Uses a SQL database to store user credentials and profile data.

### Authentication Service

- Manages access and refresh tokens for users
- Validates tokens and authorizes users to access certain resources
- Provides auth middleware for other services to use
- Uses Redis to store refresh tokens

### Project Management Service:

- Handles creation, modification, and deletion of projects.
- Handles assigning users to projects and controls their roles (owner and participants).
- Uses a SQL database to store project information and user assignments.

### Task Management Service:

- Handles creation, modification, and deletion of tasks.
- Handles creation, modification, and deletion of categories.
- Handles creation, modification, and deletion of tags.
- Tasks are tied to projects and can be assigned to individual users.
- Uses a SQL database to store task information.

### Notification Service:

- Notifies users about changes in tasks or projects.
- Listens to relevant Kafka topics for events, then sends notifications via email, push notifications, or other
  channels.

### Project Collaboration Service

- Handles comments for project members to discuss their tasks.
- Store chats in a NoSQL database for flexibility.

### Task Prioritization Service:

- Set and Update Task Priorities
- Priority-Based Task Retrieval
- Priority Rules Management
- Notifications for Priority Changes (via Notification Service)
- Uses Redis cache to avoid querying the database for every request.

### User Invitation Service:

- Manages sending, accepting, or rejecting invitations to projects.
- Interacts with the User Management Service and Project Management Service.
- Sends notifications via the Notification Service.

### Audit Service

- Logs all actions performed by users for auditing purposes.
- It listens to Kafka for all user-performed events and uses a NoSQL database for storing these events.
- Events are deleted after a certain amount of time.

### Calendar Service

- Handle task recurrence patterns (e.g., repeating every week).
- Manage non-task events (e.g., project meetings).
- It would use a SQL database for storing calendar events.

### Search Service

- Allows users to search through their tasks and projects.
- Uses Elasticsearch or another full-text search system to provide this feature.

### Frontend

- OMG, Frontend is boring :) . I don't have much experience in frontend development.
  I think I'll simplify things and go with an Angular template that I can purchase and use.
  My main focus will be on the backend development. I understand that for smaller projects, Vue or React might be better
  options, but I have some familiarity with Angular, so I feel more comfortable using it.
