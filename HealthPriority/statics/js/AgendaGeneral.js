const formButton = document.getElementById("BtnCitaGeneral")
formButton.addEventListener('click', (e) => {
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