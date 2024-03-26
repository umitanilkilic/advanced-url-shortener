![GitHub License](https://img.shields.io/github/license/umitanilkilic/advanced-url-shortener)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/umitanilkilic/advanced-url-shortener)
[![Go Report Card](https://goreportcard.com/badge/github.com/umitanilkilic/advanced-url-shortener)](https://goreportcard.com/report/github.com/umitanilkilic/advanced-url-shortener)

# URL SHORTENER - SPECS

- A sign up page
- A sign in page
- Create a new short URL page
- A short URL should have a URL that it points to 
   - A name
   - An ID 
   - A flag that keeps track whether the short URL is active and accessible - if this is set to false by the user it the short url should give 404
- List URLS page
   - Users should be able to delete individual lists here 
   - Users should be able to bulk delete lists here
   - Users should be able to order by creation date, update date, and clicks
- A detail page 
   - Users should be able to delete and edit the URL - AKA where it points to, if it is active, change the name etc. but the id should be read only
   - It should contain stats on the URL
      - How many times it has been clicked
      - IP addresses of the clicks 
      - Device info of the clicks 
      - Mobile vs Desktop

# Todo
- [x] JWT Authentication
- [ ] Dockerize the application
