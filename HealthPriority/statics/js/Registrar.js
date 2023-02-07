const form = document.querySelector('form');
const inputs = form.elements;
let data = {};

form.addEventListener('submit', (e) => {
  e.preventDefault();

  // crea un objeto JSON con los datos del formulario

  for (let i = 0; i < inputs.length; i++) {
    data[inputs[i].name] = inputs[i].value;
  }
  

    let res=   fetch("http://localhost:8080/api/registrar",{
        method:'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(
            data
        )
    }).then(json=>{
        console.log(res,json)
        alert(json.name+json.message);
    }).catch(er=>{
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
