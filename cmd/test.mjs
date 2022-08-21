
const trials = 0
const concurrency = 10000

const fetches = new Array(concurrency)

const startTime = performance.now()


for (let i = 0; i < concurrency; i++) {

    fetches[i] = (
        fetch(`http://localhost:8090/`)
    )

}

const results = await Promise.all(fetches)

const totalTime = performance.now() - startTime

const responseSet = new Set()
let dataRaces = 0

for (const resp of results) {

    const text = await resp.text()

    if (responseSet.has(text))
        dataRaces++
    else
        responseSet.add(text)

}

console.log({ dataRaces, totalTime })