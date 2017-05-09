DROP TABLE UserTravelMonths;

CREATE TABLE UserTravelMonths (
  UserTravelmonthsID int NOT NULL AUTO_INCREMENT,
  UserEmailAddress varchar(255) NOT NULL,
  TravelMonth int NOT NULL,
  PRIMARY KEY(UserTravelmonthsID),
  FOREIGN KEY (UserEmailAddress) REFERENCES Users(UserEmailAddress) ON DELETE CASCADE,
  UNIQUE KEY (UserEmailAddress, TravelMonth)
);
