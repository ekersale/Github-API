
$(document).ready(function() {
    $(".loginButton").click(function() {
        console.log("coucou");
        $.ajax({
            url: "/login?id=" + $("#userToken").val(),
            method: "POST",
            success: (function(rep) {
                console.log(rep);
                $(".loader").css("visibility", "visible");
                window.location.href = window.location.href + "index";
            }),
            error : (function(e) {
                 BootstrapDialog.show({
                message: 'Bad token, please verify the token validity'
            });
                console.log(e.message);
            })
        });
    });
    
    $(".SkipIntro").click(function() {
        window.location.href = window.location.href + "index";
    });
});