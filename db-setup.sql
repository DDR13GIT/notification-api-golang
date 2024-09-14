CREATE DATABASE notification_db;

CREATE USER 'notifuser'@'localhost' IDENTIFIED BY 'admin123';
GRANT ALL PRIVILEGES ON notification_db.* TO 'notifuser'@'localhost';
FLUSH PRIVILEGES;

USE notification_db;
DROP TABLE IF EXISTS notifications;
CREATE TABLE notifications (
    id INT AUTO_INCREMENT PRIMARY KEY,
    message VARCHAR(255) NOT NULL,
    user_id INT NOT NULL,
    is_read BOOLEAN NOT NULL
);