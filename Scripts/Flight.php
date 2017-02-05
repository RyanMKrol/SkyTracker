<?php

class Flight {

    //properties
    private $srcAirport;
    private $destAirport;
    private $price;
    private $srcCity;
    private $srcCountry;
    private $destCity;
    private $destCountry;

    //constructor
    function __construct($srcAirport, $destAirport, $price, $srcCity, $srcCountry, $destCity, $destCountry) {

      $this->srcAirport = $srcAirport;
      $this->destAirport = $destAirport;
      $this->price = $price;
      $this->srcCity = $srcCity;
      $this->srcCountry = $srcCountry;
      $this->destCity = $destCity;
      $this->destCountry = $destCountry;

    }

    //my getter function
    public function __get($property) {

      if (property_exists($this, $property)) {
        return $this->$property;
      }

    }
}

?>
