$(document).ready(function() {
	$("#generate-password").on('click', function() {
		$.ajax({
			url: "/raw",
			method: "GET",
			success: function(data) {
				
				$("#password-field").html($("<div />").text(data).html());
			},
		});
	});

	$("#generate-password").click()
});
