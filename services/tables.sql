-- Drop existing tables if they exist
DROP TABLE IF EXISTS relationships;
DROP TABLE IF EXISTS relations;
DROP TABLE IF EXISTS persons;

-- Create the 'persons' table
CREATE TABLE persons (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    gender ENUM('Male', 'Female') NOT NULL
);

-- Create the 'relations' table
CREATE TABLE relations (
    id INT AUTO_INCREMENT PRIMARY  KEY,
    name VARCHAR(255) UNIQUE NOT NULL
);

-- Create the 'relationships' table
CREATE TABLE relationships (
    id INT AUTO_INCREMENT PRIMARY KEY,
    person_id1 INT NOT NULL,
    person_id2 INT NOT NULL,
    relation_id INT NOT NULL,
    FOREIGN KEY (person_id1) REFERENCES persons(id),
    FOREIGN KEY (person_id2) REFERENCES persons(id),
    FOREIGN KEY (relation_id) REFERENCES relations(id)
);
