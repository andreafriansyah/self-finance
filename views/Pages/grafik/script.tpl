<script>
    $(document).ready(function () {
        $("#tab-grafik").addClass("active");
        var today = new Date();
        var month = today.getMonth()+1
        document.getElementById("month").innerHTML = "Bulan : " + month

        $.ajax({
            type: "GET",
            url: '/grafik/json',
            success: function(response){
                console.log(response)
            }
        });

        Morris.Donut({
            element: 'pie-chart',
            data: [
                {label: "Income", value: 30},
                {label: "Outcome", value: 15},
            ]
        })
    })
</script>