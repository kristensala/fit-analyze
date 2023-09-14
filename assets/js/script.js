async function renderView() {
    const response = await fetch("http://localhost:5432/api/fit/test-data")
    const data = await response.json()

    console.log(data)
    const records = data.records

    renderChart(records)
    renderMap(records)
}

function renderChart(records) {
    const width = 1600;
    const height = 500;
    const marginTop = 20;
    const marginRight = 30;
    const marginBottom = 30;
    const marginLeft = 40;

    // Declare the x (horizontal position) scale.
    const x = d3.scaleUtc(d3.extent(records, d => new Date(d.timestamp)), [marginLeft, width - marginRight]);

    // Declare the y (vertical position) scale.
    const y = d3.scaleLinear([0, d3.max(records, d => d.heartRate)], [height - marginBottom, marginTop]);
    const y2 = d3.scaleLinear([0, d3.max(records, d => d.power)], [height - marginBottom, marginTop]);

    // Declare the line generator.
    const line = d3.line().x(d => x(new Date(d.timestamp))).y(d => y(d.heartRate));
    const line2 = d3.line().x(d => x(new Date(d.timestamp))).y(d => y(d.power));

    // Create the SVG container.
    const svg = d3.create("svg")
        .attr("width", width)
        .attr("height", height)
        .attr("viewBox", [0, 0, width, height])
        .attr("style", "max-width: 100%; height: auto; height: intrinsic;");

    // Add the x-axis.
    svg.append("g")
        .attr("transform", `translate(0,${height - marginBottom})`)
        .call(d3.axisBottom(x).ticks(width / 80).tickSizeOuter(0));

    // Add the y-axis, remove the domain line, add grid lines and a label.
    svg.append("g")
        .attr("transform", `translate(${marginLeft},0)`)
        .call(d3.axisLeft(y).ticks(height / 40))
        .call(d3.axisLeft(y2).ticks(height / 40))
        .call(g => g.select(".domain").remove())
        .call(g => g.selectAll(".tick line").clone()
            .attr("x2", width - marginLeft - marginRight)
            .attr("stroke-opacity", 0.1))
        .call(g => g.append("text")
            .attr("x", -marginLeft)
            .attr("y", 10)
            .attr("fill", "currentColor")
            .attr("text-anchor", "start")
            .text("Workout data"));

    // Append a path for the line.
    svg.append("path")
        .attr("fill", "none")
        .attr("stroke", "steelblue")
        .attr("stroke-width", 1.5)
        .attr("d", line(records))

    svg.append("path")
        .attr("fill", "none")
        .attr("stroke", "green")
        .attr("stroke-width", 1)
        .attr("d", line2(records))

    document.getElementById("container").append(svg.node());
}

function renderMap(records) {
    const coords = records.filter(function(item) {
        if (item.latitude !== 0 || item.longitude !== 0) {
            return true;
        }
        return false;
    }).map(function(item) {
        return [item.latitude, item.longitude]
    });

    console.log(coords);

    var map = L.map('map');
    L.tileLayer('https://tile.openstreetmap.org/{z}/{x}/{y}.png', {
        maxZoom: 19,
        attribution: '&copy; <a href="http://www.openstreetmap.org/copyright">OpenStreetMap</a>'
    }).addTo(map);

    var polyline = L.polyline(coords, {color: "blue"}).addTo(map);
    map.fitBounds(polyline.getBounds());
}

function renderSummary(summary) {

}
