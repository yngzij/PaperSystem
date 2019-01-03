$().ready(function () {
    post_logout=function(url) {
        $.ajax({
            type: 'POST',
            url: url,
            success: function () {
                location.reload();
            },
            error:function () {

            }
        });

    }
})
