<script>
    $(document).ready(function () {
        $("#tab-report").addClass("menu-open");
        $("#tab-outcome").addClass("active");
        $("#search").click(function () {
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