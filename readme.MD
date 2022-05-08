# Project Structure 
- you can create many apps, but the structure of app folder is like this.   
- Inspired from Django app. Folder views is the static files (html, js, css)

## BACKEND

- .env
- main.go
- pkg
  - middleware
    - fiber_middleware.go
    - jwt_middleware.go
    - ...etc
  - utils
    - connection_url_builder.go
    - start_server.go
    - ...etc
  - database
    - <db_name>_connection.go
    - ...etc
  - config
    - fiber_config.go
    - ...etc

- apps
  - routes
    - routes.go
  - <app-name>  
    - controller:
      - <app-name>_controller.go
    - model:
      - <app-name>_model.go
    - route:
      - <app-name>_route.go   

## FULLSTACK (BE + FE)

- .env
- main.go
- pkg
  - middleware
    - fiber_middleware.go
    - jwt_middleware.go
    - ...etc
  - utils
    - connection_url_builder.go
    - start_server.go
    - ...etc
  - database
    - <db_name>_connection.go
    - ...etc
  - config
    - fiber_config.go
    - ...etc

  - templates
    - layout   
    - themes  
  - assets
    - css
    - js
    - img
  - media
   - ...

  - apps
    - routes
      - routes.go

    - <app-name>  
      - controller:
        - <app-name>_controller.go
      - model:
        - <app-name>_model.go
      - route:
        - <app-name>_route.go   
      <!-- optional -->
      - views:
        - static:
          - index.html
          - add.html
          - edit.html
          - view.html
        - css:
          - ...
        - js:
          - ...


