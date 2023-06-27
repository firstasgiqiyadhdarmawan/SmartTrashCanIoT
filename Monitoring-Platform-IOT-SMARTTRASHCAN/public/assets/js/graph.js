var data = [];
var time = [];

for (let index = 0; index < grafik_data.length; index++) {
    const element = grafik_data[index];
    if (element.Capacity > 100){
        element.Capacity = 100
    }
    data.push(element.Capacity)
    time.push(element.Date_formatter)
}

var sisa = 100 - data[0];
var pakai = data[0];

if (pakai > 100) {
    pakai = 100
}

document.getElementById("sisa").innerHTML = pakai+"%";

const ctx = document.getElementById('Donat');
const cty = document.getElementById('Line');

var dataReverse = data.reverse()
var timeReverse = time.reverse()

var dataSlice = dataReverse.slice(-8)
var timeSlice = timeReverse.slice(-8)

// Line
new Chart(cty, {
    type: 'line',
    data: {
        labels: timeSlice,
        datasets: [{
            label: 'Kapasitas Dalam Persen',
            data: dataSlice,
            borderWidth: 1,
        }]
    },
    options: {
        plugins: {
            legend: {
            position: 'bottom',
                display: true,
            }
        },
        responsive: true,
        scales: {
        x: {
            title: {
            display: true,
            text: 'Time'
            }
        },
        y: {
            title: {
            display: true,
            text: 'Capacity'
            },
            min: 0,
            max: 100,
        }
        }
    },
});

// Donat
new Chart(ctx, {
    type: 'pie',
    data: {
    labels: ['Sisa', 'Terpakai',],
    datasets: [{
        label: 'Dalam Persen',
        data: [pakai, sisa],
        borderWidth: 1
    }]
    },
    options : {
    plugins: {
            legend: {
            position: 'bottom',
                display: true,
            }
        },
    }
});