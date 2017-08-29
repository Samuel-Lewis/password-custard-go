$(document).ready(function() {
	$("#generate-password").on('click', function() {
		$.ajax({
			url: "/raw",
			method: "GET",
			success: function(data) {
				$("#password-field").html(data);
			},
		});
	});

	$("#generate-password").click()
});