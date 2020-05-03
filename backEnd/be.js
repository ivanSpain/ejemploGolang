const eLibreria=require('express')
const aplicacion=eLibreria()
const analizador=require("body-parser")

aplicacion.use(analizador.json())

aplicacion.post('/ejemploServicio',(request,response)=>{
    var respuesta={Nombre:request.body.Nombre}
    console.log("En el backend tu nombre es:"+request.body.Nombre)
    response.send(respuesta)
})

aplicacion.listen(2000,()=>{
    console.log("Escuchando en 2000")
})

