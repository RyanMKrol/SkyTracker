DROP TABLE UserSourceAirports;

CREATE TABLE UserSourceAirports (
  UserSourceID int NOT NULL AUTO_INCREMENT,
  UserEmailAddress varchar(255) NOT NULL,
  SourceAirportCode varchar(255) NOT NULL,
  PRIMARY KEY(UserSourceID),
  FOREIGN KEY (UserEmailAddress) REFERENCES Users(UserEmailAddress) ON DELETE CASCADE,
  FOREIGN KEY (SourceAirportCode) REFERENCES SourceAirports(SrcAirportCode) ON DELETE CASCADE
);
