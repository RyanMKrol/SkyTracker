DROP TABLE Emails;

CREATE TABLE Emails (
  EmailID int NOT NULL AUTO_INCREMENT,
  EmailAddress varchar(255) NOT NULL UNIQUE,
  PRIMARY KEY(EmailID)
);
