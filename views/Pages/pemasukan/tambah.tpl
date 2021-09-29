<script>
    $(document).ready(function () {
        $("#tab-report").addClass("menu-open");
        $("#tab-income").addClass("active");
    })

    $( function() {
        $( "#tanggal" ).datepicker({
            dateFormat: 'yy-mm-dd',
        });
    });
</script>