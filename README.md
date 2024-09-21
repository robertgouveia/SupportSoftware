
# Support Software

Allows for technical support specialists to execute SQL extracts seamlessly


## Demo

GIF To be added.


## Documentation

#### .env

```bash
SERVER=10.10.10
USERNAME=USERNAME
PASSWORD=PASSWORD
PORT=1433
DATABASE=master
```

Server will be the address without the final number (connection)
Username and Password for the SQL Query can be inserted here.
Port by default is 1433
Database by default is master

#### connection.txt

```bash
127
180
181
```

You will need to enter the final part of the connection IP and you can list as many as you want.

#### imports/example.sql

```sql
DECLARE @ID INT = 0;

SELECT * FROM Users WHERE ID = @ID
```

You must declare you variables for the software to pick them up.
