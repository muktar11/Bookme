CREATE TABLE IF NOT EXISTS Room(
    ID  serial PRIMARY KEY,
    RoomName VARCHAR (50) NOT NULL, 
    CreatedAt timestamp default CURRENT_TIMESTAMP not null,
    UpdatedAt timestamp default CURRENT_TIMESTAMP
);

