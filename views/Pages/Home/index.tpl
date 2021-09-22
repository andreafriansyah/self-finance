<script>
    $(document).ready(function () {
        $("#tab-home").addClass("active");
        $("#search").click(function () {
        ntable.ajax.reload();
        });
    })
</script>