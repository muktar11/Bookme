CREATE TABLE IF NOT EXISTS mail(
    ID  serial PRIMARY KEY,
    Too VARCHAR (2250) NOT NULL,
    Froom VARCHAR(2250) NOT NULL, 
    Subjeect VARCHAR (2250) NOT NULL,
    Content VARCHAR(2250) NOT NULL,
    Template VARCHAR (2250) NOT NULL, 
    created_on timestamp default CURRENT_TIMESTAMP not null,
    updated_on timestamp default CURRENT_TIMESTAMP
);

