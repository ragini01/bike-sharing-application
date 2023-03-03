# bike-sharing-application

This is a bike sharing service that allows users to rent and return bicycles. The service consists of a backend API built with Golang and a frontend web application built with React and TypeScript. The backend uses MySQL as its database.

## **Prerequisites**
    . Docker
    . Docker Compose
    
## **Installation**
   1. Clone the repository
      > git clone https://github.com/<username>/bike-sharing-application.git
      > cd bike-sharing-application
   
   2. Start the service using Docker Compose
      > docker-compose up
    The frontend should be available at http://localhost:3000 and the backend API should be available at http://localhost:8080. 
  
## **API Documentation**
   The API documentation is available in the swagger.yml file. It is not accessible in the browser due to some error internally.
  
## **Usage**
   1. Frontend web application accessible at http://localhost:3000.
   2. View the list of available bicycles and click on a bicycle from avaialble lists to view its details.
   3. Click on the "Rent bike" button to rent a bicycle.
   4. If you have rented a bicycle, you will not be able to rent another one until already rented bicycle is not returned.
   5. To return a bicycle, click on the bicycle and click on the "Return bike" button.
  
## **Development**
   To run the backend API locally, use below commands:
    > cd backend
    > go run main.go
  
  To run the frontend web application locally, use below commands:
    > cd frontend
    > npm install
    > npm start
  
