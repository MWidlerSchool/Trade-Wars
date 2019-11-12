// init
setActiveForm("name-form");

// hide inactive forms, show the active one
function setActiveForm(targetString)
{
	// hide everything
	$(".divPanel").hide();
	// show the thing with the matching ID
	$("#" + targetString).show();
}

// read in player name from name page
function readInName()
{
	var playerName = document.getElementById("name-input").value;
	// test line; replace with actual code later
	console.log(playerName);
}