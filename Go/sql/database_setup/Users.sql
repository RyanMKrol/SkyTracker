DROP TABLE Users;

CREATE TABLE Users (
  UserID int NOT NULL AUTO_INCREMENT,
  UserEmailAddress varchar(255) NOT NULL UNIQUE,
  UserBudget int DEFAULT NULL,
  UserTripMin int DEFAULT NULL,
  UserTripMax int DEFAULT NULL,
  PRIMARY KEY(UserID),
  CHECK (UserTripMin > 0),
  CHECK (UserTripMin < 29),
  CHECK (UserBudget > 0),
);
