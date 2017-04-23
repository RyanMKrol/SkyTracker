$( "form" ).on( "submit", function( event ) {

  if($("input[name='SourceAirport']:checked").length == 0){
    alert("You must select at least one source airport!");
    return false;
  }

  if($("input[name='Month']:checked").length == 0){
    alert("You must select at least one month to track!");
    return false;
  }

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

  var jsonRaw = {};

  jsonRaw["emailAddress"] = $("#EmailAddress").val()
  jsonRaw["budget"] = $( "#budgetSlider" ).slider( "value" );
  jsonRaw["tripMinLen"] = $( "#tripLengthSlider" ).slider( "values", 0 );
  jsonRaw["tripMaxLen"] = $( "#tripLengthSlider" ).slider( "values", 1 );
  jsonRaw["months"] = monthArray;
  jsonRaw["airports"] = airportArray;

  var jsonData = JSON.stringify(jsonRaw);

  $.ajax({
    type: "POST",
    url: "./../Scripts/addUser.php",
    data: {_data: jsonData},
    success: function(data){
      alert("things were done right");
      console.log(data);
    }
  });


  return false;

});
