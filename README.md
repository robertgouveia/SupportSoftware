
# Support Software

Allows for technical support specialists to execute SQL extracts seamlessly


## Demo

Server Selection <br/>
![Screenshot 2024-09-21 at 21 48 20](https://github.com/user-attachments/assets/334066b7-d600-483e-8e0a-2adf75238591)

Database Selection <br/>
![Screenshot 2024-09-21 at 21 49 36](https://github.com/user-attachments/assets/0cd54c55-e9e0-4bd4-b33b-99d8f37c4539)

Export Selection <br/>
![Screenshot 2024-09-21 at 21 50 01](https://github.com/user-attachments/assets/d3bd23b0-b4a1-4cec-954c-34160d0161c1)

User Input <br/>
![Screenshot 2024-09-21 at 21 50 37](https://github.com/user-attachments/assets/6ff485ba-23fb-41b9-9e52-de5a9c03a01d)

CSV Results <br/>
![Screenshot 2024-09-21 at 21 51 16](https://github.com/user-attachments/assets/717aea7d-0384-49c6-8932-3871a609740a)

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
