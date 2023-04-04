   CREATE TABLE features (
    id BIGSERIAL PRIMARY KEY,
    feature_name VARCHAR(128) NOT NULL,
    slug VARCHAR(16) NOT NULL,
    description TEXT NOT NULL,
    module_name VARCHAR(32) NOT NULL,
    status BIT NOT NULL,
    created_on TIMESTAMP DEFAULT NOW()  ,
    updated_on TIMESTAMP DEFAULT NOW() 
);