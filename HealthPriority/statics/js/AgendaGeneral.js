
let valor = getCookie("session")
validarSesion(valor)


const signOut = document.getElementById("signOut")

signOut.addEventListener('click',(e)=>{
    //console.log("sesion cerrada")
    setCookie("session", "", -1);
    window.location.replace("http://127.0.0.1:5500/HealthPriority/templates/ingreso.html")
})

const formButton = document.getElementById("formRegistro")

formButton.addEventListener('submit', (e) => {
  e.preventDefault();


  let fecha = document.getElementById("Fecha")
  let jornada = document.getElementById("Jornada")

  
  
  console.log(fecha.value, "tipo: ", typeof(fecha.value))
  console.log(jornada.value,"tipo: ", typeof(jornada.value))
  

    let res=   fetch("http://localhost:8080/api/agendageneral",{
        method:'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            "Fecha": fecha.value,
            "Jornada": jornada.value,
            "Tipo": "General"
        }
        )
    }).then(res=>{
        let json = res.json
        console.log(res)
        return  res.ok ? res.json():Promise.reject(res)
    }).then(json=>{
        console.log(json)
        alert("Mensaje: "+json.name+", "+json.message)
    })
    .catch(er=>{
        console.log("Error", er)
    }).finally(()=>{
        console.log("Promesa recibida")
    })
    
});

