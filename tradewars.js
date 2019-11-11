// init
setActiveForm("name-form");

// hide inactive forms, show the active one, gray out the unselected buttons and blue the selected one
function setActiveForm(targetString)
{
	// hide/grey everything
	$(".formPanel").hide();
	// show the thing with the matching ID
	$("#" + targetString).show();
}

