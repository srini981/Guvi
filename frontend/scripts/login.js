$(document).ready(function() {
	$('#login-form').submit(function(event) {
		event.preventDefault();
		var email = $('#email').val();
		var password = $('#password').val();
		if (email === '' || password === '') {
			$('#error-message').html('Please enter both username and password.');
		} else {
			// Perform AJAX request to backend to authenticate user
			$.ajax({
				url: 'http://localhost:8080/login',
				type: 'POST',
				headers: {
					"content-type": "application/json"
				},
				data: JSON.stringify({
					email: email,
					password: password
				}),
				success: function(response) {
                    //TODO: get session id and store here
					console.log(response)
                    localStorage.session_id= response;
                    var url = "/profile.html";    
                    $(location).attr('href',url);
				},
				error: function(xhr, status, error) {
					$('#error-message').html(xhr.responseText);
				}
			});
		}
	});
});
