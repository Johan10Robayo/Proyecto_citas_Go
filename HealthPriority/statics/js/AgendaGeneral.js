const formButton = document.getElementById("formRegistro")


let valor = getCookie("session")

if(valor.length == 0){
    window.location.replace("http://127.0.0.1:5500/templates/ingreso.html")
}else{
    console.log("logeado correctamente")
    console.log(valor)
    let decoded 
    try {
        decoded = decodeJWT(valor)
        console.log(decoded)
        
       
    } catch (error) {
        console.error(error);
    }

    const user = document.getElementById("user")
    user.innerHTML=`${decoded.username}`
    
    
}

const signOut = document.getElementById("signOut")

signOut.addEventListener('click',(e)=>{
    //console.log("sesion cerrada")
    setCookie("session", "", -1);
    window.location.replace("http://127.0.0.1:5501/templates/ingreso.html")
})



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
function decodeJWT(token) {
  
    const parts = token.split('.');
  
    if (parts.length !== 3) {
      throw new Error('Invalid token');
    }
  
    const header = JSON.parse(atob(parts[0]));
    const payload = JSON.parse(atob(parts[1]));
    const signature = parts[2];
  
  
    const secret = 'secret-key';
    const expectedSignature = btoa(`${parts[0]}.${parts[1]}`);
  
    /*if (signature !== expectedSignature) {
      throw new Error('Invalid signature');
    }*/
  
    return payload;
  }

  
function getCookie(cname) {
    var name = cname + "=";
    var ca = document.cookie.split(';');
    for(var i = 0; i < ca.length; i++) {
      var c = ca[i];
      while (c.charAt(0) == ' ') {
        c = c.substring(1);
      }
      if (c.indexOf(name) == 0) {
        return c.substring(name.length, c.length);
      }
    }
    return "";
}

function setCookie(cname, cvalue, exdays) {
    var d = new Date();
    d.setTime(d.getTime() + (exdays * 24 * 60 * 60 * 1000));
    var expires = "expires="+d.toUTCString();
    document.cookie = cname + "=" + cvalue + ";" + expires + ";path=/";
}