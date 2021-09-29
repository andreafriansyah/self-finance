<script>
    $(document).ready(function () {
        $("#tab-grafik").addClass("active");

        var data = [
            { y: 'Januari', a: 4000000, b: 3800000},
            { y: 'Februari', a: 3800000,  b: 3100000},
            { y: 'Maret', a: 4000000,  b: 2000000},
            { y: 'April', a: 4500000,  b: 4300000},
            { y: 'Mei', a: 4300000,  b: 2000000},
            { y: 'Juni', a: 4100000,  b: 3800000},
            { y: 'Juli', a: 3800000, b: 2000000},
            { y: 'Agustus', a: 3750000, b: 2000000},
            { y: 'September', a: 3750000, b: 2000000},
            { y: 'Oktober', a: 3750000, b: 2000000},
            { y: 'November', a: 3750000, b: 2000000},
            { y: 'December', a: 3750000, b: 2000000}
        ],
        config = {
            data: data,
            xkey: 'y',
            ykeys: ['a', 'b'],
            labels: ['Total Income', 'Total Outcome'],
            fillOpacity: 0.6,
            hideHover: 'auto',
            behaveLikeLine: true,
            resize: true,
            pointFillColors:['#ffffff'],
            pointStrokeColors: ['black'],
            lineColors:['gray','red']
        };
        config.element = 'bar-chart';
        Morris.Bar(config);
    })
</script>