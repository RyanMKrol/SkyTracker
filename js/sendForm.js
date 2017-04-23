$( "form" ).on( "submit", function( event ) {

  if($("input[name='SourceAirport']:checked").length == 0){
    alert("You must select at least one source airport!");
    return false;
  }

  if($("input[name='Month']:checked").length == 0){
    alert("You must select at least one month to track!");
    return false;
  }

  console.log($( "#budgetSlider" ).slider( "value" ))
  console.log($( "#tripLengthSlider" ).slider( "values", 0 ))
  console.log($( "#tripLengthSlider" ).slider( "values", 1 ))


  // now we've got an array for what's checked and what isn't
  var monthArray = {};
  $("input[name='Month']").each(function(){
    monthArray[$(this).val()] = $(this).is(":checked");
  });

  // now we've got an array for what's checked and what isn't
  var airportArray = {};
  $("input[name='SourceAirport']").each(function(){
    airportArray[$(this).val()] = $(this).is(":checked");
  });

  var thing = JSON.stringify(monthArray);
  var other = JSON.stringify(airportArray);

  return false;

});
