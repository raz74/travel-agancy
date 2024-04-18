# Travel Agency

Travel Agency Route Management

This Go program simulates a travel agency's system for managing cities and roads. It allows users to:

* Add new cities
* Add new roads connecting cities
* Calculate the distance between two cities and how much it takes time

### Getting Started:

1. Clone this repository.
2. Run the program using `go run main.go`.
3. Follow the menu instructions to interact with the system.


### example input:
1
2
1
21
Tehran
1
251
Qom
1
361
Kashan
2
2
2
1
T-K
21
361
]21,251[
80
600
1
2
4
361:251
3
1
21
5


### and the output should be like this :

Main Menu - Select an action:
1. Help
2. Add
3. Delete
4. Path
5. Exit
   Select a number from shown menu and enter. For example 1 is for help.
   Main Menu - Select an action:
1. Help
2. Add
3. Delete
4. Path
5. Exit
   Select model:
1. City
2. Road
   id=?
   name=?
   City with id=21 added!
   Select your next action
1. Add another City
2. Main Menu
id=?
   name=?
   City with id=251 added!
   Select your next action
1. Add another City
2. Main Menu
   id=?
   name=?
   City with id=361 added!
   Select your next action
1. Add another City
2. Main Menu
   Main Menu - Select an action:
1. Help
2. Add
3. Delete
4. Path
5. Exit
   Select model:
1. City
2. Road
   id=?
   name=?
   from=?
   to=?
   through=?
   speed_limit=?
   length=?
   bi_directional=?
   Road with id=1 added!
   Select your next action
1. Add another Road
2. Main Menu
   Main Menu - Select an action:
1. Help
2. Add
3. Delete
4. Path
5. Exit
   Kashan:Qom via Road T-K: Takes 00:07:30
   Main Menu - Select an action:
1. Help
2. Add
3. Delete
4. Path
5. Exit
   Select model:
1. City
2. Road
   City:21 deleted!
   Main Menu - Select an action:
1. Help
2. Add
3. Delete
4. Path
5. Exit
