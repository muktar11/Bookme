CREATE TABLE IF NOT EXISTS reservations(
    ID  serial PRIMARY KEY,
    FirstName VARCHAR (50) NOT NULL,
    LastName VARCHAR(50) NOT NULL, 
    Email VARCHAR (255) UNIQUE NOT NULL,
    Phone VARCHAR(50) NOT NULL, 
    StartDate  DATE NOT NULL, 
    EndDate DATE NOT NULL,
    RoomID INT,
    Processed INT, 
    CreatedAt timestamp default CURRENT_TIMESTAMP not null,
    UpdatedAt timestamp default CURRENT_TIMESTAMP
);



ALTER TABLE "reservations" ADD FOREIGN KEY ("roomid") REFERENCES "room" ("id");
CREATE INDEX ON "reservations" ("id");




CREATE TABLE room(
   roomid INT GENERATED ALWAYS AS IDENTITY,
   roomname VARCHAR(255) NOT NULL,
   CreatedAt timestamp default CURRENT_TIMESTAMP not null,
   UpdatedAt timestamp default CURRENT_TIMESTAMP
   PRIMARY KEY(customer_id)
);

CREATE TABLE reservations(
   reservationsid INT GENERATED ALWAYS AS IDENTITY,
   firstname VARCHAR(255) NOT NULL,
   lastname VARCHAR(255) NOT NULL,
   email VARCHAR(255) NOT NULL,
   phone INT(255) NOT NULL, 
   StartDate  DATE NOT NULL, 
   EndDate DATE NOT NULL,
   processed INT,
   roomid INT, 
   PRIMARY KEY(id),
   CONSTRAINT fk_room
      FOREIGN KEY(roomid) 
	  REFERENCES room(roomid)
);


INSERT INTO contacts(firstname, lastname, email, phone, startdate, enddate, roomid
VALUES('Muktar ', 'Iberahim ', 'muktar@gmail.com, (408)-111-1234', '2022-10-28 ', '2022-10-29 ', 1 ;