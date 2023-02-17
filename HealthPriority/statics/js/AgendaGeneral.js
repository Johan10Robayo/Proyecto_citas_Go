const valor = getCookie("session")
validarSesion(valor)
const fechaInput = document.getElementById('Fecha');

const fechaActual = new Date();

const fechaMaxima = new Date();
const fechaMinima = new Date();
fechaMinima.setDate(fechaActual.getDate() + 1);
fechaMaxima.setDate(fechaActual.getDate() + 5);

fechaInput.setAttribute('min', fechaMinima.toISOString().split('T')[0]);
fechaInput.setAttribute('max', fechaMaxima.toISOString().split('T')[0]);

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
  
/* 	
	
	
	
	Cedula  int64*/

    let decoded 
        try {
            decoded = decodeJWT(valor)
            console.log(decoded)
           
        } catch (error) {
            console.error(error);
    }
    console.log(`VALOR DE VALOR: ${valor}`)

    

    fetch("http://localhost:8080/api/agendageneral",{
        method:'POST',
        headers: {
            "Authorization": "Bearer " + valor,
            'Content-Type': 'application/json'
            
        },
        body: JSON.stringify({
            "Fecha": fecha.value,
            "Jornada": jornada.value,
            "Tipo": "General",
            "Cedula": decoded.id

        }
        )
    }).then(res=>{
        
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

