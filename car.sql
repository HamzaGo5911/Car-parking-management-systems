CREATE DATABASE IF NOT EXISTS car_parking;

USE car_parking;

CREATE TABLE car (
                     ID VARCHAR(36) NOT NULL,
                     LicensePlate VARCHAR(255) NOT NULL,
                     Brand VARCHAR(255) NOT NULL,
                     Model VARCHAR(255) NOT NULL,
                     Color VARCHAR(255) NOT NULL,
                     PlateNumber VARCHAR(255) NOT NULL,
                     PRIMARY KEY (ID)
);
