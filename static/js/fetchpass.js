const history = [];

$(document).ready(function() {
	
	$("#generate-password").on('click', function() {
		$.ajax({
			url: "/raw",
			method: "GET",
			success: function(data) {
				pass = $("<div />").text(data).html();
				// Send to heading
				$("#password-field").html(pass);

				logPass(pass);
			},
		});
	});

	$("#generate-password").click()
});

function logPass(pass) {
	$("#history").html(history.join('</br>'));

	history.unshift(pass);

	if (history.length > 10) {
		history.pop();
	}
}