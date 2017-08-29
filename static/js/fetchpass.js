$(document).ready(function() {
	$("#generate-password").on('click', function() {
		$.ajax({
			url: "http://localhost:3000/raw",
			method: "GET",
			success: function(data) {
				$("#password-field").html(data);
			},
		});
	});

	$("#generate-password").click()
});