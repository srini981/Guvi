$(document).ready(function() {
	$('#signup-form').submit(function(event) {
		event.preventDefault();
		var username = $('#username').val();
		var age = $('#age').val();
		var email = $('#email').val();
		var password = $('#password').val();
		var confirmPassword = $('#confirm-password').val();
		if (username === '' || email === '' || password === '' || confirmPassword === '') {
			$('#error-message').html('Please fill in all fields.');
		} else if (password !== confirmPassword) {
			$('#error-message').html('Passwords do not match.');
		} else {
			// Perform AJAX request to backend to create new user account

			var data={
				name: username,
				email: email,
				password: password,
				age:Number(age)
			}
			
			$.ajax({
				url: 'http://localhost:8080/signup',
				type: 'POST',
				headers: {
					"content-type": "application/json"
				},
				data: JSON.stringify(data),
				success: function(response) {
					alert('Signup successful!');
                    var url = "/login.html";    
                    $(location).attr('href',url);
				},
				error: function(xhr, status, error) {
					$('#error-message').html(xhr.responseText);
				}
			});
		}
	});
});
