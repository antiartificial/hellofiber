# hellofiber
Experiments with gofiber framework application design.

The objective is to write a testable service pattern and reason about the endpoints.
New seems like it's doing a lot, splitting into a setup which returns an application config
which is then passed into new to setup the routes seems like the way to go to to mock. Moving
the routes into functions creates testible handlers. Dependency injection cleans up stuff.

Todo:
Application resources (db/cache etc.) setup
