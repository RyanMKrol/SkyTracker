<?php
  //globals and includes

  include 'credentials.php';
  include 'Flight.php';

  $currentDate = date_create();

  // I don't currently need any of these variables
  // $redBound = 0.2;
  // $yellowBound = 0.35;
  // $greenBound = 0.5;
  // $hardCap = 150;
  // $minTripLength = 2;
  // $limitSize = 10;

?>
<?php
  //functions

  //this function will house the algorithm to get the flights that I'm interested in
  function getFlightsOfInterest($conn,$src,$dest){

    //get the trip length from the query
    // maybe get the cost per day from the query as well
    // $query = "SELECT * , (Price/(SELECT AveragePrice FROM Averages WHERE AirPort = '$dest')) FROM ${src}_${dest} WHERE Price < ($bound * (SELECT AveragePrice FROM Averages WHERE AirPort = '$dest')) AND DATEDIFF(ReturnDate,DepartDate) > $minTripLength AND Price < $hardCap AND Price > $previousPriceCap ORDER BY (Price/(SELECT AveragePrice FROM Averages WHERE AirPort = '$dest')) ASC limit $limitSize;";

    //this query gets the information of the flight that has the lowest cost/day
    // $query = "SELECT *,DATEDIFF(ReturnDate,DepartDate) FROM BHX_MAD WHERE (Price/DATEDIFF(ReturnDate,DepartDate)) = (SELECT Min(Price/(DATEDIFF(ReturnDate,DepartDate))) from BHX_MAD WHERE DATEDIFF(ReturnDate,DepartDate) < 14) AND DATEDIFF(ReturnDate,DepartDate) <= 14;"

    //this gets the flights that have the minimum price in the data set
    // $query = "SELECT *, DATEDIFF(ReturnDate, DepartDate) FROM BHX_MAD Where Price = (Select Min(Price) from BHX_MAD);"

    //for now i'm going to go with the cheapest flight because it will always return the same numner of flights

    $query = "SELECT *, DATEDIFF(ReturnDate, DepartDate) FROM ${src}_${dest} Where Price = (Select Min(Price) from ${src}_${dest}) limit 1;";

    $pricesArray = arraySetup($conn,$query);

    //checking that some value was returned by the database
    if (!isset($pricesArray[0])) {
      $pricesArray[0] = NULL;
    }

    //[0] because I only want to return 1 result right now
    return ($pricesArray[0]);

  }

  //this function will be used to setup my source and destination airport arrays
  function arraySetup($conn, $query){

    $result = $conn->query($query);

    $resultArray = array();

    while($row = mysqli_fetch_array($result)){
      $resultArray[] = $row;
    }

    return $resultArray;
  }

  //comparator for usort
  function cmp($a, $b){
    if ($a['Price'] == $b['Price']) {
        return 0;
    }
    return ($a['Price'] < $b['Price']) ? -1 : 1;
  }

?>
<?php
  //main

  //date time object used in reports
  global $currentDate;

  // Create connection
  $conn = new mysqli($servername, $username, $password, $database);

  // Check connection
  if ($conn->connect_error) {
      die("Connection failed: " . $conn->connect_error);
  }

  $sourcesArr = arraySetup($conn, "SELECT * FROM SourceAirports;");
  $destinationsArr = arraySetup($conn, "SELECT * FROM DestinationAirports;");

  $flightsArray = array();
  $minItem = NULL;

  //getting a minimum for each destination, from all of the sources
  foreach($destinationsArr as $destAirport) {
    foreach($sourcesArr as $srcAirport) {

      $src = $srcAirport["SrcAirportCode"];
      $dest = $destAirport["DestAirportCode"];

      // print_r($srcAirport);

      $potentialMin = getFlightsOfInterest($conn,$src,$dest);

      //checks that the minItem and potentialMin variables aren't null
      if(isset($minItem) && isset($potentialMin)){
        if($potentialMin['Price'] < $minItem['Price']){
          $minItem = $potentialMin;
          $minItem["SrcCity"] = $srcAirport["SrcCity"];
          $minItem["SrcCountry"] = $srcAirport["SrcCountry"];
          $minItem["DestCity"] = $destAirport["DestCity"];
          $minItem["DestCountry"] = $destAirport["DestCountry"];
        }
      } else {
        if(isset($potentialMin)){
          $minItem = $potentialMin;
          $minItem["SrcCity"] = $srcAirport["SrcCity"];
          $minItem["SrcCountry"] = $srcAirport["SrcCountry"];
          $minItem["DestCity"] = $destAirport["DestCity"];
          $minItem["DestCountry"] = $destAirport["DestCountry"];
        }
      }
    }

    //add the minimimum item to the array of minimums
    array_push($flightsArray, $minItem);
    $minItem = NULL;
  }


  //the report file
  $myfile = fopen(("/var/www/html/skytracker.co/Reports/Report_" . date_format($currentDate,"d-m-Y") . ".txt"), "w+");

  fwrite($myfile, "Welcome to your report! I've organised the flights by prices for you to make it all a bit easier.\n");

  usort($flightsArray, "cmp");

  $under50Flag = false;
  $under100Flag = false;
  $over100Flag = false;

  foreach($flightsArray as $flight){

    if($flight['Price'] < 50){
      if(!$under50Flag){
        $under50Flag = true;
        fwrite($myfile, "\nFlights Under £50:\n");
      }
    } else if($flight['Price'] < 100){
      if(!$under100Flag){
        $under100Flag = true;
        fwrite($myfile, "\nFlights Under £100:\n");
      }
    } else {
      if(!$over100Flag){
        $over100Flag = true;
        fwrite($myfile, "\nFlights Over £100:\n");
      }
    }

    fwrite($myfile, "${flight['SrcCity']}, ${flight['SourcePort']}\t\t=>\t${flight['DestCountry']}, ${flight['DestCity']}, ${flight['DestPort']}:\t\tLeaving - ${flight['DepartDate']}. Returning - ${flight['ReturnDate']}. Trip Length: ${flight['DATEDIFF(ReturnDate, DepartDate)']} Days. Cost: £${flight['Price']}\n");

  }

  //closing the file
  fclose($myfile);

  //close connection
  $conn->close();

?>
