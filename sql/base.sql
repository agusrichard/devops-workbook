-- Create users Table
CREATE TABLE users (
    _id SERIAL PRIMARY KEY,
    username VARCHAR(256) UNIQUE NOT NULL,
    password VARCHAR(256) NOT NULL
);

-- Create services table
CREATE TABLE services (
	_id SERIAL PRIMARY KEY,
	request_id INT UNIQUE NOT NULL,
	status VARCHAR(64) NOT NULL,
	vessel_name VARCHAR(256) NOT NULL,
	service_type VARCHAR(256) NOT NULL,
	data_agent VARCHAR(256) NOT NULL,
	cargo VARCHAR(256) NOT NULL,
	etd VARCHAR(256) NOT NULL,
	eta VARCHAR(256) NOT NULL,
    user_id INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(_id) ON DELETE CASCADE
);