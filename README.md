# FlowManagerAPI
REST API with Go for basic project management functionalities, similar to tools like Jira. The API allows for the creation, management, and tracking of projects and tasks.
# Features
- Workflow Management: Create, update, and delete workflows.
- Persistence: Utilizes PostgreSQL to store workflow data and states.
- RESTful Interface: Provides RESTful endpoints for interacting with the API, facilitating integration with other services and applications.
- Authentication and Authorization: Secures endpoints with JWT (JSON Web Token) authentication to ensure that only authorized users can access certain resources.
# Proyect structure:
![image](https://github.com/user-attachments/assets/6f4e2e66-158e-4f09-9a0a-cfa7b2742d1c)

# Database diagram:
![image](https://github.com/user-attachments/assets/10adb936-1c3a-47ea-8c2b-c741db27dc98)

# Deployment
- Build the Dockerfile with the tag fmapi:2.0
- Configure the enviroment variables that the api and postgres are going to use
- Ex:
  DB_NAME=db
  DB_USER=user
  DB_PASSWORD=secret
  JWT_SECRET=tdlZgTNhletbNDw/EnaT4Q7d00SDVXppSLlAE0AkJ/E=
- Run de dockercompose file.
- You can deploy the application in any cloud provider.

# Instalation
- Clone the repository: git clone https://github.com/rickalon/FlowManagerAPI.git
- Navigate to the project directory
- make build

# Test
- make test



