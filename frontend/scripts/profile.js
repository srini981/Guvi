$(document).ready(function() { 
    if (!localStorage["session_id"]) {
        var url = "/login.html";    
        $(location).attr('href',url);
    }
   var sessionid=localStorage["session_id"].replace("\"","")
    $.ajax({
        url: 'http://localhost:8080/profile',
        type: 'GET',
        headers: {
            "content-type": "application/json",
            "session-id":sessionid,
        },
        success: function(response) {
            $("#name").html(response.Name);
            $("#email").html(response.Email);
            $("#age").html(response.Age);
        },
        error: function(xhr, status, error) {
            alert(xhr.responseText)

            $('#error-message').html(xhr.responseText);
            var url = "/login.html";    
            $(location).attr('href',url);

        }
    });

    $(".logout-btn").click(function(){
        alert('Logged out Successfully')
        localStorage.removeItem("session_id");
        var url = "/login.html";    
        $(location).attr('href',url);
      });
});

