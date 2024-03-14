![GitHub License](https://img.shields.io/github/license/umitanilkilic/advanced-url-shortener)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/umitanilkilic/advanced-url-shortener)
[![Go Report Card](https://goreportcard.com/badge/github.com/umitanilkilic/advanced-url-shortener)](https://goreportcard.com/report/github.com/umitanilkilic/advanced-url-shortener)

### URL Shortener

![Example](https://i.hizliresim.com/t4sc3wp.gif)

This is a advanced URL shortener service implemented in Go using the Fiber web framework.  

* It provides a simple API for shortening long URLs and redirecting to the original long URL.  
* It also includes a rate limiter to prevent abuse and a simple HTML template for redirection page.




### Installation

1. Make sure you have Go installed on your machine.
2. Clone the repository:

   ```
   git clone https://github.com/umitanilkilic/advanced-url-shortener.git
   ```

3. Navigate to the project directory:

   ```
   cd advanced-url-shortener
   ```

4. Build and run the project:

   ```
   go run .
   ```

5. PostgreSQL Setup:
Create a database and a table with the following schema:

   ```sql
   CREATE TABLE IF NOT EXISTS public.shorturl
   (
      url_id integer NOT NULL,
      long_url character varying(255) COLLATE pg_catalog."default",
      created_at timestamp without time zone,
   );
   ```

### Usage

1. Send a POST request to `/shorten` with a JSON payload containing the long URL you want to shorten:

   ```json
   {
       "longUrl": "https://example.com/very/long/url/that/you/want/to/shorten"
   }
   ```

2. The service will respond with a JSON containing the shortened URL ID:

   ```json
   {
       "shortUrlId": "1234567"
   }
   ```

3. To access the original long URL, simply append the shortened URL ID to the base URL.

   For example:

   ```
   http://localhost:8080/1234567
   ```

### Dependencies

- [Fiber](https://github.com/gofiber/fiber): Fiber is a web framework for Go that's inspired by Express.js.
- [Limiter](https://github.com/gofiber/fiber/tree/v2/middleware/limiter): Rate limiter middleware for Fiber.
- [HTML Template](https://github.com/gofiber/template/tree/v2/html): HTML template engine for Fiber.

## Configuration

### Redirection Page
You can configure the redirection page by modifying the `views/redirection.html` file.
For disabling the redirection page, you can remove the 67th line from `urlshortener.go` file.

### URL Shortener Service
You can configure the URL shortener service by modifying app.env and database.env files.

## Todo
* Redis expiration time optimization

### Contributing

Contributions are welcome! Feel free to open an issue or submit a pull request on GitHub.


### License

This project is licensed under the MIT License. See the [LICENSE](https://github.com/umitanilkilic/advanced-url-shortener/blob/main/LICENSE) file for details.
