CREATE TABLE IF NOT EXISTS users(
    Id  serial PRIMARY KEY,
    FirstName VARCHAR (50) NOT NULL,
    LastName VARCHAR(50) NOT NULL, 
    Email VARCHAR (255) UNIQUE NOT NULL, 
    Password VARCHAR (50) NOT NULL,
    AccessLevel INT, 
    CreatedAt timestamp default CURRENT_TIMESTAMP not null,
    UpdatedAt timestamp default CURRENT_TIMESTAMP
);

