
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

    console.log("entro a la peticion")
    let decoded 
try {
    decoded = decodeJWT(valor)
    console.log(decoded)
} catch (error) {
    console.error(error);
}

let nombre = document.getElementById("nombre")
let tipo = document.getElementById("tipo")
let documento = document.getElementById("documento")
const file = documento.files[0];
const reader = new FileReader();

reader.onload = () => {
    if (reader.result) {
        const base64String = reader.result.split(',')[1];
        console.log('Cadena base64:', base64String);

        fetch("http://localhost:8080/api/autorizacion", {
            method:'POST',
            headers: {
                "Authorization": "Bearer " + valor,
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                "Nombre": nombre.value,
                "Tipo": tipo.value,
                "Estado": true,
                "Person_id":decoded.id,
                "Imagen": base64String,
                "NombreImagen": file.name
            })
        }).then(res => {
            console.log(res);
            return res.ok ? res.json() : Promise.reject(res);
        }).then(json => {
            console.log(json);
            alert("Mensaje: " + json.name + ", " + json.message);
        }).catch(er => {
            console.log("Error", er);
        }).finally(() => {
            console.log("Promesa recibida");
        });
    } else {
        console.error('El archivo no se ha leído correctamente');
    }
};

reader.readAsDataURL(file);
    
    
});



  