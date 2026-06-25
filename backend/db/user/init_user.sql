CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    hashed_password VARCHAR(255) NOT NULL,
    role ENUM('customer', 'admin') NOT NULL,
    created_at DATETIME NOT NULL
);

ALTER TABLE users ADD CONSTRAINT uq_users_email UNIQUE (email);
