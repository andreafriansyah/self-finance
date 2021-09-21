<script>
    $(document).ready(function(){
            console.log("masuk")
        $("form").validate({
            rules: {
                email: {
                    email: true,
                    required: true,
                },
                password: {
                    required: true,
                }
            },
            submitHandler: function (form) {
                console.log(form)
                if ($("form").hasClass("success")) {
                    return true;
                } else {
                    var data = {
                        email: $("#email").val(),
                        password: $("#password").val(),
                    }
                    console.log(data);
                    validate(data, '/auth/validate-login', form = "form")
                }
            },
        });
    })
</script>