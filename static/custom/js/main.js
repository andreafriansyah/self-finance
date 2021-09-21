// jquery validation
jQuery.validator.setDefaults({
    errorPlacement: function (error, element) {
        error.addClass("help-block");
        if (!$(element).hasClass("select-2")) {
            // Add the `help-block` class to the error element
            error.addClass("help-block invalid-feedback");
            // Add `has-error` class to the parent div.form-group
            // in order to add icons to inputs
            element.parent("div").addClass("has-error");
            if (element.prop("type") === "checkbox") {
                error.insertAfter(element.parent("label"));
            } else {
                error.insertAfter(element);
            }
        } else {
            // select 2
            element.parent("div").addClass("has-error");
            error.insertAfter(element.parent("div").children()[1]);
            $('<label class="fa fa-times form-control-feedback"></label>').insertAfter(element.parent("div").children()[1]);
            var select2 = $(element).next("span");
            $(select2).find(".selection").children("span").removeClass('select2-success').addClass('select2-error');
        }
    },
    // success: function (label, element) {
    //     // Add the span element, if doesn't exists, and apply the icon classes to it.
    //     if (!$(element).next("span")[0]) {
    //         $("<span class='glyphicon glyphicon-ok form-control-feedback'></span>").insertAfter($(element));
    //     }
    // },
    highlight: function (element, errorClass, validClass) {
        $(element).parent("div").addClass("has-error").removeClass("has-success");
        $(element).addClass("is-invalid").removeClass("is-valid");
        if ($(element).hasClass("select-2")) {
            var select2 = $(element).next("span");
            $(select2).find(".selection").children("span").removeClass('select2-success').addClass('select2-error');
        }
    },
    unhighlight: function (element, errorClass, validClass) {
        $(element).parent("div").addClass("has-success").removeClass("has-error");
        $(element).addClass("is-valid").removeClass("is-invalid");
        if (!$(element).hasClass("select-2")) {
            $(element).next("span").remove()
            $(element).next("label").remove()
        } else {
            $(element).parent("div").find("label").remove()
            var select2 = $(element).next("span");
            $(select2).find(".selection").children("span").removeClass('select2-error').addClass('select2-success');
        }
    }
});

$.validator.addMethod(
    "regex",
    function (value, element, regexp) {
        var re = new RegExp(regexp);
        return this.optional(element) || re.test(value);
    },
    "Please check your input."
);
// jquery validation

// ajax xsrf
var ajax = $.ajax;
$.extend({
    ajax: function (url, options) {
        if (typeof url === 'object') {
            options = url;
            url = undefined;
        }
        options = options || {};
        url = options.url;
        var xsrftoken = $('meta[name=_xsrf]').attr('content');
        var headers = options.headers || {};
        var domain = document.domain.replace(/\./ig, '\\.');
        if (!/^(http:|https:).*/.test(url) || eval('/^(http:|https:)\\/\\/(.+\\.)*' + domain + '.*/').test(url)) {
            headers = $.extend(headers, { 'X-Xsrftoken': xsrftoken });
        }
        options.headers = headers;
        return ajax(url, options);
    }
});
// ajax xsrf

//validate ajax
validate = (data, url, form = "form") => {
    $.ajax({
        data: data,
        url: url,
        type: "POST",
        success: function (e) {
            if (e.code !== 200) {
                setTimeout(function () {
                    $("#loadMe").modal('hide');
                    e.message.map(function (ex) {
                        iziToast.error({
                            title: 'Warning',
                            message: ex.message,
                            position: 'topRight'
                        });
                    })
                }, 1000);
            } else {
                setTimeout(function () {
                    $("#loadMe").modal('hide');
                    $(form).addClass("success");
                    $(form).submit();
                }, 1000);
            }
        },
        beforeSend: function () {
            $("#loadMe").modal('show');
        },
        error: function (e) {
            setTimeout(function () {
                $("#loadMe").modal('hide');
                iziToast.error({
                    title: 'Warning',
                    message: "Bad Request",
                    position: 'topRight'
                });
            }, 1000);
        }
    });
}

$("#btnLogout").click(function () {
    swal("Logout", "Are you sure want logout?", {
        buttons: ["Cancel", "Logout"],
    }).then((logout) => {
        if (logout) {
            window.location.href = "/auth/signout";
        }
    });
})

function formatNumber(num) {
    console.log(num);
    num = num.toString().replace(".", ",")
    return num.toString().replace(/(\d)(?=(\d{3})+(?!\d))/g, '$1.')
}
