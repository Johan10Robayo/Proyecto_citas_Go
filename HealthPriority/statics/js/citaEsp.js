const valor = getCookie("session")
validarSesion(valor)


    let decoded 
    try {
            decoded = decodeJWT(valor)
            console.log(decoded)
           
    } catch (error) {
            console.error(error);
    }

    

    fetch("http://localhost:8080/api/getAutorizaciones",{
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
        const CitaDisponible = document.getElementById('CitaDisponible');
        json.forEach(element => {
            const optionElement1 = document.createElement('option');
            optionElement1.value = `${element.Nombre}`;
            optionElement1.textContent = `${element.Nombre}`;
            CitaDisponible.appendChild(optionElement1);
        });



 

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


    
    


