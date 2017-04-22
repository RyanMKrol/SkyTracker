$( "form" ).on( "submit", function( event ) {

  if($("input[name='SourceAirport']:checked").length == 0){
    alert("You must select at least one source airport!");
    return false;
  }

  if($("input[name='Month']:checked").length == 0){
    alert("You must select at least one month to track!");
    return false;
  }

  event.preventDefault();
  console.log( $( this ).serialize() );
  console.log($( "#budgetSlider" ).slider( "value" ))
  console.log($( "#tripLength" ).slider( "values", 0 ))
  console.log($( "#tripLength" ).slider( "values", 1 ))
});
