const controller = new  AbortController();
const signal = controller.signal;

async function renderView() {
    const fitFile = document.getElementById("fitFile")

    let formData = new FormData();
    formData.append("fitFile", fitFile.files[0])

    const data = await fetch("http://localhost:5432/api/fit/upload", {
        method: "POST",
        body: formData
    }).then(async function(response) {
        return await response.json()
    });

    console.log(data)
    const records = data.records

    renderChart2(records)
    renderMap(records)
}

function renderChart2(records) {
    const width = 1200;
    const height = 500;
    const marginTop = 20;
    const marginRight = 20;
    const marginBottom = 30;
    const marginLeft = 30;

    const x = d3.scaleUtc()
    .domain(d3.extent(records, d => new Date(d.timestamp)))
    .range([marginLeft, width - marginRight])

    const y = d3.scaleLinear([0, d3.max(records, d => d.power)], [height - marginBottom, marginTop])

    const line = d3.line().x(d => x(new Date(d.timestamp))).y(d => y(d.heartRate));
    const line2 = d3.line().x(d => x(new Date(d.timestamp))).y(d => y(d.power));

    // Create the SVG container.
    const svg = d3.create("svg")
    .attr("width", width)
    .attr("height", height)
    .attr("viewBox", [0, 0, width, height])
    .attr("style", "max-width: 100%; height: auto; height: intrinsic;");

    const brush = d3.brushX()
    .extent([[marginLeft, marginTop], [width - marginRight, height - marginBottom]])
    .on("end", brushed);

    svg.append("g")
        .call(brush)
        .on("dblclick", dblclicked);

    function dblclicked() {
        const selection = d3.brushSelection(this) ? null : x.range();
        d3.select(this).call(brush.move, selection);
    }

    function brushed({selection}) {
        if (selection === null) {
            console.log("no selection")
        } else {
            const [x0, x1] = selection.map(x.invert);
            filterRecords(records, x0, x1);
        }
    }
    // Add the x-axis.
    svg.append("g")
        .attr("transform", `translate(0,${height - marginBottom})`)
        .call(d3.axisBottom(x).ticks(width / 80).tickSizeOuter(0));

    // Add the y-axis, remove the domain line, add grid lines and a label.
    svg.append("g")
        .attr("transform", `translate(${marginLeft},0)`)
        .call(d3.axisLeft(y).ticks(height / 40))
        //.call(d3.axisLeft(y).ticks(height / 40))
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
    const coords = records.map(function(item) {
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


async function renderSummary() {
    //TODO: send post request and send summary data and return summary html template
    const response = await fetch("http://localhost:5432/api/template/summary")
    const data = await response.text()

    document.getElementById("summary").innerHTML = data
}

//TODO:
function filterRecords(records, start, end) {
    var filteredRecords = records.filter(function(item) {
        if (new Date(item.timestamp) >= start && new Date(item.timestamp) <= end) {
            return true;
        }
        return false;
    }).map(function(record) {
            return record;
    });

    const jason = filteredRecords.map(item => JSON.stringify(item))

    fetch("http://localhost:5432/api/template/summary", {
        method: "POST",
        body: "{\"records\":[" + jason + "]}"
    }).then(function(res) {
        console.log("response", res);
    }).catch(function(err) {
        console.error(err);
    });
}
