DROP TABLE UserSourceAirports;
DROP TABLE UserTravelMonths;

DROP TABLE Users;

CREATE TABLE Users (
  UserID int NOT NULL AUTO_INCREMENT,
  UserEmailAddress varchar(255) NOT NULL UNIQUE,
  UserBudget int DEFAULT NULL,
  UserTripMin int DEFAULT NULL,
  UserTripMax int DEFAULT NULL,
  UserReportFrequency int DEFAULT 2,
  UserLastReport DATE,
  UserSalt varchar(255) NOT NULL UNIQUE,
  PRIMARY KEY(UserID),
  FOREIGN KEY (UserReportFrequency) REFERENCES FrequencyLookup(FLID) ON DELETE CASCADE
);

SOURCE UserSourceAirports.sql;
SOURCE UserTravelMonths.sql;
