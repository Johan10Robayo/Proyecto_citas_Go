const formButton = document.getElementById("BtnRegistar")
//let data = {}
formButton.addEventListener('click', (e) => {
  e.preventDefault();

  /*for (let i = 0; i < inputs.length; i++) {
    if(inputs[i].name=="Embarazo"){
        estado =inputs[i].value
        console.log(estado)
        if (estado== "true"){
            data['"'+inputs[i].name+'"'] = true
        }else{
            data['"'+inputs[i].name+'"'] = false
        }
    }else if (inputs[i].name=="Cedula" || inputs[i].name=="Edad" || inputs[i].name=="Telefono"){
        data['"'+inputs[i].name+'"'] = parseInt(inputs[i].value)
    }else{
        data['"'+inputs[i].name+'"'] = inputs[i].value
    }
  }*/
  let nombres = document.getElementById("Nombres")
  let apellidos = document.getElementById("Apellidos")
  let telefono = document.getElementById("Telefono")
  let edad = document.getElementById("Edad")
  let cedula = document.getElementById("Cedula")
  let embarazo = document.getElementById("Embarazo")
  let comorbilidad = document.getElementById("Comorbilidad")
  let correo = document.getElementById("Correo")
  let pass = document.getElementById("Password")
  let direccion = document.getElementById("Direccion")
  let emabarazojson
  if (embarazo.value=="true"){
    emabarazojson = true
  }else{
    emabarazojson = false
  }
  console.log(nombres.value, "tipo: ", typeof(nombres.value))
  console.log(apellidos.value,"tipo: ", typeof(apellidos.value))
  console.log(telefono.value,"tipo: ", typeof(telefono.value))
  console.log(edad.value,"tipo: ", typeof(edad.value))
  console.log(cedula.value,"tipo: ", typeof(cedula.value))
  console.log(embarazo.value,"tipo: ", typeof(embarazo.value))
  console.log(comorbilidad.value,"tipo: ", typeof(comorbilidad.value))
  console.log(correo.value,"tipo: ", typeof(correo.value))
  console.log(pass.value,"tipo: ", typeof(pass.value))
  console.log(direccion.value,"tipo: ", typeof(direccion.value))
  console.log(emabarazojson,"tipo: ", typeof(emabarazojson))

    let res=   fetch("http://localhost:8080/api/registrar",{
        method:'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            "Nombres":nombres.value,
            "Apellidos":apellidos.value,
            "Edad":parseInt(edad.value),
            "Telefono":parseInt(telefono.value),
            "Embarazo":emabarazojson,
            "Comorbilidad":comorbilidad.value,
            "Correo":correo.value,
            "Password":pass.value,
            "Cedula":parseInt(cedula.value),
            "Direccion":direccion.value
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

/*
(()=>{
    



    async function getData(){
        try {

            let res= await  fetch("http://localhost:8080/api/registrar",{
                method:'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(
                    data
                )
            })
            json = await res.json()
            console.log(res,json)
            alert(json.name+json.message);


        } catch (error) {
            console.log(error)
            alert(json.name+json.message);
        }finally{
            console.log("promesa recibida")
        }

    }


    getData();
})();
*/
