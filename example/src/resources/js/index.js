const url = 'https://forbes400.herokuapp.com/api/forbes400/real-time';

$(document).ready(() => {
  fetch(url).then((resp) => resp.json()).then((data) => {
    console.log(data[0]);
    $('.center-div').empty();
    $('.center-div').append(getImage(data[0]));
    $('.center-div').append(getInfo(data[0]));
  }).catch(function(error) {
    if (err) {
      throw err;
    }
  });
});

const getImage = (data) => {
  return `<img src="${data.squareImage}" alt="Avatar" class="avatar shadow"/>`;
}

const getInfo = (data) => {
  return `<h3> ${data.name}
            <br/> <small class="text-muted">RealTimeWorth - ${round(data.realTimeWorth)}B</small>
          </h3>`;
}

const round = (data) => {
  data = data / 1000;
  return (Math.round(data * 10) / 10);
}
