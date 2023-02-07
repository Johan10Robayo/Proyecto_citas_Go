(()=>{
    const $enterpriseName =document.getElementById("enterpriseName")
    const $enterprisePhone =document.getElementById("enterprisePhone")
    const $enterpriseImg =document.getElementById("enterpriseImg")

    const $container =document.getElementById("containerProducts")
    $fragment = document.createDocumentFragment()
    



    async function getData(){
        try {
            const valores = window.location.search;
            //Creamos la instancia
            const urlParams = new URLSearchParams(valores);

            //Accedemos a los valores
            id = urlParams.get('id');
            console.log(id)
            console.log(valores);

            let res= await  fetch("http://localhost:8080/api/enterprise/findAllProductsById",{
                method:'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    'id': id
                })
            })
            json = await res.json()
            console.log(res,json)

            coordResponse = json.location
            coordArray = coordResponse.split(',')
            lat = parseFloat(coordArray[0])
            lng = parseFloat(coordArray[1])
            iniciarMap(lat,lng)

            
            $enterpriseImg.setAttribute("style",`background-image: url('${json.imagePath}')`)
            $enterpriseName.innerHTML=`SOMOS ${json.name}`
            $enterprisePhone.innerHTML=`Llamanos al ${json.phone}`

            arrayProducts = json.productList

            arrayProducts.forEach(product=>{

                const $div = document.createElement("div")
                $div.setAttribute("class","col")

                const $img =  document.createElement("img")
                $img.setAttribute("class","images_storedesc")
                $img.setAttribute("src",`${product.imagePath}`)

                const $divBody = document.createElement("div")
                $divBody.setAttribute("class","card-body")

                const $h5 = document.createElement("h5")
                $h5.setAttribute("class","card-title")
                $h5.innerHTML=`${product.name}`

                const $pCode = document.createElement("p")
                $pCode.setAttribute("class","card-text")
                $pCode.innerHTML=`Referencia: ${product.code}`

                const $pPrice = document.createElement("p")
                $pPrice.setAttribute("class","card-text")
                $pPrice.innerHTML=`Precio: $${product.price}`

                const $pDecription = document.createElement("p")
                $pDecription.setAttribute("class","card-text")
                $pDecription.innerHTML=`Descripcion: ${product.description}`


                $div.appendChild($img)
                $div.appendChild($divBody)
                $divBody.appendChild($h5)
                $divBody.appendChild($pCode)
                $divBody.appendChild($pPrice)
                $divBody.appendChild($pDecription)

                $fragment.appendChild($div)
            

            })

            $container.appendChild($fragment)


        } catch (error) {
            console.log(error)
        }finally{
            console.log("promesa recibida")
        }

    }


    getData();
})();
