const valor = getCookie("session")
validarSesion(valor)





    let decoded 
    try {
            decoded = decodeJWT(valor)
            console.log(decoded)
           
    } catch (error) {
            console.error(error);
    }

    

    fetch("http://localhost:8080/api/me",{
        method:'POST',
        headers: {
            "Authorization": "Bearer " + valor,
            'Content-Type': 'application/json'
            
        },
        body: JSON.stringify({
            "Persona_id": decoded.id

        }
        )
    }).then(res=>{
        
        console.log(res)
        return  res.ok ? res.json():Promise.reject(res)
    }).then(json=>{
        console.log(json)
        const TablaMisCitas = document.getElementById("TablaPerfil")

        let fila = document.createElement('tr')
        let columna0 = document.createElement('td')
        columna0.innerHTML = `${json.Cedula}`
        let columna = document.createElement('td')
        columna.innerHTML = `${json.Nombres}`
        let columna2 = document.createElement('td')
        columna2.innerHTML = `${json.Apellidos}`
        let columna3 = document.createElement('td')
        columna3.innerHTML = `${json.Telefono}`
        let columna4 = document.createElement('td')
        columna4.innerHTML = `${json.Direccion}`
        let columna5 = document.createElement('td')
        columna5.innerHTML = `${json.Edad}`


        fila.appendChild(columna0)
        fila.appendChild(columna)
        fila.appendChild(columna2)
        fila.appendChild(columna3)
        fila.appendChild(columna4)
        fila.appendChild(columna5)
        TablaMisCitas.appendChild(fila);
        



 

    })
    .catch(er=>{
        console.log("Error", er)
    }).finally(()=>{
        console.log("Promesa recibida")
    })


const signOut = document.getElementById("signOut")

signOut.addEventListener('click',(e)=>{
    //console.log("sesion cerrada")
    setCookie("session", "", -1);
    window.location.replace("http://127.0.0.1:5500/HealthPriority/templates/ingreso.html")
})