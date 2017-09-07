let last_raw = "";

// Pulls password from /raw API
$(document).ready(function() {
	// GET data from /raw to show
	$("#generate-password").on('click', function() {
		$.ajax({
			url: "/raw",
			method: "GET",
			success: function(data) {
				pass = $("<div />").text(data).html();
				// Send to heading
				$("#password-field").html(pass);

				last_raw = data;
				logPass(pass);
			},
		});
	});
	$("#generate-password").click();


	// Copy on click
	$("#copy-password").on('click', function() {
		// Create temp textarea, put text in it, copy from that
		var textArea = document.createElement("textarea");
		textArea.value = last_raw;
		document.body.appendChild(textArea);

		textArea.select();
		try {
			document.execCommand('copy');
		} catch (err) {
			alert("Browser does not support copy and paste automation :(");
		}
		document.body.removeChild(textArea);
	});
});

// Pastes last 10 passwords to history box
const history = [];
function logPass(pass) {
	$("#history").html(history.join('</br>'));

	history.unshift(pass);

	if (history.length > 10) {
		history.pop();
	}
}
