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
        console.log(res)
        if(!res.ok)
            respuesta.innerHTML="Datos incorrectos, error de autenticación"
        
        return  res.ok ? res.json():Promise.reject(res)
    }).then(json=>{
        console.log(json)
        let token =json.message
        console.log(token)
        let decoded 
        try {
            decoded = decodeJWT(token)
            console.log(decoded)
           
        } catch (error) {
            console.error(error);
        }
    
        document.cookie = "session="+token+";path=/"
        if(decoded.role =="CLIENTE"){
            window.location.replace("http://127.0.0.1:5501/templates/indexLoginUser.html")
        }
            
        if(decoded.role =="FUNCIONARIO"){
            window.location.replace("http://127.0.0.1:5501/templates/indexLoginFunc.html")
        }
       
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