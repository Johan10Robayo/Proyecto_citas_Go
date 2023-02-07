const form = document.querySelector('form');
const inputs = form.elements;
let data = {};
form.addEventListener('submit', (e) => {
  e.preventDefault();

  // crea un objeto JSON con los datos del formulario

  for (let i = 0; i < inputs.length; i++) {
    data[inputs[i].name] = inputs[i].value;
  }});


(()=>{
    



    async function getData(){
        try {

            let res= await  fetch("http://localhost:8080/api/enterprise/findAllProductsById",{
                method:'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    'Nombres': id
                })
            })
            json = await res.json()
            console.log(res,json)



        } catch (error) {
            console.log(error)
        }finally{
            console.log("promesa recibida")
        }

    }


    getData();
})();
