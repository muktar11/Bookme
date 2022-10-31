CREATE TABLE IF NOT EXISTS RoomRestriction(
    ID  serial PRIMARY KEY,
    StartDate  DATE NOT NULL, 
    EndDate DATE NOT NULL,
    RoomId INT,
    ReservationID INT,
    RestrictionID INT, 
    created_on timestamp default CURRENT_TIMESTAMP not null,
    updated_on timestamp default CURRENT_TIMESTAMP
);

