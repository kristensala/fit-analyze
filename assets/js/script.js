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

