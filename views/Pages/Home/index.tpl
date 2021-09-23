<script>
    $(document).ready(function () {
        $("#tab-home").addClass("active");
        $("#search").click(function () {
        ntable.ajax.reload();
        });

        var today = new Date();

        var date = today.getFullYear()+'-'+(today.getMonth()+1)+'-'+today.getDate();
        document.getElementById("tanggal").innerHTML = date

        $.ajax({
            type: "GET",
            url: '/home/json',
            success: function(response){
                jumlah_saldo = response.data[0].jumlah_saldo
                jumlah_saldo = (jumlah_saldo).toLocaleString('id-ID', {
                    style: 'currency',
                    currency: 'IDR',
                });
                document.getElementById("saldo").innerHTML = jumlah_saldo
            }
        });
    })
</script>