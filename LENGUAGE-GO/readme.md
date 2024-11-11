# Steps to Run the Project

# 1. Clone the Repository

# Go to the terminal and type cd LANGUAGE-GO

# 2. Once the repository is cloned, build the Docker image with the following command:

docker build -t go-docker .

# 3. Run the Docker Container

docker run -p 8081:8081 go-docker

# Note: The application will be accessible at http://localhost:8081