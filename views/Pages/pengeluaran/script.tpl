<script>
    $(document).ready(function () {
        $("#tab-report").addClass("menu-open");
        $("#tab-outcome").addClass("active");
        $("#search").click(function () {
            let dari_tanggal = $("#dari_tanggal").val();
            let sampe_tanggal = $("#sampe_tanggal").val();
            ntable.ajax.reload();
        });
        var ntable = $("#table-outcome").DataTable({
        oLanguage: {
            sSearch: "Filter Data",
        },
        processing: true,
        serverSide: true,
        ajax: {
            url: "/outcome/json",
            data: function (d) {
                d.dari_tanggal = $("#dari_tanggal").val();
                d.sampe_tanggal = $("#sampe_tanggal").val();
            },
        },
        bSort: true,
        scrollX: true,
        columns: [
            {
                title: "Type",
                data: "type",
                name: "type",
            },
            {
                title: "Jumlah",
                data: "jumlah",
                name: "jumlah",
            },
            {
                title: "Tujuan",
                data: "asaltujuan",
                name: "asaltujuan",
            },
            {
                title: "Tanggal",
                data: "tanggal",
                name: "tanggal",
                width: "10%",
            },
            {
                title: "Keterangan",
                data: "keterangan",
                name: "keterangan",
                width: "10%",
            },
        ],
    });
    })
</script>
<script>
  $( function() {
    $( "#dari_tanggal" ).datepicker({
        dateFormat: 'yy-mm-dd',
    });
  });
  $( function() {
    $( "#sampe_tanggal" ).datepicker({
        dateFormat: 'yy-mm-dd',
    });
  });
</script>