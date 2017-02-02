<?php

//fucntion to create a .sql file for the table we want to store the data in
function createMySQLFile($srcAirport, $destAirport) {

  $myfile = fopen("${srcAirport}_${destAirport}.sql", "w");

  fwrite($myfile, "DROP TABLE ${srcAirport}_${destAirport};\n\n");
  fwrite($myfile, "CREATE TABLE ${srcAirport}_${destAirport} (\n");
  fwrite($myfile, "\tTripID int NOT NULL AUTO_INCREMENT,\n");
  fwrite($myfile, "\tSourcePort varchar(255) NOT NULL,\n");
  fwrite($myfile, "\testPort varchar(255) NOT NULL,\n");
  fwrite($myfile, "\tDepartDate DATE NOT NULL UNIQUE,\n");
  fwrite($myfile, "\tReturnDate DATE NOT NULL UNIQUE,\n");
  fwrite($myfile, "\tPrice int NOT NULL,\n");
  fwrite($myfile, "\tPRIMARY KEY(TripID)\n");
  fwrite($myfile, ");\n");

}

?>
