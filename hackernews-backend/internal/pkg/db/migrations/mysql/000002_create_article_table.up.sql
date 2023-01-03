CREATE TABLE IF NOT EXISTS article(
    id INT NOT NULL UNIQUE AUTO_INCREMENT,
    title VARCHAR (255) ,
    url VARCHAR (255) ,
    user_id INT ,
    FOREIGN KEY (user_id) REFERENCES user(id) ,
    PRIMARY KEY (id)
)