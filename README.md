# todo-microservice


## DevFiles

> You can use devfiles to automate and simplify your development process.

When it comes to Devfile and modern cloud-native development, the terms Inner Loop and Outer Loop are used to describe different types of workflows for software development. These loops represent the stages of development, testing, and deployment, and help distinguish between actions that happen frequently (inner loop) and those that happen less frequently (outer loop).

### Inner Loop

The Inner Loop refers to the cycle of local development and testing that developers go through on a regular basis. It includes activities that are repeated frequently, often multiple times a day. These actions are typically focused on:

- Writing Code
- Building the Application
- Running Tests
- Debugging

#### Get started 

```bash
# Initializes a new ODO (OpenDevOps) project in the current directory,
# setting up the necessary configuration files for the project.
odo init

# Starts the ODO development environment using Podman as the container platform.
# This command will build and deploy the application inside a Podman container,
# allowing developers to run and test their applications locally within a containerized environment.
odo dev --platform podman

# Deletes the component (application or service) from the ODO project,
# stopping and removing any associated containers, resources, or configurations.
# In this case, the deletion is performed using Podman as the platform.
odo delete component --platform podman
```

If you are using GO pay attention to the local version and the one used in the devfile. 
### Outer Loop

The Outer Loop refers to the activities that happen less frequently but are still crucial to the overall development lifecycle. These tasks are more focused on the integration and deployment aspects of the application. Key activities in the outer loop typically include:

- Building and Deploying to Staging or Production
- Integration Testing
- End-to-End Testing
- Deploying Microservices
- Handling CI/CD Pipelines

While the inner loop is about the local development process, the outer loop involves tasks that ensure the application is properly integrated and deployed into a more extensive environment, such as staging or production.

In Devfile terms, the outer loop might involve integration with continuous integration (CI) and continuous deployment (CD) systems, or deploying the application into cloud-native environments like Kubernetes or OpenShift. Devfiles can define commands to build and deploy the app on these platforms and integrate them into the broader outer loop processes.

