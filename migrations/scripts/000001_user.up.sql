CREATE TABLE roles (
    id BIGSERIAL PRIMARY KEY,
    role_name VARCHAR(128) NOT NULL,
    description TEXT NOT NULL,
    created_on TIMESTAMP DEFAULT NOW() ,
    updated_on TIMESTAMP DEFAULT NOW() 
);
