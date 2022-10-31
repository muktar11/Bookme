CREATE TABLE IF NOT EXISTS Restriction(
    ID  serial PRIMARY KEY,
    RestrictionName VARCHAR (50) NOT NULL, 
    created_on timestamp default CURRENT_TIMESTAMP not null,
    updated_on timestamp default CURRENT_TIMESTAMP
);

