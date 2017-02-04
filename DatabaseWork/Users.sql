DROP TABLE Users;

CREATE TABLE Users (
  UserID int NOT NULL AUTO_INCREMENT,
  UserEmailAddress varchar(255) NOT NULL UNIQUE,
  UserBudget INT NOT NULL,
  UserTripLength INT NOT NULL,
  PRIMARY KEY(UserID)
);
