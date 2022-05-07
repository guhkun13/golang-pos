# Project Structure

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
  - router
    - router.go
  - <your-app-here>
    - <app-name>_controller.go
    - <app-name>_model.go
    - <app-name>_route.go
    - <app-name>_util.go


