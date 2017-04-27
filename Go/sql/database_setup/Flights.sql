DROP TABLE Flights;

CREATE TABLE Flights (
  FlightID int NOT NULL AUTO_INCREMENT,
  SourceAirportCode varchar(255) NOT NULL,
  DestinationAirportCode varchar(255) NOT NULL,
  DepartDate DATE NOT NULL,
  ReturnDate DATE NOT NULL,
  Price int NOT NULL,
  PRIMARY KEY(FlightID),
  FOREIGN KEY (SourceAirportCode) REFERENCES SourceAirports(SrcAirportCode) ON DELETE CASCADE,
  FOREIGN KEY (DestinationAirportCode) REFERENCES DestinationAirports(DestAirportCode) ON DELETE CASCADE
);
