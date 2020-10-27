# recruit-app-go-server
Go Server for the RecruitApp Recruitment and Selection using Psychometric Analysis Project

# Installation
1. Make sure you have the latest version of Go installed. This project uses version 1.10
2. Install gorrilla/mux using ``go get``
3. Clone the repository to a location of your choice
4. Setup the database by importing the sql script -- _found in pkg/config_ -- to your database server. This project uses ``MySQL Server 5.7``. Importing the database script should create the database and it's associated tables
5. Open up the project in your preferred text editor or IDE.
6. Locate the ``db.go`` file -- in ``pkg/config`` and replace the database credentials with your environment credentials
7. Open a terminal and cd to the location of the cloned repository 
8. Run ``go run *.go`` to run the entire content of the folder
