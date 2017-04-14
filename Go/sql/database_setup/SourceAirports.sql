DROP TABLE SourceAirports;

CREATE TABLE SourceAirports (
  SrcAirportID int NOT NULL AUTO_INCREMENT,
  SrcAirportName varchar(255) NOT NULL,
  SrcAirportCode varchar(255) NOT NULL UNIQUE,
  SrcCountry varchar(255) NOT NULL,
  SrcCity varchar(255) NOT NULL,
  PRIMARY KEY(SrcAirportID)
);

INSERT INTO SourceAirports (SrcAirportName, SrcAirportCode, SrcCountry, SrcCity) VALUES ('London Heathrow','LHR','United Kingdom','London Heathrow');
INSERT INTO SourceAirports (SrcAirportName, SrcAirportCode, SrcCountry, SrcCity) VALUES ('London Gatwick','LGW','United Kingdom','London Gatwick');
INSERT INTO SourceAirports (SrcAirportName, SrcAirportCode, SrcCountry, SrcCity) VALUES ('London Luton','LTN','United Kingdom','London Luton');
INSERT INTO SourceAirports (SrcAirportName, SrcAirportCode, SrcCountry, SrcCity) VALUES ('Birmingham International','BHX','United Kingdom','Birmingham');
INSERT INTO SourceAirports (SrcAirportName, SrcAirportCode, SrcCountry, SrcCity) VALUES ('Manchester','MAN','United Kingdom','Manchester');
INSERT INTO SourceAirports (SrcAirportName, SrcAirportCode, SrcCountry, SrcCity) VALUES ('Edinburgh','EDI','United Kingdom','Edinburgh');
INSERT INTO SourceAirports (SrcAirportName, SrcAirportCode, SrcCountry, SrcCity) VALUES ('Glasgow','GLA','United Kingdom','Glasgow');
INSERT INTO SourceAirports (SrcAirportName, SrcAirportCode, SrcCountry, SrcCity) VALUES ('East Midlands','EMA','United Kingdom','East Midlands');
INSERT INTO SourceAirports (SrcAirportName, SrcAirportCode, SrcCountry, SrcCity) VALUES ('Bristol','BRS','United Kingdom','Bristol');
INSERT INTO SourceAirports (SrcAirportName, SrcAirportCode, SrcCountry, SrcCity) VALUES ('Newcastle','NCL','United Kingdom','Newcastle');

SELECT * FROM SourceAirports;
