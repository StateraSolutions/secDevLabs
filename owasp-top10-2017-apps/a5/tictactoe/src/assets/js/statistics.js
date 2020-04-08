window.onload = function() {
    fetch('http://localhost.:10005/statistics/data')
        .then(resp => resp.json())
        .then(data => {
            renderChart(data)
        })

}

function renderChart(data){
    const chart = new CanvasJS.Chart("chartContainer", {
        animationEnabled: true,
        title: {
            text: "Your statistics"
        },
        data: [{
            type: "pie",
            startAngle: 240,
            yValueFormatString: "##0.00\"%\"",
            indexLabel: "{label} {y}",
            dataPoints: data.chartData
        }]
    });
    chart.render();
    document.getElementById('games').innerHTML = data.numbers.games
    document.getElementById('wins').innerHTML = data.numbers.wins
    document.getElementById('ties').innerHTML = data.numbers.ties
    document.getElementById('loses').innerHTML = data.numbers.loses
}