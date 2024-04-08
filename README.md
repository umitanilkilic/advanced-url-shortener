![GitHub License](https://img.shields.io/github/license/umitanilkilic/advanced-url-shortener)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/umitanilkilic/advanced-url-shortener)
[![Go Report Card](https://goreportcard.com/badge/github.com/umitanilkilic/advanced-url-shortener)](https://goreportcard.com/report/github.com/umitanilkilic/advanced-url-shortener)

# ⚠️ This project is currently under development. However, the current version is functional.
# URL SHORTENER - SPECS

- [ ] A sign up page
- [ ] A sign in page
- [ ] Create a new short URL page
- [x] A short URL should have a URL that it points to 
   - [x] A name
   - [x] An ID 
   - [x] A flag that keeps track whether the short URL is active and accessible - if this is set to false by the user it the short url should give 404
- [x] List URLS page
   - [x] Users should be able to delete individual lists here 
   - [x] Users should be able to bulk delete lists here
   - [x] Users should be able to order by creation date, update date, and clicks
- [x] A detail page 
   - [x] Users should be able to delete and edit the URL - AKA where it points to, if it is active, change the name etc. but the id should be read only
   - [x] It should contain stats on the URL
      - [x] How many times it has been clicked
      - [x] IP addresses of the clicks 
      - [x] Device info of the clicks 
      - [x] Mobile vs Desktop

# HOW TO USE API

## Auth API

### Login
- **POST** `/api/auth/login`
  - Authenticate user.
  - Request Body:
    ```json
    {
        "email": "umitanilkilic@gmail.com",
        "password": "pikacu123"
    }
    ```

### Register
- **POST** `/api/auth/register`
  - Register a new user.
  - Request Body:
    ```json
    {
        "email": "umitanilkilic@gmail.com",
        "password": "pikacu123"
    }
    ```

### URL Management

#### Get URL Details
- **GET** `/api/auth/urls/:urlId`
  - Retrieve details of a specific URL.

#### Get URLs (with Pagination)
- **GET** `/api/auth/urls`
  - Retrieve URLs with pagination.
  - Query Parameters:
    - `limit`: Number of URLs per page (default: 20).
    - `page`: Page number (default: 1).

#### Create URL
- **POST** `/api/auth/urls/`
  - Create a new shortened URL.
  - Request Body:
    ```json
    {
        "name": "example",
        "long": "https://example.com",
        "alias": "example",
        "expires_at": "2022-01-01T00:00:00Z"
    }
    ```

#### Delete URLs (Bulk)
- **DELETE** `/api/auth/urls/`
  - Delete multiple URLs at once.
  - Request Body:
    ```json
    {
        "urls_ids": ["url_id1", "url_id2"]
    }
    ```

#### Delete URL
- **DELETE** `/api/auth/urls/:urlId`
  - Delete a specific URL.

#### Update URL
- **PATCH** `/api/auth/urls/:urlId`
  - Update details of a specific URL.
  - Request Body:
    ```json
    {
        "active": true,
        "alias": "example",
        "name": "example",
        "long": "https://example.com",
        "expires_at": "2022-01-01T00:00:00Z"
    }
    ```

#### Get URL Stats
- **GET** `/api/auth/urls/stats/:urlId`
  - Get statistics of a specific URL.


# Environment Variables

This application relies on the following environment variables for proper functioning:

- **APP_ADDRESS:** Specifies the address where the application is accessible.
- **DATABASE_URL:** Specifies the URL for the database connection. (Postgres)
- **AUTH_SECRET:** Specifies the secret key used for authentication and session management.
- **REDIS_URL:** Specifies the URL for the Redis database connection. 

Please ensure that these environment variables are correctly configured for the successful operation of the application.



# TODO
- [ ] Dockerize the application
- [ ] Improve configuration management
- [ ] Add tests
- [ ] Add UI
- [ ] Redis integration for caching
