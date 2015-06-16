define(function (global, modules) {

	var bar = document.querySelector("#headstbar");
	var showClass = "show";

	function display(isShow) {
		if (isShow) {
			bar.classList.add(showClass);
		} else {
			bar.classList.remove(showClass);
		}
	}

	return {
		display: display
	}

});