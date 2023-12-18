# AWS Lambda Go YouTube Search

This AWS Lambda project, developed in Go, utilizes API Gateway to handle HTTP requests for YouTube video searches. The Lambda function processes both GET and POST requests, providing a versatile interface for users to search and obtain YouTube video links.

https://uc1av5kjik.execute-api.eu-north-1.amazonaws.com/prod/search

Project Structure
Main Package (main.go):# AWS Lambda Go YouTube Search

This AWS Lambda project, developed in Go, utilizes API Gateway to handle HTTP requests for YouTube video searches. The Lambda function processes both GET and POST requests, providing a versatile interface for users to search and obtain YouTube video links.

## Project Structure

- **Main Package (`main.go`):**

  - Initializes the Gin web framework.
  - Handles incoming requests through API Gateway and routes them to appropriate handlers.

- **Search Package (`search.go`):**

  - Interacts with the YouTube Data API to search for videos based on user queries.
  - Processes links and extracts video information.

- **Handlers Package (`handlers.go`):**

  - Defines handlers for both GET and POST requests.
  - Handles incoming requests, validates parameters, and triggers YouTube searches.

- **Configuration Package (`config.go`):**

  - Manages configuration settings for the YouTube Data API.

## Usage

- **Search Endpoint:**
  - **GET:** `/search?q={query}`
    - Retrieves YouTube video link for the specified query.
  - **POST:** `/search`
    - Accepts a JSON array of links and returns a map of processed video links.

## Deployment

The project is designed to be deployed using the Serverless Framework. The `serverless.yml` file defines the AWS Lambda function and its associated API Gateway endpoints.

