let InputResultado = document.getElementById('resultado');
let num1 = document.getElementById('num1');
let num2 = document.getElementById("num2");
let check = document.getElementById("check");

const btn = document.getElementById("btnVerificar");

num1.innerText = Math.floor(Math.random() *10)
num2.innerText = Math.floor(Math.random() *10)

let valor1 = num1.innerText;
let valor2 = num2.innerText;

btn.addEventListener("click",()=>{
  let resultado = parseInt(valor1) + parseInt(valor2);
  let res = parseInt(InputResultado.value);

  if(res == resultado){
    alert("Validacion correcta");
    check.checked = true;
    document.getElementById('BtnRegistar').disabled = false;
  }else{
    InputResultado.value = " ";
    alert("Respuesta incorrecta, intente nuevamente");
    num1.innerText = Math.floor(Math.random() *10)
    num2.innerText = Math.floor(Math.random() *10)
    valor1 = num1.innerText;
    valor2 = num2.innerText;
  }

});

function iniciarMap(){
  var coord = {lat:4.137103113736017 ,lng:-73.63822874569446};
  var map = new google.maps.Map(document.getElementById('map'),{
    zoom: 15,
    center: coord
  });
  var marker = new google.maps.Marker({
    position: coord,
    map: map
  });
}