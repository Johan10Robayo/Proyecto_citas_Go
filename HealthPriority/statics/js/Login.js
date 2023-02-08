const formButton = document.getElementById("formLogin")
const respuesta = document.getElementById("response")
formButton.addEventListener('submit', (e) => {
  e.preventDefault();


  let correo = document.getElementById("email")
  let pass = document.getElementById("pass")

  console.log(correo.value, "tipo: ", typeof(correo.value))
  console.log(pass.value,"tipo: ", typeof(pass.value))

    let res=   fetch("http://localhost:8080/api/login",{
        method:'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            "User":correo.value,
            "Password":pass.value,
        }
        )
    }).then(res=>{
        let json = res.json
        console.log(res)
        if(res.ok!=200){
            respuesta.innerHTML="Datos incorrectos, error de autenticación"
        }
        return  res.ok ? res.json():Promise.reject(res)
    }).then(json=>{
        console.log(json)
        respuesta.innerHTML=""
    })
    .catch(er=>{
        console.log("Error", er)
    }).finally(()=>{
        console.log("Promesa recibida")
    })
    
});