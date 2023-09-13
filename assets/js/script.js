async function decodeFitFile() {
    const response = await fetch("http://localhost:5432/api/fit/decode")
    const data = await response.json()
    console.log(data)

    if (data !== null) {

    }
}

function calculateSelectedAreaAvgPower() {

}

function calculateSelectedAreaAvgHeartRate() {

}

async function renderChart() {
    const response = await fetch("http://localhost:5432/api/fit/test-data")
    const aapl = await response.json()
    console.log(aapl)

    const width = 1600;
    const height = 500;
    const marginTop = 20;
    const marginRight = 30;
    const marginBottom = 30;
    const marginLeft = 40;

    // Declare the x (horizontal position) scale.
    const x = d3.scaleUtc(d3.extent(aapl, d => new Date(d.timestamp)), [marginLeft, width - marginRight]);

    // Declare the y (vertical position) scale.
    const y = d3.scaleLinear([0, d3.max(aapl, d => d.heartRate)], [height - marginBottom, marginTop]);
    const y2 = d3.scaleLinear([0, d3.max(aapl, d => d.power)], [height - marginBottom, marginTop]);

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
        .attr("d", line(aapl))

    svg.append("path")
        .attr("fill", "none")
        .attr("stroke", "green")
        .attr("stroke-width", 1)
        .attr("d", line2(aapl))

    document.getElementById("container").append(svg.node());
}

