DROP TABLE UserTravelMonths;

CREATE TABLE UserTravelMonths (
  UserTravelmonthsID int NOT NULL AUTO_INCREMENT,
  UserID int NOT NULL,
  TravelMonth int NOT NULL,
  PRIMARY KEY(UserTravelmonthsID),
  FOREIGN KEY (UserID) REFERENCES Users(UserID) ON DELETE CASCADE,
  UNIQUE KEY (UserID, TravelMonth)
);
