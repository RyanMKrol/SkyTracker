DROP TABLE UserSourceAirports;

CREATE TABLE UserSourceAirports (
  UserSourceID int NOT NULL AUTO_INCREMENT,
  UserID int NOT NULL,
  SourceAirportCode varchar(255) NOT NULL,
  PRIMARY KEY(UserSourceID),
  FOREIGN KEY (UserID) REFERENCES Users(UserID) ON DELETE CASCADE,
  FOREIGN KEY (SourceAirportCode) REFERENCES SourceAirports(SrcAirportCode) ON DELETE CASCADE
);
