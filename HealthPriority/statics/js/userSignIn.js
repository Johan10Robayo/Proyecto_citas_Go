
let valor = getCookie("session")
validarSesion(valor)
alert("Sesión iniciada")


const signOut = document.getElementById("signOut")

signOut.addEventListener('click',(e)=>{
    //console.log("sesion cerrada")
    setCookie("session", "", -1);
    alert("Sesión cerrada")
    window.location.replace("http://127.0.0.1:5500/HealthPriority/templates/ingreso.html")
})

